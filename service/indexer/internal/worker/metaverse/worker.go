package metaverse

import (
	"context"

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
	for _, contract := range routerMap {
		network = append(network, contract.Network)
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("metaverse_worker")
	_, handlerSpan := tracer.Start(ctx, "metaverse_worker:Handle")

	defer handlerSpan.End()

	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		var contract MetaverseContract
		if contract = routerMap[transaction.AddressTo]; contract.Network != message.Network {
			return
		}

		for _, transfer := range transaction.Transfers {
			transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagMetaverse, transfer.Tag, contract.Type, transfer.Type)
			transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)

			if transfer.Tag == filter.TagMetaverse {
				transfer.Platform, transaction.Platform = contract.Platform, contract.Platform
			}
		}
	})

	return transactions, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{}
}

func New() worker.Worker {
	return &service{}
}
