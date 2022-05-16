package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/response"
	rabbitmq "github.com/rabbitmq/amqp091-go"
)

type GetActionListRequest struct {
	Address string `param:"address"`
}

func (h *Handler) GetActionListFunc(c echo.Context) error {
	tracer := h.TracerProvider.Tracer("handler_get_action_list")

	ctx, spanHTTPHandler := tracer.Start(c.Request().Context(), "http")
	defer spanHTTPHandler.End()

	request := GetActionListRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}

	message := protocol.Message{
		Address: request.Address,
		Network: protocol.NetworkEthereum,
	}

	messageData, err := json.Marshal(&message)
	if err != nil {
		return err
	}

	ctx, snapRabbitMQ := tracer.Start(ctx, "rabbitmq")
	if err := h.RabbitmqChannel.Publish(protocol.ExchangeJob, "", false, false, rabbitmq.Publishing{
		ContentType: protocol.ContentTypeJSON,
		Body:        messageData,
	}); err != nil {
		return err
	}

	snapRabbitMQ.End()

	ctx, spanDatabase := tracer.Start(ctx, "postgres")

	transfers := make([]model.Transfer, 0)
	if err := h.DatabaseClient.
		Model((*model.Transfer)(nil)).
		Where("address_from = ? OR address_to = ?", request.Address, request.Address).
		Find(&transfers).Error; err != nil {
		return err
	}

	spanDatabase.End()

	feeds := make([]response.Feed, 0)

	for _, transfer := range transfers {
		feeds = append(feeds, response.Feed{
			Tags:    []string{"transfer"},
			Network: transfer.Network,
			Proof:   transfer.TransactionHash,
			Actions: []response.Action{
				{
					Type:          "transfer",
					From:          transfer.AddressFrom,
					To:            transfer.AddressTo,
					Token:         "",
					TokenAddress:  transfer.TokenAddress,
					TokenStandard: transfer.TokenStandard,
					TokenID:       transfer.TokenID,
					TokenValue:    transfer.TokenValue,
				},
			},
		})
	}

	return c.JSON(http.StatusOK, &response.Response{
		Total:  int64(len(feeds)),
		Result: feeds,
	})
}
