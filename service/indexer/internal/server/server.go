package server

import (
	"context"
	"encoding/json"
	"github.com/naturalselectionlabs/pregod/common/cache"

	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/mirror"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/swap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token"
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
	datasources        []datasource.Datasource
	workers            []worker.Worker
}

func (s *Server) Initialize() (err error) {
	var exporter trace.SpanExporter

	if s.config.OpenTelemetry == nil {
		if exporter, err = opentelemetry.DialWithPath(opentelemetry.DefaultPath); err != nil {
			return err
		}
	} else {
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

<<<<<<< develop
	s.datasources = []datasource.Datasource{
		moralis.New(s.config.Moralis.Key), arweave.New(),
	}
=======
	if err := cache.Dial(*s.config.Redis); err != nil {
		return err
	}

	s.workers = append(s.workers, token.New(), swap.New())
>>>>>>> feat: add redis cache for swap pools

	s.workers = []worker.Worker{
		token.New(), swap.New(), mirror.New(),
	}

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	if err := s.SeedSwapPoolCache(); err != nil {
		return err
	}

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

func (s *Server) SeedSwapPoolCache() error {
	swapPools, err := database.GetSwapPools(s.databaseClient)
	if err != nil {
		return err
	}

	for _, pool := range swapPools {
		if err := cache.HSet(context.Background(), "swappools", pool.ContractAddress, pool); err != nil {
			return err
		}
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
			logrus.Errorln(err)

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

		for _, network := range internalWorker.Network() {
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
			logrus.Errorln(err)

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
	}

	return nil
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
