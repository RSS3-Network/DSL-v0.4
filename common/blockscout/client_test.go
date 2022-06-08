package blockscout_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/blockscout"
)

var client = blockscout.New(blockscout.NetworkXDAI)

func TestGetTransactionList(t *testing.T) {
	transactions, _, err := client.GetTransactionList(
		context.Background(),
		common.HexToAddress("0x000000A52a03835517E9d193B3c27626e1Bc96b1"),
		&blockscout.GetTransactionListOption{},
	)
	if err != nil {
		t.Fatal(err)
	}

	for _, transaction := range transactions {
		t.Log(transaction)
	}
}

func TestGetTokenTransactionList(t *testing.T) {
	transactions, _, err := client.GetTokenTransactionList(
		context.Background(),
		common.HexToAddress("0x000000A52a03835517E9d193B3c27626e1Bc96b1"),
		&blockscout.GetTokenTransactionListOption{},
	)
	if err != nil {
		t.Fatal(err)
	}

	for _, transaction := range transactions {
		t.Log(transaction)
	}
}
