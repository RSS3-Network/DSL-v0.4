package lido

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressETH   = common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84")
	AddressMatic = common.HexToAddress("0x9ee91f9f426fa633d227f7a9b000e28b9dfd8599")

	EventHashSubmitted   = crypto.Keccak256Hash([]byte("Submitted(address,uint256,address)"))
	EventHashSubmitEvent = crypto.Keccak256Hash([]byte("SubmitEvent(address,uint256)"))
)
