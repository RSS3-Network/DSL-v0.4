package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/arbitrum"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

func (w *Worker) handleArbitrum(ctx context.Context, transaction model.Transaction) (*model.Transaction, error) {
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

		if log.Topics[0] == arbitrum.EventHashMessageDelivered {
			internalTransfer, err := w.handleArbitrumDeposit(ctx, transaction, *log, sourceData.Transaction.Value())
			if err != nil {
				return nil, fmt.Errorf("failed to handle arbitrum deposit: %w", err)
			}

			internalTransaction = w.fillTransactionMetadata(internalTransaction, *internalTransfer)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		}
	}

	return &internalTransaction, nil
}

func (w *Worker) handleArbitrumDeposit(ctx context.Context, transaction model.Transaction, log types.Log, value *big.Int) (*model.Transfer, error) {
	var (
		platform string
		chainID  uint64
	)

	switch common.HexToAddress(transaction.AddressTo) {
	case arbitrum.AddressInboxOne:
		platform = arbitrum.PlatformInboxOne
		chainID = ChainIDArbitrumOne
	case arbitrum.AddressInboxNova:
		platform = arbitrum.PlatformInboxNova
		chainID = ChainIDArbitrumNova
	default:
		return nil, fmt.Errorf("invalid inbox address: %s", transaction.AddressTo)
	}

	return w.buildTransfer(ctx, transaction, log, common.HexToAddress(transaction.AddressFrom), common.HexToAddress(transaction.AddressFrom), platform, chainID, nil, value, filter.BridgeDeposit)
}
