package protocol

import (
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
)

type Message struct {
	Address     string    `json:"address"`
	Network     string    `json:"network"`
	Timestamp   time.Time `json:"timestamp"`
	BlockNumber int64     `json:"block_number"`
	IgnoreNote  bool      `json:"ignore_note"`
	Retry       int       `json:"retry"`
	IsUpdate    bool      `json:"is_update"`
}

type RefreshMessage struct {
	Address model.Address `json:"result"`
}
