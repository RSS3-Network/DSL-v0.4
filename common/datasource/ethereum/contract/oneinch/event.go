package oneinch

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	EventHashSwapped = common.BytesToHash(crypto.Keccak256([]byte("Swapped(address,address,address,address,uint256,uint256)")))
)
