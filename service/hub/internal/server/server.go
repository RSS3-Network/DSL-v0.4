package server

import (
	"net"
	"strconv"

	"github.com/labstack/echo"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler"
	rabbitmq "github.com/rabbitmq/amqp091-go"
)

var _ command.Interface = &Server{}

type Server struct {
	config             *config.Config
	httpServer         *echo.Echo
	httpHandler        *handler.Handler
	databaseClient     *database.Client
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
}

func (s *Server) Initialize() (err error) {
	s.rabbitmqConnection, err = rabbitmq.Dial(s.config.RabbitMQ.URL)
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
	}

	s.httpServer.GET("/:address/actions", s.httpHandler.GetActionListFunc)

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
