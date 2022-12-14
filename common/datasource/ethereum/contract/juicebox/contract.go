package juicebox

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressETHPaymentTerminal               = common.HexToAddress("0x594Cb208b5BB48db1bcbC9354d1694998864ec63")
	AddressProjects                         = common.HexToAddress("0xD8B4359143eda5B2d763E127Ed27c77addBc47d3")
	AddressTiered721DelegateProjectDeployer = common.HexToAddress("0x8e9c1b5b4f1b5b4f1b5b4f1b5b4f1b5b4f1b5b4f")
	AddressController                       = common.HexToAddress("0xFFdD70C318915879d5192e8a0dcbFcB0285b3C98")

	EventLaunchProject = crypto.Keccak256Hash([]byte("LaunchProject(uint256,uint256,string,address)"))
	EventPay           = crypto.Keccak256Hash([]byte("Pay(uint256,uint256,uint256,address,address,uint256,uint256,string,bytes,address)"))
)
