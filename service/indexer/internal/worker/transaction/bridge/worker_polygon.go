package bridge

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/polygon"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

func (w *Worker) handlePolygon(ctx context.Context, transaction model.Transaction) (*model.Transaction, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("failed to get ethereum client: %w", err)
	}

	etherPredicate, err := polygon.NewEtherPredicate(polygon.AddressEtherePredicate, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create ether predicate: %w", err)
	}

	erc20Predicate, err := polygon.NewERC20Predicate(polygon.AddressERC20Predicate, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get erc20 predicate: %w", err)
	}

	mintableERC20Predicate, err := polygon.NewMintableERC20Predicate(polygon.AddressMintableERC20Predicate, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get mintable erc20 predicate: %w", err)
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
		case polygon.EventHashLockedEther:
			internalTransfer, err := w.handlePolygonDepositEther(ctx, transaction, *log, etherPredicate)
			if err != nil {
				return nil, fmt.Errorf("failed to handle polygon deposit: %w", err)
			}

			internalTransaction = w.fillTransactionMetadata(internalTransaction, *internalTransfer)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		case polygon.EventHashExitedEther:
			internalTransfer, err := w.handlePolygonWithdrawEther(ctx, transaction, *log, etherPredicate)
			if err != nil {
				return nil, fmt.Errorf("failed to handle polygon withdraw: %w", err)
			}

			internalTransaction = w.fillTransactionMetadata(internalTransaction, *internalTransfer)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		case polygon.EventHashLockedERC20:
			internalTransfer, err := w.handlePolygonDepositERC20(ctx, transaction, *log, erc20Predicate)
			if err != nil {
				return nil, fmt.Errorf("failed to handle polygon withdraw: %w", err)
			}

			internalTransaction = w.fillTransactionMetadata(internalTransaction, *internalTransfer)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		case polygon.EventHashLockedMintableERC20:
			internalTransfer, err := w.handlePolygonDepositMintableERC20(ctx, transaction, *log, mintableERC20Predicate)
			if err != nil {
				return nil, fmt.Errorf("failed to handle polygon withdraw: %w", err)
			}

			internalTransaction = w.fillTransactionMetadata(internalTransaction, *internalTransfer)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		}
	}

	return &internalTransaction, nil
}

func (w *Worker) handlePolygonDepositEther(ctx context.Context, transaction model.Transaction, log types.Log, predicate *polygon.EtherPredicate) (*model.Transfer, error) {
	event, err := predicate.ParseLockedEther(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LockedEther event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.Depositor, event.DepositReceiver, polygon.PlatformBridge, ChainIDPolygon, nil, event.Amount, filter.BridgeDeposit)
}

func (w *Worker) handlePolygonWithdrawEther(ctx context.Context, transaction model.Transaction, log types.Log, predicate *polygon.EtherPredicate) (*model.Transfer, error) {
	event, err := predicate.ParseExitedEther(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ExitedEther event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, common.HexToAddress(transaction.AddressFrom), event.Exitor, polygon.PlatformBridge, ChainIDPolygon, nil, event.Amount, filter.BridgeWithdraw)
}

func (w *Worker) handlePolygonDepositERC20(ctx context.Context, transaction model.Transaction, log types.Log, predicate *polygon.ERC20Predicate) (*model.Transfer, error) {
	event, err := predicate.ParseLockedERC20(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LockedMintableERC20 event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.Depositor, event.DepositReceiver, polygon.PlatformBridge, ChainIDPolygon, &event.RootToken, event.Amount, filter.BridgeDeposit)
}

func (w *Worker) handlePolygonDepositMintableERC20(ctx context.Context, transaction model.Transaction, log types.Log, predicate *polygon.MintableERC20Predicate) (*model.Transfer, error) {
	event, err := predicate.ParseLockedMintableERC20(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LockedMintableERC20 event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.Depositor, event.DepositReceiver, polygon.PlatformBridge, ChainIDPolygon, &event.RootToken, event.Amount, filter.BridgeDeposit)
}
