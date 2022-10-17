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

	predicate, err := polygon.NewPredicate(polygon.AddressPredicate, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get predicate contract: %w", err)
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
			internalTransfer, err := w.handlePolygonDeposit(ctx, transaction, *log, predicate)
			if err != nil {
				return nil, fmt.Errorf("failed to handle polygon deposit: %w", err)
			}

			internalTransaction.Owner = internalTransfer.AddressFrom
			internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(internalTransfer.Tag, internalTransaction.Tag, internalTransfer.Type, internalTransaction.Type)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		case polygon.EventHashExitedEther:
			internalTransfer, err := w.handlePolygonWithdraw(ctx, transaction, *log, predicate)
			if err != nil {
				return nil, fmt.Errorf("failed to handle polygon withdraw: %w", err)
			}

			internalTransaction.Owner = internalTransfer.AddressFrom
			internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(internalTransfer.Tag, internalTransaction.Tag, internalTransfer.Type, internalTransaction.Type)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		}
	}

	return &internalTransaction, nil
}

func (w *Worker) handlePolygonDeposit(ctx context.Context, transaction model.Transaction, log types.Log, predicate *polygon.Predicate) (*model.Transfer, error) {
	event, err := predicate.ParseLockedEther(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LockedEther event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, event.Depositor, event.DepositReceiver, polygon.PlatformBridge, ChainIDPolygon, nil, event.Amount, filter.BridgeDeposit)
}

func (w *Worker) handlePolygonWithdraw(ctx context.Context, transaction model.Transaction, log types.Log, predicate *polygon.Predicate) (*model.Transfer, error) {
	event, err := predicate.ParseExitedEther(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ExitedEther event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, common.HexToAddress(transaction.AddressFrom), event.Exitor, polygon.PlatformBridge, ChainIDPolygon, nil, event.Amount, filter.BridgeWithdraw)
}
