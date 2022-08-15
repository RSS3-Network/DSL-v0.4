package aave

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressPoolV3Polygon = common.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD")

	EventHashSupplyV2   = common.BytesToHash(crypto.Keccak256([]byte("Deposit(address,address,address,uint256,uint16)")))
	EventHashBorrowV2   = common.BytesToHash(crypto.Keccak256([]byte("Borrow(address,address,address,uint256,uint256,uint256,uint16)")))
	EventHashRepayV2    = common.BytesToHash(crypto.Keccak256([]byte("Repay(address,address,address,uint256)")))
	EventHashWithdrawV2 = common.BytesToHash(crypto.Keccak256([]byte("Withdraw(address,address,address,uint256)")))

	EventHashSupplyV3   = common.BytesToHash(crypto.Keccak256([]byte("Supply(address,address,address,uint256,uint16)")))
	EventHashBorrowV3   = common.BytesToHash(crypto.Keccak256([]byte("Borrow(address,address,address,uint256,uint8,uint256,uint16)")))
	EventHashRepayV3    = common.BytesToHash(crypto.Keccak256([]byte("Repay(address,address,address,uint256,bool)")))
	EventHashWithdrawV3 = common.BytesToHash(crypto.Keccak256([]byte("Withdraw(address,address,address,uint256)")))
)
