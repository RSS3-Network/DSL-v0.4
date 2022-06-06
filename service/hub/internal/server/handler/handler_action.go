package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	m "github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type GetActionListRequest struct {
	Address     string   `param:"address"`
	Limit       int      `query:"limit"`
	Cursor      string   `query:"cursor"`
	Types       []string `query:"types"`
	Tags        []string `query:"tags"`
	ExculdeTags []string `query:"exclude_tags"`
	ItemSources []string `query:"item_sources"`
	Networks    []string `query:"networks"`
}

// GetActionListFunc HTTP handler for action API
// parse query parameters, query and assemble data
func (h *Handler) GetActionListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetActionListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetActionListRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}

	if request.Limit <= 0 || request.Limit > DefaultLimit {
		request.Limit = DefaultLimit
	}

	go func() {
		// create a rabbitmq job to index the latest user data
		_, rabbitmqSnap := tracer.Start(ctx, "rabbitmq")

		defer rabbitmqSnap.End()

		networks := []string{
			protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain, protocol.NetworkArweave,
		}

		for _, network := range networks {
			message := protocol.Message{
				Address: request.Address,
				Network: network,
			}

			messageData, err := json.Marshal(&message)
			if err != nil {
				return
			}

			if err := h.RabbitmqChannel.Publish(protocol.ExchangeJob, "", false, false, rabbitmq.Publishing{
				ContentType: protocol.ContentTypeJSON,
				Body:        messageData,
			}); err != nil {
				return
			}
		}
	}()

	transfers, total, err := h.getActionListDatabase(ctx, request)
	if err != nil {
		return err
	}

	transferMap := make(map[string][]model.Transfer)
	for _, transfer := range transfers {
		transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
	}

	result := make([]m.Feed, 0)

	for transactionHash, transfers := range transferMap {
		feed := m.Feed{
			TransactionHash: transactionHash,
			Actions:         transfers,
		}

		result = append(result, feed)
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  total,
		Cursor: result[len(result)-1].TransactionHash,
		Result: result,
	})
}

// getActionListDatabase get transfer data from database
func (h *Handler) getActionListDatabase(c context.Context, request GetActionListRequest) ([]model.Transfer, int64, error) {
	tracer := otel.Tracer("getActionListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transfers := make([]model.Transfer, 0)
	transactionHashList := []string{}
	total := int64(0)
	sql := h.DatabaseClient.
		Model(&model.Transfer{}).
		Select("transaction_hash").
		Where("LOWER(address_from) = ? OR LOWER(address_to) = ? OR (network = 'arweave' AND LOWER(metadata #>> '{mirror, contributor}') = ?)",
			strings.ToLower(request.Address),
			strings.ToLower(request.Address),
			strings.ToLower(request.Address),
		)

	if len(request.Cursor) > 0 {
		sql = sql.Where("transaction_hash < ?", request.Cursor)
	}

	if len(request.Types) > 0 {
		sql = sql.Where("\"type\" in ? ", request.Types)
	}

	if len(request.Tags) > 0 {
		sql = sql.Where("tags && ?", pq.StringArray(request.Tags))
	}

	if len(request.ExculdeTags) > 0 {
		sql = sql.Where("tags && ? = FALSE", pq.StringArray(request.ExculdeTags))
	}

	if len(request.ItemSources) > 0 {
		sql = sql.Where("source IN ?", request.ItemSources)
	}

	if len(request.Networks) > 0 {
		sql = sql.Where("network IN ?", request.Networks)
	}

	if err := sql.Group("transaction_hash").Limit(request.Limit).Find(&transactionHashList).
		Limit(-1).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := h.DatabaseClient.Model(&model.Transfer{}).
		Where("transaction_hash IN ?", transactionHashList).
		Find(&transfers).Error; err != nil {
		return nil, 0, err
	}

	return transfers, total, nil
}
