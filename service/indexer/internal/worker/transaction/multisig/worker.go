package multisig

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/safe"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"

	"go.uber.org/zap"
)

var _ worker.Worker = (*MultiSign)(nil)

type MultiSign struct {
	gnosisSafeABI      abi.ABI
	gnosisSafeFilterer *safe.GnosisSafeFilterer
}

func (m *MultiSign) Name() string {
	return protocol.PlatformGnosisSafe
}

func (m *MultiSign) Networks() []string {
	return protocol.EthclientNetworks
}

func (m *MultiSign) Initialize(ctx context.Context) error {
	gnosisSafeABI, err := contract.ContractABIs.Open("safe/GnosisSafe.abi")
	if err != nil {
		return fmt.Errorf("open gnosis safe abi: %w", err)
	}

	m.gnosisSafeABI, err = abi.JSON(gnosisSafeABI)
	if err != nil {
		return fmt.Errorf("parse gnosis safe abi: %w", err)
	}

	return nil
}

func (m *MultiSign) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	var (
		waitGroup sync.WaitGroup
		locker    sync.Mutex
	)

	result := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		// Ignore transactions that have been processed upstream
		if transaction.Tag != "" && transaction.Type != "" {
			continue
		}

		transaction := transaction
		transaction.Transfers = make([]model.Transfer, 0)

		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			var sourceData ethereum.SourceData
			if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
				zap.L().Error("unmarshal source data", zap.Error(err), zap.String("transaction_hash", transaction.Hash))

				return
			}

			// Upstream may not have built the correct source data
			if sourceData.Transaction == nil || sourceData.Receipt == nil {
				zap.L().Error("invalid source data", zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

				return
			}

			for _, log := range sourceData.Receipt.Logs {
				transfer, err := m.handle(ctx, transaction, *log)
				if err != nil {
					if !errors.Is(err, ethereum.ErrorUnsupportedEvent) {
						zap.L().Warn("handle gnosis safe event", zap.Error(err), zap.String("transaction_hash", transaction.Hash))
					}

					continue
				}

				transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
				transaction.Platform = transfer.Platform
				transaction.Transfers = append(transaction.Transfers, *transfer)
			}

			locker.Lock()
			defer locker.Unlock()

			if transaction.Tag != "" && transaction.Type != "" {
				result = append(result, transaction)
			}
		}()
	}

	waitGroup.Wait()

	return result, nil
}

func (m *MultiSign) handle(ctx context.Context, transaction model.Transaction, log types.Log) (transfer *model.Transfer, err error) {
	switch log.Topics[0] {
	case safe.EventHashAddedOwner:
		transfer, err = m.handleGnosisSafeEventAddedOwner(ctx, transaction, log)
	case safe.EventHashRemovedOwner:
		transfer, err = m.handleGnosisSafeEventRemovedOwner(ctx, transaction, log)
	case safe.EventHashChangedThreshold:
		transfer, err = m.handleGnosisSafeEventChangedThreshold(ctx, transaction, log)
	case safe.EventHashExecutionSuccess:
		transfer, err = m.handleGnosisSafeEventExecutionSuccess(ctx, transaction, log)
	case safe.EventHashExecutionFailure:
		transfer, err = m.handleGnosisSafeEventExecutionFailure(ctx, transaction, log)
	case safe.EventHashSafeSetup:
		transfer, err = m.handleGnosisSafeEventSafeSetup(ctx, transaction, log)
	default:
		return nil, ethereum.ErrorUnsupportedEvent
	}

	return transfer, err
}

func (m *MultiSign) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	gnosisSafeFilterer, _ := safe.NewGnosisSafeFilterer(safe.AddressGnosisSafe, nil)

	return &MultiSign{
		gnosisSafeFilterer: gnosisSafeFilterer,
	}
}
