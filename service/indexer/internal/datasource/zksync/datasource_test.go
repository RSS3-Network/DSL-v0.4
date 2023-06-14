package zksync_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/zksync"
)

var datasource = zksync.New()

func TestName(t *testing.T) {
	transactions, err := datasource.Handle(context.Background(), &protocol.Message{
		Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
		Network: protocol.NetworkZkSync,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, transaction := range transactions {
		t.Log(transaction, len(transaction.Transfers))
	}
}
