package model

import (
	"github.com/naturalselectionlabs/pregod/common/database/model"
)

type Feed struct {
	TransactionHash string           `json:"transaction_hash"`
	Actions         []model.Transfer `json:"actions"`
}
