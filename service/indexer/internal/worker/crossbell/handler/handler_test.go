package handler_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multibase"
)

func TestName(t *testing.T) {
	linkKey := common.HexToHash("0xaba33d387a5fd695aad15a440e089f5285abd0d7d89af71646e6872ec46946a4")

	// bafkreiaz47imwc6xhohd6e25gjrscwokpeowyttbblejsvq7ahc46qj6ga
	t.Log(cid.NewCidV1(multibase.Base16, linkKey.Bytes()))
}
