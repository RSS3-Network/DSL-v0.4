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
	Reindex     bool      `json:"reindex"`
	IgnoreNote  bool      `json:"ignore_note"`
	Retry       int       `json:"retry"`
}

type RefreshMessage struct {
	Status     bool          `json:"status"`
	FinishedAt time.Time     `json:"finished_at,omitempty"` // 完成时间
	UpdatedAt  time.Time     `json:"updated_at,omitempty"`  // 本次最早的更新时间
	Address    model.Address `json:"result"`
}
