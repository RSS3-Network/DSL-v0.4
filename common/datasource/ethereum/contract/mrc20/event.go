package mrc20

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	EventHashTransfer    = common.BytesToHash(crypto.Keccak256([]byte("Transfer(address,address,uint256)")))
	EventHashApproval    = common.BytesToHash(crypto.Keccak256([]byte("Approval(address,address,uint256)")))
	EventHashLogTransfer = common.BytesToHash(crypto.Keccak256([]byte("LogTransfer(address,address,address,uint256,uint256,uint256,uint256,uint256)")))
)
