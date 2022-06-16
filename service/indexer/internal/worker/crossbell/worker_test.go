package crossbell_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell"
)

func TestName(t *testing.T) {
	message := protocol.Message{
		Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
		Network: protocol.NetworkCrossbell,
	}

	datasource := blockscout.New()

	transactions, err := datasource.Handle(context.Background(), &message)
	if err != nil {
		t.Fatal(err)
	}

	worker := crossbell.New()

	if err := worker.Initialize(context.Background()); err != nil {
		t.Fatal(err)
	}

	transactions, err = worker.Handle(context.Background(), &message, transactions)
	if err != nil {
		t.Fatal(err)
	}

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			t.Log(string(transfer.Metadata))
		}
	}
}
