package protocol

import (
	"time"
)

type Message struct {
	Address       string    `json:"address"`
	Network       string    `json:"network"`
	Timestamp     time.Time `json:"timestamp"`
	BlockNumber   int64     `json:"block_number"`
	Reindex       bool      `json:"reindex"`
	IgnoreNote    bool      `json:"ignore_note"`
	IgnoreTrigger bool      `json:"ignore_trigger"`
}
