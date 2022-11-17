package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type Service struct {
	employer           *shedlock.Employer
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
}

func New() (s *Service) {
	s = &Service{
		employer: shedlock.New(),
	}

	if err := s.connectMQ(); err != nil {
		loggerx.Global().Fatal("connect mq failed", zap.Error(err))
	}

	go func() {
		maxRetry := 3
		for {
			<-s.rabbitmqConnection.NotifyClose(make(chan *rabbitmq.Error))
			loggerx.Global().Error("rabbitmq connection closed, reconnecting...")
			time.Sleep(10 * time.Second)
			if err := s.connectMQ(); err != nil {
				loggerx.Global().Error("connect mq failed", zap.Error(err))
			} else {
				loggerx.Global().Info("connect mq success", zap.Error(err))
			}
			maxRetry--
			if maxRetry == 0 {
				panic("rabbitmq reconnect failed")
			}
		}
	}()
	return s
}

func (s *Service) connectMQ() error {
	var err error
	s.rabbitmqConnection, err = rabbitmq.Dial(config.ConfigHub.RabbitMQ.String())
	if err != nil {
		loggerx.Global().Error("rabbitmq dail failed", zap.Error(err))
		return err
	}

	s.rabbitmqChannel, err = s.rabbitmqConnection.Channel()
	if err != nil {
		loggerx.Global().Error("rabbitmqConnection failed", zap.Error(err))
		return err
	}

	if err := s.rabbitmqChannel.ExchangeDeclare(protocol.ExchangeJob, "direct", true, false, false, false, nil); err != nil {
		loggerx.Global().Error("rabbitmqChannel ExchangeDeclare failed", zap.Error(err))
		return err
	}
	return nil
}

// publishIndexerMessage create a rabbitmq job to index the latest user data
func (s *Service) PublishIndexerMessage(ctx context.Context, message protocol.Message) {
	tracer := otel.Tracer("PublishIndexerMessage")
	_, rabbitmqSnap := tracer.Start(ctx, "rabbitmq")

	defer rabbitmqSnap.End()

	if len(message.Address) == 0 {
		return
	}

	if !message.IgnoreNote {
		var address model.Address
		_ = database.Global().Model(&dbModel.Address{}).
			Where("address = ?", message.Address).
			First(&address).Error
		if address.Address != "" && address.UpdatedAt.After(time.Now().Add(-2*time.Minute)) {
			return
		}

		address.Address = message.Address
		dao.InitializeAddressStatus(ctx, address)
	}

	networks := []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
		protocol.NetworkArweave, protocol.NetworkXDAI, protocol.NetworkZkSync, protocol.NetworkCrossbell,
		protocol.NetworkOptimism, protocol.NetworkAvalanche, protocol.NetworkCelo, protocol.NetworkFantom,
		protocol.NetworkEIP1577,
	}

	for _, network := range networks {
		message.Network = network

		messageData, err := json.Marshal(&message)
		if err != nil {
			return
		}

		if err := s.rabbitmqChannel.Publish(protocol.ExchangeJob, protocol.IndexerWorkRoutingKey, false, false, rabbitmq.Publishing{
			ContentType: protocol.ContentTypeJSON,
			Body:        messageData,
		}); err != nil {
			loggerx.Global().Error("publish indexer message failed", zap.Error(err))
			return
		}
	}
}

// publishIndexerAssetMessage create a rabbitmq job to index the latest user data
func (s *Service) PublishIndexerAssetMessage(ctx context.Context, address string) {
	tracer := otel.Tracer("PublishIndexerAssetMessage")
	_, rabbitmqSnap := tracer.Start(ctx, "rabbitmq")

	defer rabbitmqSnap.End()

	networks := []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon,
	}

	for _, network := range networks {
		message := protocol.Message{
			Address: address,
			Network: network,
		}

		messageData, err := json.Marshal(&message)
		if err != nil {
			return
		}

		if err := s.rabbitmqChannel.Publish(protocol.ExchangeJob, protocol.IndexerAssetRoutingKey, false, false, rabbitmq.Publishing{
			ContentType: protocol.ContentTypeJSON,
			Body:        messageData,
		}); err != nil {
			loggerx.Global().Error("publish indexer asset message failed", zap.Error(err))
			return
		}
	}
}
