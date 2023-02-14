package metaverse

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"

	"go.uber.org/zap"

	"github.com/naturalselectionlabs/pregod/internal/allowlist"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	lop "github.com/samber/lo/parallel"
	"go.opentelemetry.io/otel"
)

var (
	_       worker.Worker = (*service)(nil)
	network []string
)

type service struct{}

func (s *service) Name() string {
	return "metaverse"
}

func (s *service) Networks() []string {
	return network
}

func (s *service) Initialize(ctx context.Context) error {
	for _, contract := range RouterMap {
		network = append(network, contract.Network)
	}

	for address, name := range Agent {
		allowlist.AllowList.Add(address.String(), name)
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("metaverse_worker")
	_, handlerSpan := tracer.Start(ctx, "metaverse_worker:Handle")

	defer handlerSpan.End()

	internalTransactions := make([]model.Transaction, 0)

	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		var contract MetaverseContract
		if contract = RouterMap[common.HexToAddress(transaction.AddressTo)]; contract.Network != message.Network {
			return
		}

		var action MetaverseContract
		if events, exist := EventsRouterMap[common.HexToAddress(transaction.AddressTo)]; exist {
			var sourceData ethereum.SourceData
			if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
				zap.L().Error("failed to unmarshal source data", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

				return
			}

			for _, log := range sourceData.Receipt.Logs {
				// Filter anonymous log
				if len(log.Topics) == 0 {
					continue
				}

				for _, topic := range log.Topics {
					if _, exist = events[topic]; exist {
						action = events[topic]
						goto update
					}
				}
			}
		}

	update:
		for index := range transaction.Transfers {
			if action.Type != "" {
				transaction.Transfers[index].Tag, transaction.Transfers[index].Type = filter.UpdateTagAndType(filter.TagMetaverse, transaction.Transfers[index].Tag, action.Type, transaction.Transfers[index].Type)
				transaction.Transfers[index].Platform = action.Platform
			}

			// default type
			if contract.Type != "" {
				transaction.Transfers[index].Tag, transaction.Transfers[index].Type = filter.UpdateTagAndType(filter.TagMetaverse, transaction.Transfers[index].Tag, contract.Type, transaction.Transfers[index].Type)
			}

			transaction.Tag, transaction.Type = filter.UpdateTagAndType(transaction.Transfers[index].Tag, transaction.Tag, transaction.Transfers[index].Type, transaction.Type)

			if transaction.Transfers[index].Tag == filter.TagMetaverse {
				transaction.Transfers[index].Platform, transaction.Platform = contract.Platform, contract.Platform
			}
		}

		internalTransactions = append(internalTransactions, transaction)
	})

	return internalTransactions, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{}
}

func New() worker.Worker {
	return &service{}
}
