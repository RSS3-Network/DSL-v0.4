package lens_test

import (
	"context"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/lens"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

var (
	datasource = lens.New()
)

func TestName(t *testing.T) {
	transfers, err := datasource.Handle(context.Background(), &protocol.Message{
		Address: "0x3a5bd1e37b099ae3386d13947b6a90d97675e5e3",
		Network: protocol.NetworkPolygon,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, transfer := range transfers {
		t.Log(transfer)
	}
}
