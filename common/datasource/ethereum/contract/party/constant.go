package party

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	EventHashPartyBidDeployed = common.BytesToHash(crypto.Keccak256([]byte("PartyBidDeployed(address,address,address,uint256,address,uint256,address,uint256,address,uint256,string,string)")))

	EventHashPartyBuyDeployed = common.BytesToHash(crypto.Keccak256([]byte("PartyBuyDeployed(address,address,address,uint256,uint256,uint256,address,uint256,address,uint256,string,string)")))

	EventHashCollectionPartyDeployed = common.BytesToHash(crypto.Keccak256([]byte("CollectionPartyDeployed(address,address,address,uint256,uint256,address[],address,uint256,address,uint256,string,string)")))

	EventHashContributed = common.BytesToHash(crypto.Keccak256([]byte("Contributed(address,uint256,uint256,uint256)")))
	EventHashClaimed     = common.BytesToHash(crypto.Keccak256([]byte("Claimed(address,uint256,uint256,uint256)")))

	EventHashBid       = common.BytesToHash(crypto.Keccak256([]byte("Bid(uint256)")))
	EventHashFinalized = common.BytesToHash(crypto.Keccak256([]byte("Finalized(uint8,uint256,uint256,uint256)")))

	EventHashPbBought  = common.BytesToHash(crypto.Keccak256([]byte("Bought(address,address,uint256,uint256,uint256)")))
	EventHashPbExpired = common.BytesToHash(crypto.Keccak256([]byte("Expired(address)")))

	EventHashPcBought  = common.BytesToHash(crypto.Keccak256([]byte("Bought(uint256,address,address,uint256,uint256,uint256)")))
	EventHashPcExpired = common.BytesToHash(crypto.Keccak256([]byte("Expired(address)")))

	AddressPartyBidDeployed        = common.HexToAddress("0x744c2BE04d079eDdb21c1a9BB13bB5259A368614")
	AddressPartyBuyDeployed        = common.HexToAddress("0x10Fb0f8860c11cCcdf85f2f56e3B8e1DdEb2BA3F")
	AddressCollectionPartyDeployed = common.HexToAddress("0xd084d7849d4EBE564A2a41e085b2a74f6DDe5300")
	AddressPartyBid                = common.HexToAddress("0x744c2BE04d079eDdb21c1a9BB13bB5259A368614")
	AddressPartyBuy                = common.HexToAddress("0x2045427276B2Ad409202eea1E0C81E150f3203E4")
	AddressPartyCollection         = common.HexToAddress("0x0C696f63A8Cfd4b456f725f1174f1D5B48D1e876")

	AddressMapToMarket = map[string]string{
		"0x96e5b0519983f2f984324b926e6d28C3A4Eb92A1": "foundation",
		"0x11c07cE1315a3b92C9755F90cDF40B04b88c5731": "zora",
		"0x9319DAd8736D752C5c72DB229f8e1b280DC80ab1": "nouns",
		"0xe761685c3E5Af758fc905CfdaC8Ca9f767CF84D7": "koans",
		"opensea": "opensea",
	}

	PlatformPartyBid = "PartyBid"
)
