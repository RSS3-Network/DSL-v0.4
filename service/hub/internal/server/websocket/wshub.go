package websocket

import (
	"encoding/json"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.uber.org/zap"
	"golang.org/x/exp/maps"
)

type WSHub struct {
	Clients    map[*WSClient]string
	Action     chan map[string][]byte
	Broadcast  chan []byte
	Register   chan *WSClient
	Unregister chan *WSClient
}

func NewHub() *WSHub {
	return &WSHub{
		Action:     make(chan map[string][]byte),
		Broadcast:  make(chan []byte),
		Register:   make(chan *WSClient),
		Unregister: make(chan *WSClient),
		Clients:    make(map[*WSClient]string),
	}
}

func (h *WSHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = string(client.ClientId)
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
		case data := <-h.Action:
			for client := range h.Clients {
				if message, ok := data[string(client.ClientId)]; ok {
					request := model.WebSocketRequest{}
					var response []byte
					if err := json.Unmarshal(message, &request); err != nil || request.Action == "" || request.Id == nil {
						loggerx.Global().Error("failed to unmarshal websocket message", zap.Error(err))
						response, _ = json.Marshal(model.WebsocketResponse{Status: "error", Result: map[string]any{"msg": "failed to unmarshal websocket message"}})
						select {
						case client.Send <- response:
						default:
							close(client.Send)
							delete(h.Clients, client)
							delete(GetClientMaps(), string(client.ClientId))
						}
					}

					switch request.Action {
					case model.Subscribe:
						for _, address := range request.AddressArr {
							client.AddressMap[strings.ToLower(address)] = struct{}{}
						}
						response, _ = json.Marshal(model.WebsocketResponse{Id: *request.Id, Status: "success", Result: map[string]any{"msg": request.Action, "address": maps.Keys(client.AddressMap)}})
						select {
						case client.Send <- response:
						default:
							close(client.Send)
							delete(h.Clients, client)
							delete(GetClientMaps(), string(client.ClientId))
						}
					case model.Unsubscribe:
						for _, address := range request.AddressArr {
							delete(client.AddressMap, strings.ToLower(address))
						}
						response, _ = json.Marshal(model.WebsocketResponse{Id: *request.Id, Status: "success", Result: map[string]any{"msg": request.Action, "address": maps.Keys(client.AddressMap)}})
						select {
						case client.Send <- response:
						default:
							close(client.Send)
							delete(h.Clients, client)
							delete(GetClientMaps(), string(client.ClientId))
						}
					case model.Query:
						response, _ = json.Marshal(model.WebsocketResponse{Id: *request.Id, Status: "success", Result: map[string]any{"msg": request.Action, "address": maps.Keys(client.AddressMap)}})
						select {
						case client.Send <- response:
						default:
							close(client.Send)
							delete(h.Clients, client)
							delete(GetClientMaps(), string(client.ClientId))
						}
					default:
						response, _ = json.Marshal(model.WebsocketResponse{Id: *request.Id, Status: "error", Result: map[string]any{"msg": "unsupport action: " + request.Action}})
						select {
						case client.Send <- response:
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
}
