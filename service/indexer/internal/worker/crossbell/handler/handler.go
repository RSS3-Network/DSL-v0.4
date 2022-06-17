package handler

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
)

var (
	AddressGenesis            = common.HexToAddress("0x0000000000000000000000000000000000000000")
	AddressProfileProxy       = common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")
	AddressLinkListTokenProxy = common.HexToAddress("0xFc8C75bD5c26F50798758f387B698f207a016b6A")

	EventNameTransfer = "Transfer"

	ErrorUnknownUnknownEvent    = errors.New("unknown event")
	ErrorUnknownContractAddress = errors.New("unknown contract address")
)

type Interface interface {
	Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error)
}

var _ Interface = (*handler)(nil)

type handler struct {
	ethereumClient *ethclient.Client

	profileHandler  Interface
	linkListHandler Interface
}

func (h *handler) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	// Transfer type actions should be handled by the token worker
	if transfer.Source != protocol.SourceOrigin {
		return h.handleTransfer(ctx, transaction, transfer)
	}

	switch common.HexToAddress(transaction.AddressTo) {
	case AddressProfileProxy:
		return h.profileHandler.Handle(ctx, transaction, transfer)
	case AddressLinkListTokenProxy:
		return h.linkListHandler.Handle(ctx, transaction, transfer)
	default:
		return nil, ErrorUnknownContractAddress
	}
}

func (h *handler) handleTransfer(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	transfer.Type = EventNameTransfer

	return &transfer, nil
}

func New(ethereumClient *ethclient.Client, abi abi.ABI) (Interface, error) {
	profileContract, err := contract.NewERC721(AddressProfileProxy, ethereumClient)
	if err != nil {
		return nil, err
	}

	linkListContract, err := contract.NewERC721(AddressLinkListTokenProxy, ethereumClient)
	if err != nil {
		return nil, err
	}

	return &handler{
		ethereumClient: ethereumClient,
		profileHandler: &profile{
			contract: profileContract,
			abi:      abi,
		},
		linkListHandler: &linkList{
			contract: linkListContract,
			abi:      abi,
		},
	}, nil
}
