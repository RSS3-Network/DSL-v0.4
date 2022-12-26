package ethereum

import "github.com/ethereum/go-ethereum/common"

var (
	AddressGenesis = common.HexToAddress("0x0000000000000000000000000000000000000000")
	HashGenesis    = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
)

const RPCMaxConcurrency = 200

// https://etherscan.io/accounts/label/burn
func IsBlackHoleAddress(address common.Address) bool {
	switch address {
	case AddressGenesis,
		common.HexToAddress("0x0000000000000000000000000000000000000001"),
		common.HexToAddress("0x0000000000000000000000000000000000000002"),
		common.HexToAddress("0x0000000000000000000000000000000000000003"),
		common.HexToAddress("0x0000000000000000000000000000000000000004"),
		common.HexToAddress("0x0000000000000000000000000000000000000005"),
		common.HexToAddress("0x0000000000000000000000000000000000000006"),
		common.HexToAddress("0x0000000000000000000000000000000000000007"),
		common.HexToAddress("0x0000000000000000000000000000000000000008"),
		common.HexToAddress("0x0000000000000000000000000000000000000009"),
		common.HexToAddress("0x0123456789012345678901234567890123456789"),
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x1234567890123456789012345678901234567890"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
		common.HexToAddress("0x6666666666666666666666666666666666666666"),
		common.HexToAddress("0x8888888888888888888888888888888888888888"),
		common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"),
		common.HexToAddress("0xdead000000000000000042069420694206942069"),
		common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"),
		common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffff"),
		common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		common.HexToAddress("0x000000000000000000000000000000000000dEaD"):

		return true
	default:
		return false
	}
}
