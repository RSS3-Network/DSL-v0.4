package datasource_asset

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Datasource interface {
	Name() string
	Networks() []string
	Handle(ctx context.Context, message *protocol.Message) ([]model.Asset, error)
}
