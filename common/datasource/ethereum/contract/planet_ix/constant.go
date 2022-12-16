package planet_ix

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate abigen --abi ./contract/marketplace.abi --pkg contract --type Marketplace --out ./contract/marketplace.go

var (
	PlanetIXMarketplaceContractAddress = common.HexToAddress("0x5a98e7ce0f72995fdc13d91255443f374c1299a6")

	EventOwnershipTransferred   = crypto.Keccak256Hash([]byte("OwnershipTransferred(address,address)"))
	EventPurchased              = crypto.Keccak256Hash([]byte("Purchased(address,address,uint256,uint256)"))                      // buy
	EventPurchasedWithSignature = crypto.Keccak256Hash([]byte("PurchasedWithSignature(address,address,address,uint256,uint256)")) // sell
	EventSaleCancelled          = crypto.Keccak256Hash([]byte("SaleCancelled(uint256)"))
	EventSaleRequested          = crypto.Keccak256Hash([]byte("SaleRequested(address,uint256,address,uint256[],uint256)"))
)
