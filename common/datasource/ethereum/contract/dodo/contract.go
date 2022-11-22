package dodo

import (
	"github.com/ethereum/go-ethereum/crypto"
)

var EventHashDODOSwap = crypto.Keccak256Hash([]byte("DODOSwap(address,address,uint256,uint256,address,address)"))
