package zerox

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var EventHashTransformedERC20 = common.BytesToHash(crypto.Keccak256([]byte("TransformedERC20(address,address,address,uint256,uint256)")))
