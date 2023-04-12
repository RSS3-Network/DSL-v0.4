package name_service

import (
	"context"
	"encoding/json"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shopspring/decimal"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

var _ worker.Worker = (*internal)(nil)

type internal struct {
	tokenClient *token.Client
}

func (i *internal) Name() string {
	return filter.CollectibleEdit
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
	}
}

func (i *internal) Initialize(ctx context.Context) error {
	return nil
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	var (
		result    = make([]model.Transaction, 0)
		waitGroup sync.WaitGroup
		locker    sync.Mutex
	)

	for _, transaction := range transactions {
		// Unsupported Platform
		platform, exists := platformMap[common.HexToAddress(transaction.AddressTo)]
		if !exists {
			continue
		}

		// Initialize the transaction
		transaction := transaction
		transaction.Transfers = make([]model.Transfer, 0)
		transaction.Platform = platform

		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			var sourceData ethereum.SourceData
			if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
				zap.L().Warn("unmarshal source data", zap.Error(err), zap.String("source_data", string(transaction.SourceData)))

				return
			}

			for _, log := range sourceData.Receipt.Logs {
				// Filter anonymous log
				if len(log.Topics) == 0 {
					continue
				}

				var (
					transfer *model.Transfer
					err      error
				)

				switch log.Topics[0] {
				// ENS NameRenewed
				case ens.EventNameRenewed:
					transfer, err = i.handleENSNameRenewed(ctx, message, transaction, *log, platform)
				default:
					continue
				}

				if err != nil {
					zap.L().Error("handle event", zap.Error(err), zap.Stringer("topic_first", log.Topics[0]))

					return
				}

				transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagCollectible, transfer.Tag, i.Name(), transfer.Type)

				transaction.Tag, transaction.Type = filter.UpdateTagAndType(filter.TagCollectible, transaction.Tag, i.Name(), transaction.Type)
				transaction.Owner = transaction.AddressFrom
				transaction.Transfers = append(transaction.Transfers, *transfer)
			}

			locker.Lock()
			defer locker.Unlock()

			result = append(result, transaction)
		}()
	}

	waitGroup.Wait()

	return result, nil
}

func (i *internal) buildNameServiceMetadata(name string, action string, tokenType *token.Native, cost *big.Int, expires *big.Int, key, value string) (json.RawMessage, error) {
	nameServiceMetadata := metadata.NameService{
		Action: action,
		Name:   name,
	}

	if expires != nil {
		nameServiceMetadata.Expiry = lo.ToPtr(time.Unix(expires.Int64(), 0))
	}

	if tokenType != nil && cost != nil {
		internalTokenValue := decimal.NewFromBigInt(cost, 0)

		internalTokenDisplay := internalTokenValue.Shift(-int32(tokenType.Decimals))

		standard := protocol.TokenStandardNative

		nameServiceMetadata.Cost = &metadata.Token{
			Name:         tokenType.Name,
			Symbol:       tokenType.Symbol,
			Decimals:     tokenType.Decimals,
			Image:        tokenType.Logo,
			Standard:     standard,
			Value:        &internalTokenValue,
			ValueDisplay: &internalTokenDisplay,
		}
	}

	if key != "" {
		nameServiceMetadata.Key = key
		nameServiceMetadata.Value = value
	}

	return json.Marshal(&nameServiceMetadata)
}

func (i *internal) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &internal{
		tokenClient: token.New(),
	}
}
