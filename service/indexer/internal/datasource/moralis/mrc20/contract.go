package mrc20

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate abigen --abi ./event.abi --pkg mrc20 --type MRC20 --out ./mrc20.go

var (
	LogTransferContractAddress = common.HexToAddress("0x0000000000000000000000000000000000001010")

	EventLogTransfer = common.BytesToHash(crypto.Keccak256([]byte("LogTransfer(address,address,address,uint256,uint256,uint256,uint256,uint256)")))
)
