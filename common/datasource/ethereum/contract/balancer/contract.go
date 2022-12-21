package balancer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressVault = common.HexToAddress("0xBA12222222228d8Ba445958a75a0704d566BF2C8")

	EventHashSwap           = crypto.Keccak256Hash([]byte("Swap(bytes32,address,address,uint256,uint256)"))
	EventPoolBalanceChanged = crypto.Keccak256Hash([]byte("PoolBalanceChanged(bytes32,address,address[],int256[],uint256[])"))
)
