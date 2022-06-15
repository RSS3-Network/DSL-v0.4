package worker

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Worker interface {
	Name() string
	Networks() []string
	Initialize(ctx context.Context) error
	Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error)
	Jobs() []Job
}
