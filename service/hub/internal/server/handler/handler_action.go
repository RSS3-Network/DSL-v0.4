package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/protocol"
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

	transfers, err := h.DatabaseClient.Transfer.Query().All(ctx)
	if err != nil {
		return err
	}

	spanDatabase.End()

	return c.JSON(http.StatusOK, transfers)
}
