package opensea

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	ItemTypeERC20 uint8 = 1

	Platform = "OpenSea"
)

var (
	AddressSeaport          = common.HexToAddress("0x00000000006c3852cbEf3e08E8dF289169EdE581")
	AddressWyvernExchangeV1 = common.HexToAddress("0x7Be8076f4EA4A4AD08075C2508e481d6C946D12b")
	AddressWyvernExchangeV2 = common.HexToAddress("0x7f268357A8c2552623316e2562D90e642bB538E5")

	EventHashOrderFulfilled = crypto.Keccak256Hash([]byte("OrderFulfilled(bytes32,address,address,address,(uint8,address,uint256,uint256)[],(uint8,address,uint256,uint256,address)[])"))
	EventHashOrdersMatched  = crypto.Keccak256Hash([]byte("OrdersMatched(bytes32,bytes32,address,address,uint256,bytes32)"))
)
