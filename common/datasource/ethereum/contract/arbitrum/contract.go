package arbitrum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressInboxOne  = common.HexToAddress("0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f")
	AddressInboxNova = common.HexToAddress("0xc4448b71118c9071Bcb9734A0EAc55D18A153949")

	EventHashMessageDelivered = crypto.Keccak256Hash([]byte("MessageDelivered(uint256,bytes32,address,uint8,address,bytes32,uint256,uint64)"))

	PlatformInboxOne  = "Arbitrum Inbox One"
	PlatformInboxNova = "Arbitrum Inbox Nova"
)
