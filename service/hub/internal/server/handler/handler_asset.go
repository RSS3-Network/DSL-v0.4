package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

var cursorKey = "%v:%v:%v"

// GetAssetsFunc HTTP handler for asset API
func (h *Handler) GetAssetsFunc(c echo.Context) error {
	tracer := otel.Tracer("GetAssetsFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetAssetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	if request.Limit <= 0 || request.Limit > model.DefaultLimit {
		request.Limit = model.DefaultLimit
	}

	request.Address = strings.ToLower(request.Address)
	assetList, total, err := h.getAssets(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	// publish mq message
	if request.Refresh || len(assetList) == 0 {
		go h.publishIndexerAssetMessage(ctx, request.Address)
	}

	if len(assetList) == 0 {
		return c.JSON(http.StatusOK, &model.Response{
			Result: make([]dbModel.Asset, 0),
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		last := assetList[len(assetList)-1]
		cursor = fmt.Sprintf(cursorKey, last.Network, last.TokenAddress, last.TokenID)
	}

	return c.JSON(http.StatusOK, &model.Response{
		Total:  &total,
		Cursor: cursor,
		Result: assetList,
	})
}

func (h *Handler) getAssets(c context.Context, request model.GetAssetRequest) ([]dbModel.Asset, int64, error) {
	tracer := otel.Tracer("getAssets")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	assetList := make([]dbModel.Asset, 0)
	total := int64(0)
	sql := database.Global().
		Model(&dbModel.Asset{}).
		Where("owner = ?", request.Address)

	if request.BlockSpam == nil || *request.BlockSpam {
		sql = sql.Where("token_address NOT IN ?", allowlist.SpamList.Keys())
	}

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.TokenAddress) > 0 {
		sql = sql.Where("token_address = ?", strings.ToLower(request.TokenAddress))
	}

	if len(request.TokenId) > 0 {
		sql = sql.Where("token_id = ?", request.TokenId)
	}

	if len(request.Cursor) > 0 {
		param := strings.Split(request.Cursor, ":")
		if len(param) < 3 {
			return nil, 0, fmt.Errorf("cursor is wrong")
		}

		var lastItem dbModel.Asset

		if err := database.Global().
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

		if err := svc.rabbitmqChannel.Publish(protocol.ExchangeJob, protocol.IndexerAssetRoutingKey, false, false, rabbitmq.Publishing{
			ContentType: protocol.ContentTypeJSON,
			Body:        messageData,
		}); err != nil {
			return
		}
	}
}
