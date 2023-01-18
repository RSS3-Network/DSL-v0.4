package datasource

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

const DatasourceLimit = 40000

type Datasource interface {
	Name() string
	Networks() []string
	Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error)
}
