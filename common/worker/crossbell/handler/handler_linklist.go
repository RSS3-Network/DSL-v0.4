package handler

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell/contract/linklist"
)

var _ Interface = (*linkListHandler)(nil)

type linkListHandler struct {
	linkListContract *linklist.LinkList
}

func (l *linkListHandler) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	return nil, crossbell.ErrorUnknownEvent
}
