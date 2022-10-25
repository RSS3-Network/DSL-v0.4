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
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/websocket"
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
	DeliveryCh         <-chan rabbitmq.Delivery
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

	deliveryCh, err := s.rabbitmqChannel.Consume(s.rabbitmqQueue.Name, "", true, false, false, false, nil)
	if err != nil {
		loggerx.Global().Error("rabbitmq Consume Refresh Msg failed", zap.Error(err))
		return
	}
	s.DeliveryCh = deliveryCh

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

		address.Address = message.Address
		dao.InitializeAddressStatus(ctx, address)
	}

	networks := []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
		protocol.NetworkArweave, protocol.NetworkXDAI, protocol.NetworkZkSync, protocol.NetworkCrossbell,
		protocol.NetworkEIP1577, protocol.NetworkFarcaster,
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

func (s *Service) SubscribeIndexerRefreshMessage(client *websocket.WSClient) {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case delivery := <-s.DeliveryCh:
			message := protocol.RefreshMessage{}
			if err := json.Unmarshal(delivery.Body, &message); err != nil {
				loggerx.Global().Error("failed to unmarshal message", zap.Error(err))
				continue
			}
			if _, ok := s.WsHub.Clients[client]; !ok {
				return
			}
			s.WsHub.Broadcast <- delivery.Body
		case <-ticker.C:
			if _, ok := s.WsHub.Clients[client]; !ok {
				return
			}
		}
	}
}
