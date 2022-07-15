package weth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressEthereum = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	AddressPolygon  = common.HexToAddress("0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619")

	EventHashDeposit    = common.BytesToHash(crypto.Keccak256([]byte("Deposit(address,uint256)")))
	EventHashWithdrawal = common.BytesToHash(crypto.Keccak256([]byte("Withdrawal(address,uint256)")))
)
