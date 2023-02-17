package staking

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/rss3"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"

	"go.uber.org/zap"
)

var _ worker.Worker = (*Staking)(nil)

type Staking struct {
	tokenClient *token.Client
}

func (s *Staking) Name() string {
	return "staking"
}

func (s *Staking) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
	}
}

func (s *Staking) Initialize(ctx context.Context) error {
	return nil
}

func (s *Staking) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	var (
		result = make([]model.Transaction, 0)
		locker sync.Mutex
	)

	waitGroup := sync.WaitGroup{}

	for _, transaction := range transactions {
		// Filter unsupported transactions
		if _, exists := platformMap[common.HexToAddress(transaction.AddressTo)]; !exists {
			continue
		}

		transaction := transaction
		transaction.Transfers = make([]model.Transfer, 0)

		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			var sourceData ethereum.SourceData
			if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
				zap.L().Error("unmarshal source data", zap.Error(err), zap.String("hash", transaction.Hash))

				return
			}

			for _, log := range sourceData.Receipt.Logs {
				transfer, err := s.handle(ctx, transaction, *log)
				if err != nil {
					// Ignore unimportant events
					if errors.Is(err, ethereum.ErrorUnsupportedEvent) {
						continue
					}

					zap.L().Error("handle staking event", zap.Error(err), zap.String("hash", transaction.Hash))

					return
				}

				transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
				transaction.Platform = transfer.Platform
				transaction.Transfers = append(transaction.Transfers, *transfer)
			}

			if transaction.Tag != "" && transaction.Type != "" {
				locker.Lock()
				defer locker.Unlock()

				result = append(result, transaction)
			}
		}()
	}

	waitGroup.Wait()

	return result, nil
}

func (s *Staking) handle(ctx context.Context, transaction model.Transaction, log types.Log) (transfer *model.Transfer, err error) {
	switch log.Topics[0] {
	case rss3.EventHashDeposited:
		transfer, err = s.handleRSS3StakingDeposited(ctx, transaction, log)
	case rss3.EventHashWithdrawn:
		transfer, err = s.handleRSS3StakingWithdrawn(ctx, transaction, log)
	case rss3.EventHashRewardsClaimed:
		transfer, err = s.handleRSS3StakingRewardsClaimed(ctx, transaction, log)
	default:
		return nil, ethereum.ErrorUnsupportedEvent
	}

	return transfer, err
}

func (s *Staking) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &Staking{
		tokenClient: token.New(),
	}
}
