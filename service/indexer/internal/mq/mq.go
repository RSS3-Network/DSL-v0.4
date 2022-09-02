package mq

import (
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	rabbitmq "github.com/rabbitmq/amqp091-go"
)

var (
	RabbitmqConnection *rabbitmq.Connection
	RabbitmqChannel    *rabbitmq.Channel

	RabbitmqQueue      rabbitmq.Queue
	RabbitmqAssetQueue rabbitmq.Queue
)

func Initialize() (err error) {
	RabbitmqConnection, err = rabbitmq.Dial(config.ConfigIndexer.RabbitMQ.String())
	if err != nil {
		return err
	}

	RabbitmqChannel, err = RabbitmqConnection.Channel()
	if err != nil {
		return err
	}

	if err := RabbitmqChannel.ExchangeDeclare(
		protocol.ExchangeJob, "direct", true, false, false, false, nil,
	); err != nil {
		return err
	}

	if RabbitmqQueue, err = RabbitmqChannel.QueueDeclare(
		protocol.IndexerWorkQueue, false, false, false, false, nil,
	); err != nil {
		return err
	}

	if err := RabbitmqChannel.QueueBind(
		RabbitmqQueue.Name, protocol.IndexerWorkRoutingKey, protocol.ExchangeJob, false, nil,
	); err != nil {
		return err
	}

	// asset mq
	if RabbitmqAssetQueue, err = RabbitmqChannel.QueueDeclare(
		protocol.IndexerAssetQueue, false, false, false, false, nil,
	); err != nil {
		return err
	}

	if err := RabbitmqChannel.QueueBind(
		RabbitmqAssetQueue.Name, protocol.IndexerAssetRoutingKey, protocol.ExchangeJob, false, nil,
	); err != nil {
		return err
	}

	return nil
}
