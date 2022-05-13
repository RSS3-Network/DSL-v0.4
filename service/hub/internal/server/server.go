package server

import (
	"context"
	"net"
	"strconv"

	"entgo.io/ent/dialect"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.uber.org/zap"
)

var _ command.Interface = &Server{}

type Server struct {
	config             *config.Config
	httpServer         *echo.Echo
	httpHandler        *handler.Handler
	databaseClient     *database.Client
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	exporter           *jaeger.Exporter
	tracerProvider     *trace.TracerProvider
	logger             *zap.Logger
}

func (s *Server) Initialize() (err error) {
	s.logger, err = zap.NewProduction()
	if err != nil {
		return err
	}

	s.exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(s.config.OpenTelemetry.String())))
	if err != nil {
		return err
	}

	s.tracerProvider = trace.NewTracerProvider(
		trace.WithBatcher(s.exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("pregod-hub"),
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

	if err := s.rabbitmqChannel.ExchangeDeclare(protocol.ExchangeJob, "direct", true, false, false, false, nil); err != nil {
		return err
	}

	s.httpServer = echo.New()
	s.httpServer.HideBanner = true
	s.httpServer.HidePort = true

	s.httpHandler = &handler.Handler{
		DatabaseClient:     s.databaseClient,
		RabbitmqConnection: s.rabbitmqConnection,
		RabbitmqChannel:    s.rabbitmqChannel,
		TracerProvider:     s.tracerProvider,
	}

	s.httpServer.HTTPErrorHandler = s.httpHandler.ErrorFunc

	s.httpServer.GET("/addresses/:address/actions", s.httpHandler.GetActionListFunc)

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	return s.httpServer.Start(net.JoinHostPort(s.config.HTTP.Host, strconv.Itoa(s.config.HTTP.Port)))
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
