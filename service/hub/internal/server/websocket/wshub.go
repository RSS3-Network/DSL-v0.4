package websocket

import (
	"encoding/json"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"
)

type WSHub struct {
	Clients    map[*WSClient]bool
	Broadcast  chan []byte
	Register   chan *WSClient
	Unregister chan *WSClient
}

func NewHub() *WSHub {
	return &WSHub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *WSClient),
		Unregister: make(chan *WSClient),
		Clients:    make(map[*WSClient]bool),
	}
}

func (h *WSHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				delete(GetClientMaps(), string(client.ClientId))
				close(client.Send)
			}
		case data := <-h.Broadcast:
			message := protocol.RefreshMessage{}
			if err := json.Unmarshal(data, &message); err != nil {
				loggerx.Global().Error("failed to unmarshal message", zap.Error(err))
				continue
			}
			for client := range h.Clients {
				if _, ok := client.AddressMap[message.Address.Address]; ok {
					select {
					case client.Send <- data:
					default:
						close(client.Send)
						delete(h.Clients, client)
						delete(GetClientMaps(), string(client.ClientId))
					}
				}
			}
		}
	}
}
