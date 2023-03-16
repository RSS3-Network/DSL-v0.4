package foundation

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressFoundationMarket = common.HexToAddress("0xcDA72070E455bb31C7690a170224Ce43623d0B6f")

	EventOfferAccepted    = crypto.Keccak256Hash([]byte("OfferAccepted(address,uint256,address,address,uint256,uint256,uint256)"))
	EventBuyPriceAccepted = crypto.Keccak256Hash([]byte("BuyPriceAccepted(address,uint256,address,address,uint256,uint256,uint256)"))
)
