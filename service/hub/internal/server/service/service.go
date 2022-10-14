package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/common/websocket"
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
	rabbitmqQueue      rabbitmq.Queue
	WsHub              *websocket.WSHub
}

func New() (s *Service) {
	s = &Service{
		employer: shedlock.New(),
		WsHub:    websocket.NewHub(),
	}

	var err error
	s.rabbitmqConnection, err = rabbitmq.Dial(config.ConfigHub.RabbitMQ.String())
	if err != nil {
		loggerx.Global().Error("rabbitmq dail failed", zap.Error(err))
	}

	s.rabbitmqChannel, err = s.rabbitmqConnection.Channel()
	if err != nil {
		loggerx.Global().Error("rabbitmqConnection failed", zap.Error(err))
	}

	if err := s.rabbitmqChannel.ExchangeDeclare(protocol.ExchangeJob, "direct", true, false, false, false, nil); err != nil {
		loggerx.Global().Error("rabbitmqChannel ExchangeDeclare Job failed", zap.Error(err))
	}

	if err := s.rabbitmqChannel.ExchangeDeclare(protocol.ExchangeRefresh, "fanout", true, false, false, false, nil); err != nil {
		loggerx.Global().Error("rabbitmqChannel ExchangeDeclare Refresh failed", zap.Error(err))
	}

	if s.rabbitmqQueue, err = s.rabbitmqChannel.QueueDeclare(
		"", false, false, true, false, nil,
	); err != nil {
		loggerx.Global().Error("rabbitmq QueueDeclare failed", zap.Error(err))
	}

	if err := s.rabbitmqChannel.QueueBind(
		s.rabbitmqQueue.Name, "", protocol.ExchangeRefresh, false, nil,
	); err != nil {
		loggerx.Global().Error("rabbitmq QueueBind failed", zap.Error(err))
	}

	return s
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

		dao.InitializeAddressStatus(ctx, message.Address)
	}

	networks := []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
		protocol.NetworkArweave, protocol.NetworkXDAI, protocol.NetworkZkSync, protocol.NetworkCrossbell,
		protocol.NetworkEIP1577,
	}

	go func() {
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
				return
			}
		}
	}()
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
			return
		}
	}
}

func (s *Service) SubscribeIndexerRefreshMessage() {
	deliveryCh, err := s.rabbitmqChannel.Consume(s.rabbitmqQueue.Name, "", true, false, false, false, nil)
	if err != nil {
		return
	}
	for delivery := range deliveryCh {
		message := protocol.RefreshMessage{}
		if err := json.Unmarshal(delivery.Body, &message); err != nil {
			loggerx.Global().Error("failed to unmarshal message", zap.Error(err))
			continue
		}
		fmt.Println("get msg from mq: ", message)
		s.WsHub.Broadcast <- delivery.Body
	}
}
