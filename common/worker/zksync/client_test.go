package zksync_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/common/worker/zksync"
)

var client = zksync.New()

func TestGetTokenList(t *testing.T) {
	tokenList, err := client.GetTokenList()
	if err != nil {
		t.Fatal(err)
	}

	for _, token := range tokenList {
		t.Log(token.ID, token.Address, token.Symbol, token.Kind, token.IsNFT)
	}
}
