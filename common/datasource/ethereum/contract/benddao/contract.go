package benddao

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressLendPool = common.HexToAddress("0x70b97A0da65C15dfb0FFA02aEE6FA36e507C2762")

	AddressBendDAO = common.HexToAddress("0x3b968d2d299b895a5fcf3bba7a64ad0f566e6f88")

	EventLendPoolDeposit  = crypto.Keccak256Hash([]byte("Deposit(address,address,uint256,address,uint16)"))
	EventLendPoolWithdraw = crypto.Keccak256Hash([]byte("Withdraw(address,address,uint256,address)"))
)
