package handler

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
	"go.opentelemetry.io/otel"
)

type Interface interface {
	Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error)
}

var _ Interface = (*handler)(nil)

type handler struct {
	ethereumClient   *ethclient.Client
	characterHandler Interface
	profileHandler   Interface
	linkListHandler  Interface
}

func (h *handler) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handle")

	_, span := tracer.Start(ctx, "worker_crossbell_handle:Handle")

	defer span.End()

	// Transfer type actions should be handled by the token worker
	if transfer.Source != protocol.SourceOrigin {
		return &transfer, nil
	}

	switch common.HexToAddress(transaction.AddressTo) {
	case contract.AddressCharacter:
		// Broken change
		if transaction.BlockNumber >= contract.BrokenBlockNumber {
			return h.characterHandler.Handle(ctx, transaction, transfer)
		}

		return h.profileHandler.Handle(ctx, transaction, transfer)
	case contract.AddressLinkList:
		return h.linkListHandler.Handle(ctx, transaction, transfer)
	default:
		return nil, contract.ErrorUnknownContractAddress
	}
}

func New(ethereumClient *ethclient.Client) Interface {
	return &handler{
		ethereumClient:   ethereumClient,
		characterHandler: &character{},
		profileHandler:   &profile{},
		linkListHandler:  &linkList{},
	}
}
