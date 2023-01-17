package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/constant"
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

	if err = s.rabbitmqChannel.ExchangeDeclare(protocol.ExchangeJob, "direct", true, false, false, false, nil); err != nil {
		loggerx.Global().Error("rabbitmqChannel ExchangeDeclare failed", zap.Error(err))
		return err
	}

	if err = s.rabbitmqChannel.ExchangeDeclare(protocol.ExchangeRefresh, "fanout", true, false, false, false, nil); err != nil {
		loggerx.Global().Error("rabbitmqChannel ExchangeDeclare Refresh failed", zap.Error(err))
		return err
	}

	if s.rabbitmqQueue, err = s.rabbitmqChannel.QueueDeclare(
		"", false, false, true, false, nil,
	); err != nil {
		loggerx.Global().Error("rabbitmq QueueDeclare failed", zap.Error(err))
		return err
	}

	if err = s.rabbitmqChannel.QueueBind(
		s.rabbitmqQueue.Name, "", protocol.ExchangeRefresh, false, nil,
	); err != nil {
		loggerx.Global().Error("rabbitmq QueueBind failed", zap.Error(err))
		return err
	}

	deliveryCh, err := s.rabbitmqChannel.Consume(s.rabbitmqQueue.Name, "", true, false, false, false, nil)
	if err != nil {
		loggerx.Global().Error("rabbitmq Consume Refresh Msg failed", zap.Error(err))
		return err
	}
	s.DeliveryCh = deliveryCh

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
		protocol.NetworkArweave,
		protocol.NetworkXDAI, protocol.NetworkZkSync, protocol.NetworkCrossbell,
		protocol.NetworkOptimism, protocol.NetworkAvalanche, protocol.NetworkCelo, protocol.NetworkFantom,
		protocol.NetworkEIP1577,
	}

	routingKey := getRoutingKeyByHeader(ctx)

	for _, network := range networks {
		message.Network = network

		messageData, err := json.Marshal(&message)
		if err != nil {
			return
		}

		if err := s.rabbitmqChannel.Publish(protocol.ExchangeJob, routingKey, false, false, rabbitmq.Publishing{
			ContentType: protocol.ContentTypeJSON,
			Body:        messageData,
		}); err != nil {
			loggerx.Global().Error("publish indexer message failed", zap.Error(err), zap.String("address", message.Address), zap.String("network", network))
			return
		}
	}
}

// if not valid, return empty string
func getRoutingKeyByHeader(ctx context.Context) string {
	var (
		httpHeader http.Header
		ok         bool
	)

	// default
	res := protocol.IndexerWorkRoutingKey

	header := ctx.Value(constant.HEADER_CTX_KEY)
	if header == nil {
		return res
	}
	if httpHeader, ok = header.(http.Header); !ok {
		return res
	}

	// check origin
	// TODO to be removed
	for _, domain := range []string{"rss3.io", "hoot.it", "openscan.it"} {
		if strings.Contains(httpHeader.Get("Origin"), domain) {
			return protocol.IndexerWorkRoutingKeyIO
		}
	}

	// check api key
	apiKey := httpHeader.Get(constant.API_KEY_HEADER)
	if apiKey == constant.IO_API_Key {
		return protocol.IndexerWorkRoutingKeyIO
	}

	return res
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

func (s *Service) SubscribeIndexerRefreshMessage(client *websocket.WSClient) {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case delivery := <-s.DeliveryCh:
			if len(delivery.Body) == 0 {
				loggerx.Global().Info("wait mq reconnected")
				time.Sleep(3 * time.Second)
				continue
			}
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
