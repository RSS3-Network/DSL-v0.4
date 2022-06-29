package handler

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
)

var _ Interface = (*character)(nil)

type character struct {
	contract any
}

func (c *character) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	return nil, contract.ErrorUnknownUnknownEvent
}
