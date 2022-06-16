package handler

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
)

var _ Interface = (*linkList)(nil)

type linkList struct {
	contract *contract.ERC721
	abi      abi.ABI
}

func (l *linkList) Handle(ctx context.Context, transaction *model.Transaction, log *types.Log) (*model.Transfer, error) {
	return nil, nil
}
