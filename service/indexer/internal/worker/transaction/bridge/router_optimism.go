package bridge

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/optimism"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"go.uber.org/zap"
)

func (w *Worker) handleOptimism(ctx context.Context, transaction model.Transaction) (*model.Transaction, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("failed to get ethereum client: %w", err)
	}

	bridge, err := optimism.NewBridge(optimism.AddressGateway, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge contract: %w", err)
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, err
	}

	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	for _, log := range sourceData.Receipt.Logs {
		if len(log.Topics) == 0 {
			return nil, fmt.Errorf("invalid topics of log: %v", log)
		}

		switch log.Topics[0] {
		case optimism.EventHashETHDepositInitiated:
			internalTransfer, err := w.handleOptimismDeposit(ctx, transaction, *log, bridge)
			if err != nil {
				zap.L().Error("failed to handle optimism deposit", zap.Error(err))

				continue
			}

			internalTransaction.Owner = internalTransfer.AddressFrom
			internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(internalTransfer.Tag, internalTransaction.Tag, internalTransfer.Type, internalTransaction.Type)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		case optimism.EventHashERC20DepositInitiated:
			internalTransfer, err := w.handleOptimismDepositERC20(ctx, transaction, *log, bridge)
			if err != nil {
				zap.L().Error("failed to handle optimism deposit", zap.Error(err))

				continue
			}

			internalTransaction.Owner = internalTransfer.AddressFrom
			internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(internalTransfer.Tag, internalTransaction.Tag, internalTransfer.Type, internalTransaction.Type)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		case optimism.EventHashETHWithdrawalFinalized:
			internalTransfer, err := w.handleOptimismWithdraw(ctx, transaction, *log, bridge)
			if err != nil {
				zap.L().Error("failed to handle optimism withdraw", zap.Error(err))

				continue
			}

			internalTransaction.Owner = internalTransfer.AddressFrom
			internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(internalTransfer.Tag, internalTransaction.Tag, internalTransfer.Type, internalTransaction.Type)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		case optimism.EventHashERC20WithdrawalFinalized:
			internalTransfer, err := w.handleOptimismWithdrawERC20(ctx, transaction, *log, bridge)
			if err != nil {
				zap.L().Error("failed to handle optimism withdraw", zap.Error(err))

				continue
			}

			internalTransaction.Owner = internalTransfer.AddressFrom
			internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(internalTransfer.Tag, internalTransaction.Tag, internalTransfer.Type, internalTransaction.Type)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		}
	}

	return &internalTransaction, nil
}

func (w *Worker) handleOptimismDeposit(ctx context.Context, transaction model.Transaction, log types.Log, bridge *optimism.Bridge) (*model.Transfer, error) {
	event, err := bridge.ParseETHDepositInitiated(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ETHDepositInitiated event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.From, event.To, optimism.PlatformBridge, ChainIDOptimism, nil, event.Amount, filter.BridgeDeposit)
}

func (w *Worker) handleOptimismDepositERC20(ctx context.Context, transaction model.Transaction, log types.Log, bridge *optimism.Bridge) (*model.Transfer, error) {
	event, err := bridge.ParseERC20DepositInitiated(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20DepositInitiated event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.From, event.To, optimism.PlatformBridge, ChainIDOptimism, &event.L1Token, event.Amount, filter.BridgeDeposit)
}

func (w *Worker) handleOptimismWithdraw(ctx context.Context, transaction model.Transaction, log types.Log, bridge *optimism.Bridge) (*model.Transfer, error) {
	event, err := bridge.ParseETHWithdrawalFinalized(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ETHWithdrawalFinalized event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.From, event.To, optimism.PlatformBridge, ChainIDOptimism, nil, event.Amount, filter.BridgeWithdraw)
}

func (w *Worker) handleOptimismWithdrawERC20(ctx context.Context, transaction model.Transaction, log types.Log, bridge *optimism.Bridge) (*model.Transfer, error) {
	event, err := bridge.ParseERC20WithdrawalFinalized(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20WithdrawalFinalized event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.From, event.To, optimism.PlatformBridge, ChainIDOptimism, &event.L1Token, event.Amount, filter.BridgeWithdraw)
}
