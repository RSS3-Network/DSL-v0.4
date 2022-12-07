package quix

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressWrapperSeaportProxy = common.HexToAddress("0xC78A09D6a4badecc7614A339FD264B7290361ef1")
	AddressSeaport             = common.HexToAddress("0x998EF16Ea4111094EB5eE72fC2c6f4e6E8647666")
	AddressExchangeV5          = common.HexToAddress("0x3F9DA045b0F77d707eA4061110339C4Ea8ecFA70")

	EventHashOrderFulfilled  = crypto.Keccak256Hash([]byte("OrderFulfilled(bytes32,address,address,address,(uint8,address,uint256,uint256)[],(uint8,address,uint256,uint256,address)[])"))
	EventHashSellOrderFilled = crypto.Keccak256Hash([]byte("SellOrderFilled(address,address,address,uint256,uint256)"))
)
