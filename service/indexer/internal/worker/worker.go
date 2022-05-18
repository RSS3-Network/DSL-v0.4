package worker

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Worker interface {
	Name() string
	Network() []string
	Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error)
}
