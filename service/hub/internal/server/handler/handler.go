package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	utils "github.com/naturalselectionlabs/pregod/common/utils/interface"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/service"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

type Handler struct {
	service *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) apiReport(path string, c echo.Context) {
	report := map[string]interface{}{
		"index":       model.EsIndex,
		"path":        path,
		"ts":          time.Now().Format("2006-01-02 15:04:05"),
		"count":       true,
		"api_key":     c.Get("API-KEY"),
		"remote_addr": c.RealIP(),
	}

	output, _ := json.Marshal(report)
	fmt.Printf("[DATABEAT]%v\n", string(output))
}

func (h *Handler) filterReport(path string, request interface{}) {
	b, _ := json.Marshal(request)
	report := make(map[string]interface{})

	err := json.Unmarshal(b, &report)
	if err != nil {
		return
	}

	for k, v := range report {
		if utils.IfInterfaceValueIsNil(v) {
			delete(report, k)
		}
	}

	if path == model.PostNotes {
		report["index"] = model.EsIndex
		report["path"] = path
		report["ts"] = time.Now().Format("2006-01-02 15:04:05")
		for _, address := range report["address"].([]interface{}) {
			report["address"] = address

			output, _ := json.Marshal(report)
			fmt.Printf("[DATABEAT]%v\n", string(output))
		}
	}
}

// publishIndexerMessage create a rabbitmq job to index the latest user data
func (h *Handler) publishIndexerMessage(ctx context.Context, message protocol.Message) {
	tracer := otel.Tracer("publishIndexerMessage")
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

		h.initializeIndexerStatus(ctx, message.Address)
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

			if err := h.RabbitmqChannel.Publish(protocol.ExchangeJob, protocol.IndexerWorkRoutingKey, false, false, rabbitmq.Publishing{
				ContentType: protocol.ContentTypeJSON,
				Body:        messageData,
			}); err != nil {
				return
			}
		}
	}()
}

func (h *Handler) initializeIndexerStatus(ctx context.Context, address string) {
	status := dbModel.Address{
		Address:          address,
		IndexingNetworks: protocol.SupportNetworks,
		Status:           false,
	}

	if err := database.Global().
		Model(dbModel.Address{}).
		Clauses(clause.OnConflict{
			UpdateAll: true,
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).
		Create(map[string]interface{}{
			"address":           status.Address,
			"status":            status.Status,
			"done_networks":     status.DoneNetworks,
			"indexing_networks": status.IndexingNetworks,
		}).Error; err != nil {
		loggerx.Global().Error("failed to upsert address", zap.Error(err), zap.String("address", address))
		return
	}
}
