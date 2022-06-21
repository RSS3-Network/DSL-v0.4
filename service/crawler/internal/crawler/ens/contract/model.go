package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var TopicHashNameRegistered = common.HexToHash("0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f")

type NameRegisteredData struct {
	Name    string
	Cost    *big.Int
	Expires *big.Int
}
