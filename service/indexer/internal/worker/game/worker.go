package game

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
	_           worker.Worker = (*service)(nil)
	gameNetwork []string
)

type service struct{}

func (s *service) Name() string {
	return "game"
}

func (s *service) Networks() []string {
	return gameNetwork
}

func (s *service) Initialize(ctx context.Context) error {
	for _, game := range routerMap {
		gameNetwork = append(gameNetwork, game.Network)
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("game_worker")
	_, handlerSpan := tracer.Start(ctx, "game_worker:Handle")

	defer handlerSpan.End()

	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		var gameContract GameContract
		if gameContract = routerMap[transaction.AddressTo]; gameContract.Network != message.Network {
			return
		}

		for _, transfer := range transaction.Transfers {
			transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagGame, transfer.Tag, gameContract.Type, transfer.Type)
			transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)

			if transfer.Tag == filter.TagGame {
				transfer.Platform, transaction.Platform = gameContract.Platform, gameContract.Platform
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
