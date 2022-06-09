package model

import (
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
)

type Feed struct {
	TransactionHash string           `json:"transaction_hash"`
	Timestamp       time.Time        `json:"timestamp"`
	Network         string           `json:"network"`
	Actions         []model.Transfer `json:"actions"`
}

type Feeds []Feed

func (f Feeds) Len() int {
	return len(f)
}

func (f Feeds) Less(i, j int) bool {
	return f[i].Timestamp.After(f[j].Timestamp)
}

func (f Feeds) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
