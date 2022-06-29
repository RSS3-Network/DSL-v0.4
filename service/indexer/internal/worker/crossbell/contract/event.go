package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EventProfileCreated struct {
	ProfileID *big.Int
	Creator   common.Address
	To        common.Address

	EventProfileCreatedUnIndexed
}

type EventProfileCreatedUnIndexed struct {
	Handle    string
	Timestamp *big.Int
}
