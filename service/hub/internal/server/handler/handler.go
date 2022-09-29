package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	utils "github.com/naturalselectionlabs/pregod/common/utils/interface"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/websocket"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

const (
	DefaultLimit = 500

	// path
	GetNotes     = "/notes/"
	PostNotes    = "/notes"
	GetProfiles  = "/profiles/"
	PostProfiles = "/profiles"
	GetNS        = "/ns/"

	EsIndex = "pregod-v1-visit-path"
)

type Handler struct {
	RabbitmqConnection *rabbitmq.Connection
	RabbitmqChannel    *rabbitmq.Channel
	RabbitmqQueue      *rabbitmq.Queue
	WsHub              *websocket.WSHub
	HubId              string
}

type Response struct {
	Total         *int64          `json:"total,omitempty"`
	Cursor        string          `json:"cursor,omitempty"`
	Result        any             `json:"result,omitempty"`
	AddressStatus []model.Address `json:"address_status,omitempty"`
	Message       string          `json:"message,omitempty"`
}

type GetRequest struct {
	Address   string    `param:"address" json:"address" validate:"required"`
	Limit     int       `query:"limit" json:"limit"`
	Cursor    string    `query:"cursor" json:"cursor"`
	Type      []string  `query:"type" json:"type"`
	Tag       []string  `query:"tag" json:"tag" validate:"required_with=Type"`
	Network   []string  `query:"network" json:"network"`
	Platform  []string  `query:"platform" json:"platform"`
	Timestamp time.Time `query:"timestamp" json:"timestamp"`
	Hash      string    `query:"hash" json:"hash"`
	// includes POAP in the response
	IncludePoap bool `query:"include_poap" json:"include_poap"`
	Refresh     bool `query:"refresh" json:"refresh"`
	Reindex     bool `query:"reindex" json:"reindex"`
	Page        int  `query:"page" json:"page"`
	QueryStatus bool `query:"query_status" json:"query_status"`
	// returns a count of transactions only
	CountOnly bool   `query:"count_only" json:"count_only"`
	SocketId  string `query:"socket_id" json:"socket_id"`
}

type GetExchangeRequest struct {
	ExchangeType string   `param:"exchange_type"`
	Cursor       int      `query:"cursor"`
	Name         []string `query:"name"`
	Network      []string `query:"network"`
}

type GetPlatformRequest struct {
	PlatformType string   `param:"platform_type"`
	Network      []string `query:"network"`
}

type GetAssetRequest struct {
	Address      string   `param:"address" validate:"required"`
	Network      []string `query:"network"`
	TokenAddress string   `query:"token_address" validate:"required_with=TokenId"`
	TokenId      string   `query:"token_id"`
	Cursor       string   `query:"cursor"`
	Limit        int      `query:"limit"`
	Refresh      bool     `query:"refresh"`
	BlockSpam    *bool    `query:"block_spam"` // Default true
}

type BatchGetNotesRequest struct {
	Address     []string  `json:"address" validate:"required"`
	Type        []string  `json:"type"`
	Tag         []string  `json:"tag" validate:"required_with=Type"`
	Network     []string  `json:"network"`
	Platform    []string  `json:"platform"`
	Timestamp   time.Time `json:"timestamp"`
	Limit       int       `json:"limit"`
	Cursor      string    `json:"cursor"`
	Refresh     bool      `json:"refresh"`
	IncludePoap bool      `json:"include_poap"`
	Page        int       `json:"page"`
	QueryStatus bool      `json:"query_status"`
	CountOnly   bool      `json:"count_only"`
}

type BatchGetProfilesRequest struct {
	Address  []string `json:"address" validate:"required"`
	Network  []string `json:"network"`
	Platform []string `json:"platform"`
	Refresh  bool     `query:"refresh"`
}

type APIKeyRequest struct {
	Address string `json:"address" validate:"required"`
}

type Transactions []model.Transaction

// Len()
func (t Transactions) Len() int {
	return len(t)
}

// Less()
func (t Transactions) Less(i, j int) bool {
	return t[i].Timestamp.Unix() > t[j].Timestamp.Unix() ||
		(t[i].Timestamp.Unix() == t[j].Timestamp.Unix() && t[i].Index > t[j].Index)
}

// Swap()
func (t Transactions) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (h *Handler) apiReport(path string, c echo.Context) {
	report := map[string]interface{}{
		"index":       EsIndex,
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

	if path == PostNotes {
		report["index"] = EsIndex
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
	status := model.Address{
		Address:          address,
		IndexingNetworks: protocol.SupportNetworks,
		Status:           false,
	}

	if err := database.Global().
		Model(model.Address{}).
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
