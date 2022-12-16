package bridge

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/arbitrum"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

func (w *Worker) handleArbitrumHashMessageDelivered(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("unmarshal source data: %w", err)
	}

	var chainID uint64

	switch common.HexToAddress(transaction.AddressTo) {
	case arbitrum.AddressInboxOne:
		chainID = ChainIDArbitrumOne
	case arbitrum.AddressInboxNova:
		chainID = ChainIDArbitrumNova
	default:
		return nil, fmt.Errorf("invalid inbox address: %s", transaction.AddressTo)
	}

	return w.buildTransfer(ctx, transaction, log, common.HexToAddress(transaction.AddressFrom), common.HexToAddress(transaction.AddressFrom), platform, chainID, nil, sourceData.Transaction.Value(), filter.BridgeDeposit)
}
