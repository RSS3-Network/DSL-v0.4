package server

import (
	"net"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ command.Interface = &Server{}

type Server struct {
	config             *config.Config
	httpServer         *echo.Echo
	httpHandler        *handler.Handler
	databaseClient     *gorm.DB
	redisClient        *redis.Client
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	logger             *zap.Logger
}

func (s *Server) Initialize() (err error) {
	s.logger, err = zap.NewProduction()
	if err != nil {
		return err
	}

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
			semconv.ServiceNameKey.String("pregod-hub"),
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

	if err := s.rabbitmqChannel.ExchangeDeclare(protocol.ExchangeJob, "direct", true, false, false, false, nil); err != nil {
		return err
	}

	if s.redisClient, err = cache.Dial(s.config.Redis); err != nil {
		return err
	}

	s.httpServer = echo.New()

	s.httpServer.HideBanner = true
	s.httpServer.HidePort = true

	s.httpHandler = &handler.Handler{
		DatabaseClient:     s.databaseClient,
		RabbitmqConnection: s.rabbitmqConnection,
		RabbitmqChannel:    s.rabbitmqChannel,
	}

	s.httpServer.HTTPErrorHandler = s.httpHandler.ErrorFunc

	s.httpServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

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
