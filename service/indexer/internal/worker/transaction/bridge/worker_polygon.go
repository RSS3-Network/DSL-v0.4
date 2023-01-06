package bridge

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/polygon"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

func (w *Worker) handlePolygonLockedEther(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := polygon.NewEtherPredicateFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("get bridge contract: %w", err)
	}

	event, err := filterer.ParseLockedEther(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LockedEther event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.Depositor, event.DepositReceiver, polygon.PlatformBridge, ChainIDPolygon, nil, event.Amount, filter.BridgeDeposit)
}

func (w *Worker) handlePolygonExitedEther(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := polygon.NewEtherPredicateFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("get bridge contract: %w", err)
	}

	event, err := filterer.ParseExitedEther(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ExitedEther event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, common.HexToAddress(transaction.AddressFrom), event.Exitor, polygon.PlatformBridge, ChainIDPolygon, nil, event.Amount, filter.BridgeWithdraw)
}

func (w *Worker) handlePolygonLockedERC20(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := polygon.NewERC20PredicateFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("get bridge contract: %w", err)
	}

	event, err := filterer.ParseLockedERC20(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LockedMintableERC20 event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.Depositor, event.DepositReceiver, polygon.PlatformBridge, ChainIDPolygon, &event.RootToken, event.Amount, filter.BridgeDeposit)
}

func (w *Worker) handlePolygonLockedMintableERC20(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := polygon.NewMintableERC20PredicateFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("get bridge contract: %w", err)
	}

	event, err := filterer.ParseLockedMintableERC20(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LockedMintableERC20 event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.Depositor, event.DepositReceiver, polygon.PlatformBridge, ChainIDPolygon, &event.RootToken, event.Amount, filter.BridgeDeposit)
}
