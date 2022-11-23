package lido

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	Address = common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84")

	EventHashSubmitted = crypto.Keccak256Hash([]byte("Submitted(address,uint256,address)"))
)
