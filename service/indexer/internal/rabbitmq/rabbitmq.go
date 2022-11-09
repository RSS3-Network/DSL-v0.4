package rabbitmq

import (
	"sync"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
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
	err = connect(mq)

	go func() {
		maxRetry := 3
		for {
			<-rabbitmqConnection.NotifyClose(make(chan *rabbitmq.Error))
			loggerx.Global().Error("rabbitmq connection closed, reconnecting...")
			if err := connect(mq); err != nil {
				loggerx.Global().Error("connect mq failed", zap.Error(err))
			}
			maxRetry--
			if maxRetry == 0 {
				panic("rabbitmq reconnect failed")
			}
		}
	}()

	return err
}

func connect(mq *configx.RabbitMQ) (err error) {
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
