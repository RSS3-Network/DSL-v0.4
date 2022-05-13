package server

import (
	"encoding/json"

	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

var _ command.Interface = &Server{}

type Server struct {
	config             *config.Config
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	rabbitmqQueue      rabbitmq.Queue
	databaseClient     *database.Client
}

func (s *Server) Initialize() (err error) {
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

		switch "" {
		case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartCain:
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
