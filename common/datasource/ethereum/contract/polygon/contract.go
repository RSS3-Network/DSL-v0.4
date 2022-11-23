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
	AddressStakingValidatorShare  = common.HexToAddress("0x11cc04dD962e82D411587c56b815E8f8141Eb7D5")
	AddressStakingInfo            = common.HexToAddress("0xa59C847Bd5aC0172Ff4FE912C5d29E5A71A7512B")

	EventHashLockedEther             = crypto.Keccak256Hash([]byte("LockedEther(address,address,uint256)"))
	EventHashExitedEther             = crypto.Keccak256Hash([]byte("ExitedEther(address,uint256)"))
	EventHashLockedERC20             = crypto.Keccak256Hash([]byte("LockedERC20(address,address,address,uint256)"))
	EventHashLockedMintableERC20     = crypto.Keccak256Hash([]byte("LockedMintableERC20(address,address,address,uint256)"))
	EventHashShareMinted             = crypto.Keccak256Hash([]byte("ShareMinted(uint256,address,uint256,uint256)"))
	EventHashDelegatorClaimedRewards = crypto.Keccak256Hash([]byte("DelegatorClaimedRewards(uint256,address,uint256)"))

	PlatformBridge = "Polygon Bridge"
)
