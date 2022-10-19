package opensea

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressSeaport        = common.HexToAddress("0x00000000006c3852cbEf3e08E8dF289169EdE581")
	AddressWyvernExchange = common.HexToAddress("0x7Be8076f4EA4A4AD08075C2508e481d6C946D12b")

	EventHashOrderFulfilled = crypto.Keccak256Hash([]byte("OrderFulfilled(bytes32,address,address,address,(uint8,address,uint256,uint256)[],(uint8,address,uint256,uint256,address)[])"))
	EventHashOrdersMatched  = crypto.Keccak256Hash([]byte("OrdersMatched(bytes32,bytes32,address,address,uint256,bytes32)"))

	Platform = "OpenSea"
)
