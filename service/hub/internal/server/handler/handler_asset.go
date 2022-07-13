package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

var cursorKey = "%v:%v:%v"

// GetAssetListFunc HTTP handler for asset API
func (h *Handler) GetAssetListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetAssetListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	request.Address = strings.ToLower(request.Address)

	if request.Limit <= 0 || request.Limit > DefaultLimit {
		request.Limit = DefaultLimit
	}

	assetList, total, err := h.getAssetListDatabase(ctx, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	// publish mq message
	if request.Refresh || len(assetList) == 0 {
		go h.publishIndexerAssetMessage(ctx, request.Address)
	}

	if len(assetList) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]dbModel.Transaction, 0),
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		last := assetList[len(assetList)-1]
		cursor = fmt.Sprintf(cursorKey, last.Network, last.TokenAddress, last.TokenID)
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  total,
		Cursor: cursor,
		Result: assetList,
	})
}

func (h *Handler) getAssetListDatabase(c context.Context, request GetRequest) ([]dbModel.Asset, int64, error) {
	tracer := otel.Tracer("getAssetListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	assetList := make([]dbModel.Asset, 0)
	total := int64(0)
	sql := h.DatabaseClient.
		Model(&dbModel.Asset{}).
		Where("owner = ?", request.Address)

	if len(request.Cursor) > 0 {
		param := strings.Split(request.Cursor, ":")
		if len(param) < 3 {
			return nil, 0, fmt.Errorf("cursor is wrong")
		}

		var lastItem dbModel.Asset

		if err := h.DatabaseClient.
			Where("network = ?", param[0]).
			Where("token_address = ?", param[1]).
			Where("token_id", param[2]).
			First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND token_id < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.TokenID)
	}

	if err := sql.Count(&total).Limit(request.Limit).Order("timestamp DESC, token_id DESC").Find(&assetList).Error; err != nil {
		return nil, 0, err
	}

	return assetList, total, nil
}

// publishIndexerAssetMessage create a rabbitmq job to index the latest user data
func (h *Handler) publishIndexerAssetMessage(ctx context.Context, address string) {
	tracer := otel.Tracer("publishIndexerAssetMessage")
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

		if err := h.RabbitmqChannel.Publish(protocol.ExchangeJob, protocol.IndexerAssetRoutingKey, false, false, rabbitmq.Publishing{
			ContentType: protocol.ContentTypeJSON,
			Body:        messageData,
		}); err != nil {
			return
		}
	}
}
