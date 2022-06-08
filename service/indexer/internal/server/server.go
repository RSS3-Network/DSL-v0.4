package server

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/poap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/mirror"
	poapworker "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/poap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/swap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token/coinmarketcap"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ command.Interface = &Server{}

type Server struct {
	config             *config.Config
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	rabbitmqQueue      rabbitmq.Queue
	databaseClient     *gorm.DB
	redisClient        *redis.Client
	datasources        []datasource.Datasource
	workers            []worker.Worker
	employer           *shedlock.Employer
}

func (s *Server) Initialize() (err error) {
	var exporter trace.SpanExporter

	if s.config.OpenTelemetry == nil {
		if exporter, err = opentelemetry.DialWithPath(opentelemetry.DefaultPath); err != nil {
			return err
		}
	} else if s.config.OpenTelemetry.Enabled {
		if exporter, err = opentelemetry.DialWithURL(s.config.OpenTelemetry.String()); err != nil {
			return err
		}
	}

	otel.SetTracerProvider(trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("pregod-indexer"),
			semconv.ServiceVersionKey.String("v0.5.0"),
		)),
	))

	s.databaseClient, err = database.Dial(s.config.Postgres.String(), true)
	if err != nil {
		return err
	}

	if s.redisClient, err = cache.Dial(s.config.Redis); err != nil {
		return err
	}

	coinmarketcap.Init(s.config.CoinMarketCap.APIKey)

	s.rabbitmqConnection, err = rabbitmq.Dial(s.config.RabbitMQ.String())
	if err != nil {
		return err
	}

	s.rabbitmqChannel, err = s.rabbitmqConnection.Channel()
	if err != nil {
		return err
	}

	if err := s.rabbitmqChannel.ExchangeDeclare(
		protocol.ExchangeJob, "direct", true, false, false, false, nil,
	); err != nil {
		return err
	}

	if s.rabbitmqQueue, err = s.rabbitmqChannel.QueueDeclare(
		"", false, false, false, false, nil,
	); err != nil {
		return err
	}

	if err := s.rabbitmqChannel.QueueBind(
		s.rabbitmqQueue.Name, "", protocol.ExchangeJob, false, nil,
	); err != nil {
		return err
	}

	s.datasources = []datasource.Datasource{
		moralis.New(s.config.Moralis.Key), arweave.New(), poap.New(), blockscout.New(),
	}

	s.workers = []worker.Worker{
		token.New(s.databaseClient), swap.New(s.databaseClient), mirror.New(), poapworker.New(),
	}

	s.employer = shedlock.New(s.redisClient)

	for _, internalWorker := range s.workers {
		if err := internalWorker.Initialize(context.Background()); err != nil {
			return err
		}

		if internalWorker.Jobs() == nil {
			continue
		}

		for _, job := range internalWorker.Jobs() {
			if err := s.employer.AddJob(job.Name(), job.Spec(), job.Timeout(), job); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	s.employer.Start()

	defer s.employer.Stop()

	deliveryCh, err := s.rabbitmqChannel.Consume(s.rabbitmqQueue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for delivery := range deliveryCh {
		message := protocol.Message{}
		if err := json.Unmarshal(delivery.Body, &message); err != nil {
			logrus.Errorln(err)

			continue
		}

		go func() {
			if err := s.handle(context.Background(), &message); err != nil {
				logrus.Errorln(err)
			}
		}()
	}

	return nil
}

func (s *Server) handle(ctx context.Context, message *protocol.Message) (err error) {
	logrus.Infoln(message.Address, message.Network)

	// TODO Query the latest time of existing data, which is used to avoid data being reworked

	transfers := make([]model.Transfer, 0)

	for _, internalDatasource := range s.datasources {
		supportedNetwork := false

		for _, network := range internalDatasource.Networks() {
			if network == message.Network {
				supportedNetwork = true

				break
			}
		}

		if !supportedNetwork {
			continue
		}

		internalTransfers, err := internalDatasource.Handle(ctx, message)
		if err != nil {
			logrus.Errorln(internalDatasource.Name(), err)

			return err
		}

		transfers = append(transfers, internalTransfers...)
	}

	if len(transfers) == 0 {
		return nil
	}

	if err := s.databaseClient.
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).
		Create(transfers).Error; err != nil {
		return err
	}

	for _, internalWorker := range s.workers {
		supportedNetwork := false

		for _, network := range internalWorker.Networks() {
			if network == message.Network {
				supportedNetwork = true

				break
			}
		}

		if !supportedNetwork {
			continue
		}

		internalTransfers, err := internalWorker.Handle(context.Background(), message, transfers)
		if err != nil {
			logrus.Errorln(internalWorker.Name(), err)

			return err
		}

		if len(internalTransfers) == 0 {
			continue
		}

		if err := s.databaseClient.
			Clauses(clause.OnConflict{
				DoUpdates: clause.AssignmentColumns([]string{"metadata"}),
				UpdateAll: true,
			}).
			Create(internalTransfers).Error; err != nil {
			logrus.Errorln(err)
			return err
		}

		transfers = internalTransfers
	}

	return nil
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
