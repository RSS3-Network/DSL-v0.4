package uniswap

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	EventHashTransfer = common.BytesToHash(crypto.Keccak256([]byte("Transfer(address,address,uint256)")))
	EventHashSwap     = common.BytesToHash(crypto.Keccak256([]byte("Swap(address,address,int256,int256,uint160,uint128,int24)")))
	EventHashApproval = common.BytesToHash(crypto.Keccak256([]byte("Approval(address,address,uint256)")))
)
