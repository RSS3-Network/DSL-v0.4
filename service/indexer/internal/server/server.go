package server

import (
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/message"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

var _ command.Interface = &Server{}

type Server struct {
	config             *config.Config
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	databaseClient     *database.Client
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

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	queue, err := s.rabbitmqChannel.QueueDeclare("", false, false, false, false, nil)
	if err != nil {
		return err
	}

	if err := s.rabbitmqChannel.QueueBind(queue.Name, "", "", false, nil); err != nil {
		return err
	}

	deliveryCh, err := s.rabbitmqChannel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for delivery := range deliveryCh {
		logrus.Infoln(delivery)

		switch "" {
		case message.NetworkEthereum, message.NetworkPolygon, message.NetworkBinanceSmartCain:
		case message.NetworkZkSync:
		}
	}

	return nil
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
