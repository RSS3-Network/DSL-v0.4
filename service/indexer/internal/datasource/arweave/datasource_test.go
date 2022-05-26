package arweave_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/mirror"
)

var (
	datasource = arweave.New()
)

func TestName(t *testing.T) {
	transfers, err := datasource.Handle(context.Background(), &protocol.Message{
		Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
		Network: protocol.NetworkArweave,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, transfer := range transfers {
		t.Log(transfer)
	}

	client := mirror.New()

	transfers, err = client.Handle(context.Background(), &protocol.Message{
		Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
		Network: protocol.NetworkArweave,
	}, transfers)
	if err != nil {
		t.Fatal(err)
	}

	for _, transfer := range transfers {
		metadataModel := metadata.Metadata{}

		if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
			t.Fatal(err)
		}

		t.Log(metadataModel.Mirror)
	}
}
