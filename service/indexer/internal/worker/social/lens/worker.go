package lens

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	lens_comm "github.com/naturalselectionlabs/pregod/common/worker/lens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	lop "github.com/samber/lo/parallel"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ worker.Worker = (*service)(nil)

type service struct {
	commWorkerClient *lens_comm.Client
}

func (s *service) Name() string {
	return protocol.PlatformLens
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkPolygon,
	}
}

func (s *service) Initialize(ctx context.Context) (err error) {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (internalTransactions []model.Transaction, err error) {
	tracer := otel.Tracer("worker_lens")
	_, trace := tracer.Start(ctx, "worker_len:Handle")

	defer func() { opentelemetry.Log(trace, transactions, internalTransactions, err) }()

	var mu sync.Mutex
	opt := lop.NewOption().WithConcurrency(ethereum.RPCMaxConcurrency)
	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		addressTo := common.HexToAddress(transaction.AddressTo)
		if transaction.Source != protocol.SourceKurora && transaction.Source != protocol.SourceAlchemy &&
			addressTo != lens.HubProxyContractAddress && addressTo != lens.ProfileProxyContractAddress {
			return
		}

		transferMap := make(map[int64]model.Transfer)
		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// get receipt
		internalTransfers, err := s.commWorkerClient.HandleReceipt(ctx, &transaction)
		if err != nil {
			zap.L().Error("failed to handle receipt", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

			return
		}

		if len(internalTransfers) > 0 {
			//nolint:gocritic
			transaction.Transfers = append(internalTransfers, transferMap[protocol.IndexVirtual])
			transaction.Platform = protocol.PlatformLens

			for _, transfer := range transaction.Transfers {
				transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
			}
		}

		mu.Lock()
		internalTransactions = append(internalTransactions, transaction)
		mu.Unlock()
	}, opt)

	return internalTransactions, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{}
}

func New() worker.Worker {
	return &service{
		commWorkerClient: lens_comm.New(),
	}
}
