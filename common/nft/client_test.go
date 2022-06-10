package nft_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/nft"
	"github.com/stretchr/testify/assert"
)

func setup() {
	nft.Initialize("56ff3a0f0ad14ec19dc933038c69bfbe")
}

func teardown() {
	nft.Initialize("")
}

func TestGetMetadata(t *testing.T) {
	metadata, err := nft.GetMetadata(
		nft.NetworkEthereum,
		common.HexToAddress("0xacbe98efe2d4d103e221e04c76d7c55db15c8e89"),
		big.NewInt(1),
	)

	assert.Nil(t, err)
	assert.Contains(t, string(metadata), "right fruit")
}
