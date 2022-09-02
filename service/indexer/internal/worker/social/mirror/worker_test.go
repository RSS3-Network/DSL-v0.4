package mirror_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/mirror"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
)

func TestName(t *testing.T) {
	datasource := arweave.New()

	transactions, err := datasource.Handle(context.Background(), &protocol.Message{
		Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
		Network: protocol.NetworkArweave,
	})
	if err != nil {
		t.Log(err)
	}

	worker := mirror.New()
	transactions, err = internalModel.Handle(context.Background(), &protocol.Message{}, transactions)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			t.Log(string(transfer.Metadata))
		}
	}
}
