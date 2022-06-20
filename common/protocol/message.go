package protocol

import "time"

type Message struct {
	Address   string    `json:"address"`
	Network   string    `json:"network"`
	Timestamp time.Time `json:"timestamp"`
}
