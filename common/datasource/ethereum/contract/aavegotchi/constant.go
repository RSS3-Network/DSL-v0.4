package aavegotchi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate abigen --abi ./contract/ERC1155MarketplaceFacet.abi --pkg contract --type ERC1155Marketplace --out ./contract/ERC1155MarketplaceFacet.go
//go:generate abigen --abi ./contract/ERC721MarketplaceFacet.abi --pkg contract --type ERC721Marketplace --out ./contract/ERC721MarketplaceFacet.go
//go:generate abigen --abi ./contract/AavegotchiGameFacets.abi --pkg contract --type Game --out ./contract/AavegotchiGameFacets.go

var (
	AavegotchiContractAddress = common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d")

	EventERC1155ExecutedListing  = crypto.Keccak256Hash([]byte("ERC1155ExecutedListing(uint256,address,address,address,uint256,uint256,uint256,uint256,uint256)"))
	EventERC1155ListingCancelled = crypto.Keccak256Hash([]byte("ERC1155ListingCancelled(uint256)"))
	EventERC1155ListingAdd       = crypto.Keccak256Hash([]byte("ERC1155ListingAdd(uint256,address,address,uint256,uint256,uint256,uint256,uint256)"))

	EventERC721ListingAdd      = crypto.Keccak256Hash([]byte("ERC721ListingAdd(uint256,address,address,uint256,uint256,uint256)"))
	EventERC721ExecutedListing = crypto.Keccak256Hash([]byte("ERC721ExecutedListing(uint256,address,address,address,uint256,uint256,uint256,uint256)"))

	EventClaimAavegotchi = crypto.Keccak256Hash([]byte("ClaimAavegotchi(uint256)"))
)
