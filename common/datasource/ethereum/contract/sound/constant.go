package sound

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate abigen --abi ./contract/artist.abi --pkg contract --type Artist --out ./contract/artist.go
//go:generate abigen --abi ./contract/artist_v3.abi --pkg contract --type ArtistV3 --out ./contract/artist_v3.go
//go:generate abigen --abi ./contract/artist_v5.abi --pkg contract --type ArtistV5 --out ./contract/artist_v5.go

var (
	// AddressArtistCreator https://etherscan.io/address/0x78E3aDc0E811e4f93BD9F1f9389b923c9A3355c2#code
	AddressArtistCreator = common.HexToAddress("0x78e3adc0e811e4f93bd9f1f9389b923c9a3355c2")

	// EventHashArtistCreated 0x23748b43b77f98380e738976c6324996908ffc1989994dd3c68631c87a65a7c0
	EventHashArtistCreated = crypto.Keccak256Hash([]byte("CreatedArtist(uint256,string,string,address)"))

	// EventHashEditionCreatedV1 0xb3131d7d301f8caeb40981cffc627b1fdf324b5e4a23845b61c1a6ad2a25f385
	EventHashEditionCreatedV1 = crypto.Keccak256Hash([]byte("EditionCreated(uint256,address,uint256,uint32,uint32,uint32,uint32)"))

	// EventHashEditionCreatedV3 0xb56f9ba6a8af17a142f8ad372c6fc283e81d8c6193b86afe7f6ca901acd698f3
	EventHashEditionCreatedV3 = crypto.Keccak256Hash([]byte("EditionCreated(uint256,address,uint256,uint32,uint32,uint32,uint32,uint32,address)"))

	// EventHashEditionPurchasedV3 0xe38cb07a52e5d88a83de7c9d29c2841118103e462d20f8c526b35872f9977785
	EventHashEditionPurchasedV3 = crypto.Keccak256Hash([]byte("EditionPurchased(uint256,uint256,uint32,address)"))

	// EventHashEditionPurchasedV5 0xc3e4ad784e3889561b308df24921cf08a47410a8ed63b8afe80c50a002efb251
	EventHashEditionPurchasedV5 = crypto.Keccak256Hash([]byte("EditionPurchased(uint256,uint256,uint32,address,uint256)"))
)
