package snapshot

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

const (
	Source = "snapshot"
)

type Datasource struct {
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {

}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	// get snapshot list

	// set snapshot in Transfers
}
