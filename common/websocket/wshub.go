package websocket

import (
	"encoding/json"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"
)

type WSHub struct {
	// Registered clients.
	Clients map[*WSClient]bool
	// Inbound messages from the clients.
	Broadcast chan []byte
	// Register requests from the clients.
	Register chan *WSClient
	// Unregister requests from clients.
	Unregister chan *WSClient
	// unique id key:client value: unique
	ClientId map[*WSClient]string
}

func NewHub() *WSHub {
	return &WSHub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *WSClient),
		Unregister: make(chan *WSClient),
		Clients:    make(map[*WSClient]bool),
		ClientId:   make(map[*WSClient]string),
	}
}

func (h *WSHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.ClientId[client] = string(client.ClientId)
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				delete(h.ClientId, client)
				close(client.Send)
			}
		case data := <-h.Broadcast:
			message := protocol.RefreshMessage{}
			if err := json.Unmarshal(data, &message); err != nil {
				loggerx.Global().Error("failed to unmarshal message", zap.Error(err))
				continue
			}
			for client := range h.Clients {
				if string(client.ClientId) == message.SocketId {
					select {
					case client.Send <- data:
					default:
						close(client.Send)
						delete(h.Clients, client)
						delete(h.ClientId, client)
					}
				}
			}
		}
	}
}
