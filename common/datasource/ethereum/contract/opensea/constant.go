package opensea

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressSeaport = common.HexToAddress("0x00000000006c3852cbEf3e08E8dF289169EdE581")

	EventHashOrderFulfilled = common.BytesToHash(crypto.Keccak256([]byte("OrderFulfilled(bytes32,address,address,address,(uint8,address,uint256,uint256)[],(uint8,address,uint256,uint256,address)[])")))

	PlatformOpenSea = "OpenSea"
)
