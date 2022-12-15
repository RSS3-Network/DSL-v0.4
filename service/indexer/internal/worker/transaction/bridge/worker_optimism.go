package bridge

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/optimism"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

func (w *Worker) handleOptimismETHDepositInitiated(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	bridge, err := optimism.NewBridge(optimism.AddressGateway, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("get bridge contract: %w", err)
	}

	event, err := bridge.ParseETHDepositInitiated(log)
	if err != nil {
		return nil, fmt.Errorf("parse ETHDepositInitiated event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.From, event.To, optimism.PlatformBridge, ChainIDOptimism, nil, event.Amount, filter.BridgeDeposit)
}

func (w *Worker) handleOptimismERC20DepositInitiated(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	bridge, err := optimism.NewBridge(optimism.AddressGateway, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("get bridge contract: %w", err)
	}

	event, err := bridge.ParseERC20DepositInitiated(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20DepositInitiated event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.From, event.To, optimism.PlatformBridge, ChainIDOptimism, &event.L1Token, event.Amount, filter.BridgeDeposit)
}

func (w *Worker) handleOptimismETHWithdrawalFinalized(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	bridge, err := optimism.NewBridge(optimism.AddressGateway, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("get bridge contract: %w", err)
	}

	event, err := bridge.ParseETHWithdrawalFinalized(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ETHWithdrawalFinalized event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.From, event.To, optimism.PlatformBridge, ChainIDOptimism, nil, event.Amount, filter.BridgeWithdraw)
}

func (w *Worker) handleOptimismERC20WithdrawalFinalized(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	bridge, err := optimism.NewBridge(optimism.AddressGateway, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("get bridge contract: %w", err)
	}

	event, err := bridge.ParseERC20WithdrawalFinalized(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20WithdrawalFinalized event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.From, event.To, optimism.PlatformBridge, ChainIDOptimism, &event.L1Token, event.Amount, filter.BridgeWithdraw)
}
