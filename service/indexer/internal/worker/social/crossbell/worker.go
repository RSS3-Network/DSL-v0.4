package crossbell

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	contract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/worker/crossbell"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"

	lop "github.com/samber/lo/parallel"

	"go.opentelemetry.io/otel"

	"go.uber.org/zap"
)

const (
	Name     = "Crossbell"
	Endpoint = "https://rpc.crossbell.io"
)

var _ worker.Worker = (*service)(nil)

type service struct {
	crossbellClient *crossbell.Client
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkCrossbell,
	}
}

func (s *service) Initialize(ctx context.Context) (err error) {
	s.crossbellClient, err = crossbell.New()
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_crossbell")

	_, snap := tracer.Start(ctx, "worker_crossbell:Handle")

	defer snap.End()

	internalTransactions := make([]model.Transaction, 0)
	opt := lop.NewOption().WithConcurrency(ethereum.RPCMaxConcurrency)

	var mu sync.Mutex

	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		addressTo := common.HexToAddress(transaction.AddressTo)

		// Processing of transactions for contracts Profile and LinkList only
		if addressTo != contract.AddressCharacter && addressTo != contract.AddressLinkList {
			return
		}

		transaction.Platform = protocol.PlatformCrossbell

		// Retain the action model of the transfer type
		transferMap := make(map[int64]model.Transfer)

		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)
		transaction.Transfers = append(transaction.Transfers, transferMap[protocol.IndexVirtual])

		// Get the raw data directly via source data
		var sourceData ethereum.SourceData
		if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
			zap.L().Error("failed to unmarshal source data", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

			return
		}

		internalTransfers, err := s.crossbellClient.HandleReceipt(ctx, message, &transaction, sourceData.Receipt, transferMap)
		if err != nil {
			zap.L().Error("failed to handle receipt", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

			return
		}

		transaction.Transfers = append(transaction.Transfers, internalTransfers...)

		for _, transfer := range internalTransfers {
			transaction.Platform = transfer.Platform

			// Use first transfer as the main transfer
			continue
		}

		for _, transfer := range transaction.Transfers {
			transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
		}

		transaction.Owner = message.Address

		mu.Lock()
		internalTransactions = append(internalTransactions, transaction)
		mu.Unlock()
	}, opt)

	return internalTransactions, nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &service{}
}
