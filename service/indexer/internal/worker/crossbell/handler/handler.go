package handler

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract/character"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract/linklist"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract/periphery"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract/profile"
	"go.opentelemetry.io/otel"
)

type Interface interface {
	Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error)
}

var _ Interface = (*handler)(nil)

type handler struct {
	characterHandler Interface
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

	var log types.Log

	if err := json.Unmarshal(transfer.SourceData, &log); err != nil {
		return nil, err
	}

	switch common.HexToAddress(transaction.AddressTo) {
	case contract.AddressCharacter:
		return h.characterHandler.Handle(ctx, transaction, transfer)
	case contract.AddressLinkList:
		return h.linkListHandler.Handle(ctx, transaction, transfer)
	default:
		return nil, contract.ErrorUnknownContractAddress
	}
}

func New(ethereumClient *ethclient.Client) (Interface, error) {
	profileContract, err := profile.NewProfile(contract.AddressCharacter, ethereumClient)
	if err != nil {
		return nil, err
	}

	characterContract, err := character.NewCharacter(contract.AddressCharacter, ethereumClient)
	if err != nil {
		return nil, err
	}

	peripheryContract, err := periphery.NewPeriphery(contract.AddressPeriphery, ethereumClient)
	if err != nil {
		return nil, err
	}

	linkListContract, err := linklist.NewLinkList(contract.AddressLinkList, ethereumClient)
	if err != nil {
		return nil, err
	}

	return &handler{
		characterHandler: &characterHandler{
			characterContract: characterContract,
			peripheryContract: peripheryContract,
			profileHandler: &profileHandler{
				profileContract: profileContract,
			},
		},
		linkListHandler: &linkListHandler{
			linkListContract: linkListContract,
		},
	}, nil
}
