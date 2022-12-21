package masknetwork

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressHappyTokenPool = common.HexToAddress("0xF9F7C1496c21bC0180f4B64daBE0754ebFc8A8c0")

	EventSwapSuccess = crypto.Keccak256Hash([]byte("SwapSuccess(bytes32,address,address,address,uint256,uint256,uint128,bool)"))
)
