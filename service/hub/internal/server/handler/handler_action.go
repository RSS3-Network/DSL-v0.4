package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	rabbitmq "github.com/rabbitmq/amqp091-go"
)

type GetActionListRequest struct {
	Address string `param:"address"`
}

func (h *Handler) GetActionListFunc(c echo.Context) error {
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

	if err := h.RabbitmqChannel.Publish(protocol.ExchangeJob, "", false, false, rabbitmq.Publishing{
		ContentType: protocol.ContentTypeJSON,
		Body:        messageData,
	}); err != nil {
		return err
	}

	return nil
}
