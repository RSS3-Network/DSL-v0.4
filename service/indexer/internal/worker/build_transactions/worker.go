package build_transactions

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

const Name = "buildTransaction"

var _ worker.Worker = (*service)(nil)

type service struct{}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkPolygon,           // source: alchemy, lens
		protocol.NetworkXDAI,              // source: blockscount
		protocol.NetworkCrossbell,         // source: kurora
		protocol.NetworkArbitrum,          // source: kurora
		protocol.NetworkOptimism,          // source: kurora
		protocol.NetworkAvalanche,         // source: kurora
		protocol.NetworkFantom,            // source: kurora
		protocol.NetworkCelo,              // source: kurora
		protocol.NetworkBinanceSmartChain, // source: moralis
		protocol.NetworkEthereum,          // source: ethereum
		protocol.NetworkBase,              // source: base
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (internalTransactions []model.Transaction, err error) {
	buildTransactions := lo.ToSlicePtr(transactions) // [] -> []*
	if buildTransactions, err = ethereum.BuildTransactions(ctx, message, buildTransactions); err != nil {
		loggerx.Global().Error("failed to build transactions", zap.Error(err))

		return nil, err
	}

	// []* -> []
	transactions = lo.Map(buildTransactions, func(x *model.Transaction, _ int) model.Transaction {
		return *x
	})

	return transactions, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{}
}

func New() worker.Worker {
	return &service{}
}
