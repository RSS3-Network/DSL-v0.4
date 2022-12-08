package zerox

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressExchangeProxy = common.HexToAddress("0xDef1C0ded9bec7F1a1670819833240f027b25EfF")

	EventERC721OrderFilled  = crypto.Keccak256Hash([]byte("ERC721OrderFilled(uint8,address,address,uint256,address,uint256,address,uint256,address)"))
	EventERC1155OrderFilled = crypto.Keccak256Hash([]byte("ERC1155OrderFilled(uint8,address,address,uint256,address,uint256,address,uint256,uint128,address)"))
)
