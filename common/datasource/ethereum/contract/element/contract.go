package element

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressExchange    = common.HexToAddress("0xEAF5453b329Eb38Be159a872a6ce91c9A8fb0260")
	AddressNativeToken = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

	EventERC721SellOrderFilled  = crypto.Keccak256Hash([]byte("ERC721SellOrderFilled(bytes32,address,address,uint256,address,uint256,(address,uint256)[],address,uint256)"))
	EventERC1155SellOrderFilled = crypto.Keccak256Hash([]byte("ERC1155SellOrderFilled(bytes32,address,address,uint256,address,uint256,(address,uint256)[],address,uint256,uint128)"))
)
