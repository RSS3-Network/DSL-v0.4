package nswap

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressExchange = common.HexToAddress("0x080fa1fb48e0b1bd251348efd02c1e7a12a931ac")

	EventHashEventMatch  = crypto.Keccak256Hash([]byte("EventMatch(bytes32,bytes32,address,address,uint256,uint256,(bytes4,address,uint256),(bytes4,address,uint256))"))
	EventHashEventCancel = crypto.Keccak256Hash([]byte("EventCancel(bytes32,address,(bytes4,address,uint256),(bytes4,address,uint256))"))
)
