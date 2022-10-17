package polygon

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressPredicate = common.HexToAddress("0x499a865ac595e6167482d2bd5A224876baB85ab4")
	AddressBridge    = common.HexToAddress("0xA0c68C638235ee32657e8f720a23ceC1bFc77C77")

	EventHashLockedEther = crypto.Keccak256Hash([]byte("LockedEther(address,address,uint256)"))
	EventHashExitedEther = crypto.Keccak256Hash([]byte("ExitedEther(address,uint256)"))

	PlatformBridge = "Polygon Bridge"
)
