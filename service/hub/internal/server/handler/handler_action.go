package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/response"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type GetActionListRequest struct {
	Address string `param:"address"`
}

func (h *Handler) GetActionListFunc(c echo.Context) error {
	tracer := otel.Tracer("handler_get_action_list")

	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetActionListRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}

	ctx, rabbitmqSnap := tracer.Start(ctx, "rabbitmq")

	networks := []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}

	for _, network := range networks {
		message := protocol.Message{
			Address: request.Address,
			Network: network,
		}

		messageData, err := json.Marshal(&message)
		if err != nil {
			return err
		}

		if err := h.RabbitmqChannel.Publish(protocol.ExchangeJob, "", false, false, rabbitmq.Publishing{
			ContentType: protocol.ContentTypeJSON,
			Body:        messageData,
		}); err != nil {
			return err
		}
	}

	rabbitmqSnap.End()

	ctx, postgresSnap := tracer.Start(ctx, "postgres")

	transfers := make([]model.Transfer, 0)
	if err := h.DatabaseClient.
		Model((*model.Transfer)(nil)).
		Where("address_from = ? OR address_to = ?", strings.ToLower(request.Address), strings.ToLower(request.Address)).
		Order("timestamp DESC").
		Find(&transfers).Error; err != nil {
		return err
	}

	// Hash map is unordered
	transferMap := make(map[string][]model.Transfer)
	for _, transfer := range transfers {
		transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
	}

	postgresSnap.End()

	feeds := make([]response.Feed, 0)

	for _, transfers := range transferMap {
		firstTransfer := transfers[0]

		feed := response.Feed{
			Timestamp: firstTransfer.Timestamp,
			Network:   firstTransfer.Network,
			Proof:     firstTransfer.TransactionHash,
		}

		for _, transfer := range transfers {
			if strings.ContainsAny(transfer.Type, "swap_") {
				metadataModel := metadata.Metadata{}

				if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
					return err
				}

				feed.Tag = "swap"
				feed.Platform = metadataModel.Swap.Name
			}

			feed.Actions = append(feed.Actions, response.Action{
				Type:     transfer.Type,
				From:     transfer.AddressFrom,
				To:       transfer.AddressTo,
				Metadata: transfer.Metadata,
			})
		}

		feeds = append(feeds, feed)
	}

	sort.Sort(response.Feeds(feeds))

	return c.JSON(http.StatusOK, &response.Response{
		Total:  int64(len(feeds)),
		Result: feeds,
	})
}
