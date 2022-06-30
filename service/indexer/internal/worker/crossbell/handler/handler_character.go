package handler

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract/character"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract/periphery"
)

var _ Interface = (*characterHandler)(nil)

type characterHandler struct {
	characterContract *character.Character
	peripheryContract *periphery.Periphery
}

func (c *characterHandler) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	return nil, contract.ErrorUnknownEvent
}
