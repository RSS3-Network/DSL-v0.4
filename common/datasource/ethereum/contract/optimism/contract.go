package optimism

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressGateway = common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1")

	EventHashETHDepositInitiated      = crypto.Keccak256Hash([]byte("ETHDepositInitiated(address,address,uint256,bytes)"))
	EventHashETHWithdrawalFinalized   = crypto.Keccak256Hash([]byte("ETHWithdrawalFinalized(address,address,uint256,bytes)"))
	EventHashERC20DepositInitiated    = crypto.Keccak256Hash([]byte("ERC20DepositInitiated(address,address,address,address,uint256,bytes)"))
	EventHashERC20WithdrawalFinalized = crypto.Keccak256Hash([]byte("ERC20WithdrawalFinalized(address,address,address,address,uint256,bytes)"))

	PlatformBridge = "Optimism Bridge"
)
