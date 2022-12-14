package juicebox

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressETHPaymentTerminal = common.HexToAddress("0x594Cb208b5BB48db1bcbC9354d1694998864ec63")
	AddressProjects           = common.HexToAddress("0xD8B4359143eda5B2d763E127Ed27c77addBc47d3")

	EventMint = crypto.Keccak256Hash([]byte("Mint(uint256,uint256,address,uint256,address)"))
	EventPay  = crypto.Keccak256Hash([]byte("Pay(uint256,uint256,uint256,address,address,uint256,uint256,string,bytes,address)"))
)
