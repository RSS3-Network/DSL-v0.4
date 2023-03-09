package zora

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressAsks = common.HexToAddress("0x6170B3C3A54C3d8c854934cBC314eD479b2B29A3")

	EventAskFilled = crypto.Keccak256Hash([]byte("AskFilled(address,uint256,address,address,(address,address,address,uint16,uint256))"))
)
