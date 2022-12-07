package tofunft

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressMarketplace = common.HexToAddress("0x7bc8b1B5AbA4dF3Be9f9A32daE501214dC0E4f3f")

	// EventEvInventoryUpdate hash is 0x5beea7b3b87c573953fec05007114d17712e5775d364acc106d8da9e74849033
	EventEvInventoryUpdate = crypto.Keccak256Hash([]byte("EvInventoryUpdate(uint256,(address,address,address,uint256,uint256,uint256,uint8,uint8))"))
)
