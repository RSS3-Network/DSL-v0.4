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
	// a list of EVM networks where the address is NOT deployed as a contract
	ValidEVMNetworkList []string
}

type RefreshMessage struct {
	Address model.Address `json:"result"`
}
