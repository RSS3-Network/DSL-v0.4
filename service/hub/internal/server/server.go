package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/middlewarex"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/validatorx"
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
	httpServer         *echo.Echo
	httpHandler        *handler.Handler
	databaseClient     *gorm.DB
	redisClient        *redis.Client
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	logger             *zap.Logger
}

func (s *Server) Initialize() (err error) {
	// do not check the err here
	s.logger, _ = zap.NewProduction()
	// if err != nil {
	//	return err
	//}

	var exporter trace.SpanExporter

	if config.ConfigHub.OpenTelemetry == nil {
		if exporter, err = opentelemetry.DialWithPath(opentelemetry.DefaultPath); err != nil {
			logger.Global().Error("opentelemetry DialWithPath failed", zap.Error(err))
		}
	} else if config.ConfigHub.OpenTelemetry.Enabled {
		if exporter, err = opentelemetry.DialWithURL(config.ConfigHub.OpenTelemetry.String()); err != nil {
			logger.Global().Error("opentelemetry DialWithURL failed", zap.Error(err))
		}
	}

	otel.SetTracerProvider(trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("pregod-1-1-hub"),
			semconv.ServiceVersionKey.String(protocol.Version),
		)),
	))

	s.databaseClient = config.ConfigHub.DatabaseClient

	s.rabbitmqConnection, err = rabbitmq.Dial(config.ConfigHub.RabbitMQ.String())
	if err != nil {
		logger.Global().Error("rabbitmq dail failed", zap.Error(err))
	}

	s.rabbitmqChannel, err = s.rabbitmqConnection.Channel()
	if err != nil {
		logger.Global().Error("rabbitmqConnection failed", zap.Error(err))
	}

	if err := s.rabbitmqChannel.ExchangeDeclare(protocol.ExchangeJob, "direct", true, false, false, false, nil); err != nil {
		logger.Global().Error("rabbitmqChannel ExchangeDeclare failed", zap.Error(err))
	}

	if s.redisClient, err = cache.Dial(config.ConfigHub.Redis); err != nil {
		logger.Global().Error("redis dial failed", zap.Error(err))
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

	s.httpServer.Validator = validatorx.Default

	s.httpServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	s.httpServer.Use(middlewarex.ZapLogger(s.logger))

	s.httpServer.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":        "ok",
			"version":       fmt.Sprintf("%s-%s", protocol.Version, protocol.Build),
			"documentation": "https://docs.rss3.io/PreGod/api/",
		})
	})

	// GET
	s.httpServer.GET("/notes/:address", s.httpHandler.GetNotesFunc, middlewarex.APIMiddleware)
	s.httpServer.GET("/assets/:address", s.httpHandler.GetAssetsFunc, middlewarex.APIMiddleware)
	s.httpServer.GET("/exchanges/:exchange_type", s.httpHandler.GetExchangeListFunc)
	s.httpServer.GET("/platforms/:platform_type", s.httpHandler.GetPlatformListFunc)
	s.httpServer.GET("/profiles/:address", s.httpHandler.GetProfilesFunc, middlewarex.APIMiddleware)
	s.httpServer.GET("/ns/:address", s.httpHandler.GetNameResolve)

	// POST
	s.httpServer.POST("/notes", s.httpHandler.BatchGetNotesFunc, middlewarex.CheckAPIKeyMiddleware)
	s.httpServer.POST("/profiles", s.httpHandler.BatchGetProfilesFunc, middlewarex.CheckAPIKeyMiddleware)

	// API KEY
	s.httpServer.POST("/apikey/apply", s.httpHandler.PostAPIKeyFunc)
	s.httpServer.GET("/apikey", s.httpHandler.GetAPIKeyFunc)

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	return s.httpServer.Start(net.JoinHostPort(config.ConfigHub.HTTP.Host, strconv.Itoa(config.ConfigHub.HTTP.Port)))
}
