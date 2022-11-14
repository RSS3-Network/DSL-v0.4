package websocket

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"
)

const (
	waitTime = 10 * time.Second
)

var globalLocker sync.RWMutex

var clientMaps map[string]*WSClient

var once sync.Once

func ReplaceGlobal(clientId string, client *WSClient) {
	globalLocker.Lock()

	defer globalLocker.Unlock()

	clientMaps[clientId] = client
}

func GetClientMaps() map[string]*WSClient {
	once.Do(func() {
		clientMaps = make(map[string]*WSClient)
	})

	globalLocker.RLock()

	defer globalLocker.RUnlock()

	item := make(map[string]*WSClient)
	for key, value := range clientMaps {
		item[key] = value
	}

	return item
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
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				loggerx.Global().Info("websocket is disconnected", zap.String("websocket_id", string(c.ClientId)))
			}
			break
		}

		c.Hub.Action <- map[string][]byte{string(c.ClientId): message}
	}
}
