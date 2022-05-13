package server

import (
	"context"
	"encoding/json"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	moralisx "github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/indexer"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/indexer/moralis"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

var _ command.Interface = &Server{}

type Server struct {
	config             *config.Config
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	rabbitmqQueue      rabbitmq.Queue
	databaseClient     *database.Client
	indexerMoralis     indexer.Worker
	exporter           *jaeger.Exporter
	tracerProvider     *trace.TracerProvider
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

	s.databaseClient, err = database.Open(dialect.Postgres, s.config.Postgres.String())
	if err != nil {
		return err
	}

	if err := s.databaseClient.Schema.Create(context.Background()); err != nil {
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

	s.indexerMoralis = &moralis.Indexer{
		TracerProvider: s.tracerProvider,
		DatabaseClient: s.databaseClient,
		MoralisClient:  moralisx.NewClient(s.config.Moralis.Key),
	}

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

		logrus.Infoln(message.Address, message.Network)

		switch message.Network {
		case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartCain:
			if err := s.indexerMoralis.Handle(context.Background(), &message); err != nil {
				logrus.Errorln(err)

				continue
			}
		case protocol.NetworkZkSync:
		}
	}

	return nil
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
