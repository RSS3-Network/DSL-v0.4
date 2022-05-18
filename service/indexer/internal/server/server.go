package server

import (
	"context"
	"encoding/json"

	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/swap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
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
	exporter           *jaeger.Exporter
	tracerProvider     *trace.TracerProvider
	workers            []worker.Worker
	datasources        []datasource.Datasource
}

func (s *Server) Initialize() (err error) {
	s.exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(s.config.OpenTelemetry.String())))
	if err != nil {
		return err
	}

	s.tracerProvider = trace.NewTracerProvider(
		trace.WithBatcher(s.exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("pregod-indexer"),
		)),
	)

	otel.SetTracerProvider(s.tracerProvider)

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

	s.workers = append(s.workers, token.New(), swap.New())

	s.datasources = append(s.datasources, moralis.New(s.config.Moralis.Key))

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
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

func (s *Server) handle(ctx context.Context, message *protocol.Message) (err error) {
	logrus.Infoln(message.Address, message.Network)

	// TODO Query the latest time of existing data, which is used to avoid data being reworked

	transfers := make([]model.Transfer, 0)

	for _, internalDatasource := range s.datasources {
		transfers, err = internalDatasource.Handle(ctx, message)
		if err != nil {
			logrus.Errorln(err)

			return err
		}
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
			}
		}

		if !supportedNetwork {
			continue
		}

		transfers, err := internalWorker.Handle(context.Background(), message, transfers)
		if err != nil {
			logrus.Errorln(err)

			return err
		}

		if len(transfers) == 0 {
			continue
		}

		if err := s.databaseClient.
			Clauses(clause.OnConflict{
				DoUpdates: clause.AssignmentColumns([]string{"metadata"}),
				UpdateAll: true,
			}).
			Create(transfers).Error; err != nil {
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
