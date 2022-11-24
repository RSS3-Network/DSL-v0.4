package polygon

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressBridge                 = common.HexToAddress("0xA0c68C638235ee32657e8f720a23ceC1bFc77C77")
	AddressEtherePredicate        = common.HexToAddress("0x499a865ac595e6167482d2bd5A224876baB85ab4")
	AddressERC20Predicate         = common.HexToAddress("0x40ec5B33f54e0E8A33A975908C5BA1c14e5BbbDf")
	AddressMintableERC20Predicate = common.HexToAddress("0x9923263fA127b3d1484cFD649df8f1831c2A74e4")

	EventHashLockedEther         = crypto.Keccak256Hash([]byte("LockedEther(address,address,uint256)"))
	EventHashExitedEther         = crypto.Keccak256Hash([]byte("ExitedEther(address,uint256)"))
	EventHashLockedERC20         = crypto.Keccak256Hash([]byte("LockedERC20(address,address,address,uint256)"))
	EventHashLockedMintableERC20 = crypto.Keccak256Hash([]byte("LockedMintableERC20(address,address,address,uint256)"))

	PlatformBridge = "Polygon Bridge"
)
