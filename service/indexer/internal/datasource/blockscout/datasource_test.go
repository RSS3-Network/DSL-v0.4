package blockscout_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
)

func TestName(t *testing.T) {
	blockscoutClient := blockscout.New()

	transactions, err := blockscoutClient.Handle(context.Background(), &protocol.Message{
		Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
		Network: protocol.NetworkXDAI,
	})
	if err != nil {
		t.Log(err)
	}

	for _, transaction := range transactions {
		t.Log(transaction.Hash, len(transaction.Transfers))
	}
}
