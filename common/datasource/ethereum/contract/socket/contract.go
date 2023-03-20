package socket

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressRegistry = common.HexToAddress("0xc30141B657f4216252dc59Af2e7CdB9D8792e1B0")

	EventHashBridgeSend = crypto.Keccak256Hash([]byte("Send(bytes32,address,address,address,uint256,uint64,uint64,uint32)"))
)
