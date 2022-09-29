package protocol

import (
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
)

type Message struct {
	Address       string    `json:"address"`
	Network       string    `json:"network"`
	Timestamp     time.Time `json:"timestamp"`
	BlockNumber   int64     `json:"block_number"`
	Reindex       bool      `json:"reindex"`
	IgnoreNote    bool      `json:"ignore_note"`
	IgnoreTrigger bool      `json:"ignore_trigger"`
	Retry         int       `json:"retry"`
	Refresh       bool      `json:"refresh"`
	HubId         string    `json:"hub_id"`
	SocketId      string    `json:"socket_id"`
}

type RefreshMessage struct {
	SocketId string        `json:"socket_id"`
	Address  model.Address `json:"result"`
}
