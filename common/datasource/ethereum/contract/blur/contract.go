package blur

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressMarketplace  = common.HexToAddress("0x39da41747a83aeE658334415666f3EF92DD0D541")
	AddressMarketplace2 = common.HexToAddress("0x000000000000Ad05Ccc4F10045630fb830B95127")
	AddressPool         = common.HexToAddress("0x0000000000A39bb272e79075ade125fd351887Ac")

	EventOrdersMatched = crypto.Keccak256Hash([]byte("OrdersMatched(address,address,(address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),bytes32,(address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),bytes32)"))
)
