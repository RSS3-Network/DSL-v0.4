package looksrare

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressExchange = common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a")

	EventHashRoyaltyPayment = common.BytesToHash(crypto.Keccak256([]byte("RoyaltyPayment(address,uint256,address,address,uint256)")))
	EventHashTakerBid       = common.BytesToHash(crypto.Keccak256([]byte("TakerBid(bytes32,uint256,address,address,address,address,address,uint256,uint256,uint256)")))
	EventHashTakerAsk       = common.BytesToHash(crypto.Keccak256([]byte("TakerAsk(bytes32,uint256,address,address,address,address,address,uint256,uint256,uint256)")))

	PlatformLooksrare = "Looksrare"
)
