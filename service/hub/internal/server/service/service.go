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
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type Service struct {
	dao                *dao.Dao
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
}

func New(d *dao.Dao) (s *Service) {
	s = &Service{
		dao: d,
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
		loggerx.Global().Error("rabbitmqChannel ExchangeDeclare failed", zap.Error(err))
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

		s.dao.InitializeAddressStatus(ctx, message.Address)
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
