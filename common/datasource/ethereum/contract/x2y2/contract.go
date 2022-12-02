package x2y2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressExchange = common.HexToAddress("0x74312363e45DCaBA76c59ec49a7Aa8A65a67EeD3")

	EventHashEvInventory = crypto.Keccak256Hash([]byte("EvInventory(bytes32,address,address,uint256,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes),(uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]))"))
)
