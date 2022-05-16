package worker

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
)

type Worker interface {
	Handle(ctx context.Context, transfers []model.Transfer) ([]model.Transfer, error)
}
