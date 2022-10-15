package websocket

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.uber.org/zap"
	"golang.org/x/exp/maps"
)

const (
	waitTime = 10 * time.Second
)

var ClientMaps map[string]*WSClient

var once sync.Once

func GetClientMaps() map[string]*WSClient {
	once.Do(func() {
		ClientMaps = make(map[string]*WSClient)
	})
	return ClientMaps
}

type WSClient struct {
	Hub        *WSHub
	Conn       *websocket.Conn
	Send       chan []byte
	ClientId   []byte
	AddressMap map[string]struct{}
}

func (c *WSClient) WriteMsg() {
	ticker := time.NewTicker(waitTime)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					return
				}
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			if _, err := w.Write(message); err != nil {
				return
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *WSClient) ReadMsg() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			return
		}
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				loggerx.Global().Info("websocket is disconnected", zap.String("websocket_id", string(c.ClientId)))
			}
			break
		}

		request := model.WebSocketRequest{}

		if err := json.Unmarshal(message, &request); err != nil || request.Action == "" || request.Id == nil {
			loggerx.Global().Error("failed to unmarshal websocket message", zap.Error(err))
			if err = c.Conn.WriteJSON(model.WebsocketResponse{Status: "error", Result: map[string]any{"msg": "failed to unmarshal websocket message"}}); err != nil {
				return
			}
			continue
		}

		switch request.Action {
		case model.Subscribe:
			for _, address := range request.AddressArr {
				c.AddressMap[strings.ToLower(address)] = struct{}{}
			}
			if err = c.Conn.WriteJSON(model.WebsocketResponse{Id: *request.Id, Status: "success", Result: map[string]any{"msg": request.Action, "address": maps.Keys(c.AddressMap)}}); err != nil {
				return
			}

		case model.Unsubscribe:
			for _, address := range request.AddressArr {
				delete(c.AddressMap, strings.ToLower(address))
			}
			if err = c.Conn.WriteJSON(model.WebsocketResponse{Id: *request.Id, Status: "success", Result: map[string]any{"msg": request.Action, "address": maps.Keys(c.AddressMap)}}); err != nil {
				return
			}

		case model.Query:
			if err = c.Conn.WriteJSON(model.WebsocketResponse{Id: *request.Id, Status: "success", Result: map[string]any{"msg": request.Action, "address": maps.Keys(c.AddressMap)}}); err != nil {
				return
			}
		default:
			if err = c.Conn.WriteJSON(model.WebsocketResponse{Id: *request.Id, Status: "error", Result: map[string]any{"msg": "unsupport action: " + request.Action}}); err != nil {
				return
			}
		}
	}
}
