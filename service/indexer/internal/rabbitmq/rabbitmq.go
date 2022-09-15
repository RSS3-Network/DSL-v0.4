package rabbitmq

import (
	"sync"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	rabbitmq "github.com/rabbitmq/amqp091-go"
)

var (
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	rabbitmqQueue      rabbitmq.Queue
	rabbitmqAssetQueue rabbitmq.Queue

	globalLocker sync.RWMutex
)

func GetRabbitmqConnection() *rabbitmq.Connection {
	globalLocker.RLock()

	defer globalLocker.RUnlock()

	return rabbitmqConnection
}

func GetRabbitmqChannel() *rabbitmq.Channel {
	globalLocker.RLock()

	defer globalLocker.RUnlock()

	return rabbitmqChannel
}

func GetRabbitmqQueue() rabbitmq.Queue {
	globalLocker.RLock()

	defer globalLocker.RUnlock()

	return rabbitmqQueue
}

func GetRabbitmqAssetQueue() rabbitmq.Queue {
	globalLocker.RLock()

	defer globalLocker.RUnlock()

	return rabbitmqAssetQueue
}

func InitializeMQ(mq *configx.RabbitMQ) (err error) {
	rabbitmqConnection, err = rabbitmq.Dial(mq.String())
	if err != nil {
		return err
	}

	rabbitmqChannel, err = rabbitmqConnection.Channel()
	if err != nil {
		return err
	}

	if err := rabbitmqChannel.ExchangeDeclare(
		protocol.ExchangeJob, "direct", true, false, false, false, nil,
	); err != nil {
		return err
	}

	if rabbitmqQueue, err = rabbitmqChannel.QueueDeclare(
		protocol.IndexerWorkQueue, false, false, false, false, nil,
	); err != nil {
		return err
	}

	if err := rabbitmqChannel.QueueBind(
		rabbitmqQueue.Name, protocol.IndexerWorkRoutingKey, protocol.ExchangeJob, false, nil,
	); err != nil {
		return err
	}

	// asset mq
	if rabbitmqAssetQueue, err = rabbitmqChannel.QueueDeclare(
		protocol.IndexerAssetQueue, false, false, false, false, nil,
	); err != nil {
		return err
	}

	if err := rabbitmqChannel.QueueBind(
		rabbitmqAssetQueue.Name, protocol.IndexerAssetRoutingKey, protocol.ExchangeJob, false, nil,
	); err != nil {
		return err
	}

	return nil
}
