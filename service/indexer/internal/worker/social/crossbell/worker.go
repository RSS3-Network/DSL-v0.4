package crossbell

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/handler"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const (
	Name = "crossbell"

	Endpoint = "https://rpc.crossbell.io"
)

var _ worker.Worker = (*service)(nil)

type service struct {
	ethereumClient *ethclient.Client
	handler        handler.Interface
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
	if s.ethereumClient, err = ethclient.Dial(Endpoint); err != nil {
		return err
	}

	if s.handler, err = handler.New(s.ethereumClient); err != nil {
		return err
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_crossbell")

	_, snap := tracer.Start(ctx, "worker_crossbell:Handle")

	defer snap.End()

	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		addressTo := common.HexToAddress(transaction.AddressTo)

		// Processing of transactions for contracts Profile and LinkList only
		if addressTo != contract.AddressCharacter && addressTo != contract.AddressLinkList {
			continue
		}

		transaction.Platform = s.Name()

		// Retain the action model of the transfer type
		transferMap := make(map[int64]model.Transfer)

		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)
		transaction.Transfers = append(transaction.Transfers, transferMap[protocol.IndexVirtual])

		// Get the raw data directly via Ethereum RPC node
		receipt, err := s.ethereumClient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
		if err != nil {
			return nil, err
		}

		internalTransfers, err := s.handleReceipt(ctx, message, transaction, receipt, transferMap)
		if err != nil {
			return nil, err
		}

		transaction.Transfers = append(transaction.Transfers, internalTransfers...)

		for _, transfer := range transaction.Transfers {
			transaction.Tag = filter.UpdateTag(transfer.Tag, transaction.Tag)
		}

		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (s *service) handleReceipt(ctx context.Context, message *protocol.Message, transaction model.Transaction, receipt *types.Receipt, transferMap map[int64]model.Transfer) ([]model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell")

	_, snap := tracer.Start(ctx, "worker_crossbell:handleReceipt")

	defer snap.End()

	internalTransfers := make([]model.Transfer, 0)

	for _, log := range receipt.Logs {
		logIndex := int64(log.Index)

		transfer, exist := transferMap[logIndex]
		if !exist {
			sourceData, err := json.Marshal(log)
			if err != nil {
				return nil, err
			}

			// Virtual transfer
			transfer = model.Transfer{
				TransactionHash: transaction.Hash,
				Timestamp:       transaction.Timestamp,
				Index:           logIndex,
				AddressFrom:     transaction.AddressFrom,
				AddressTo:       transaction.AddressTo,
				Metadata:        metadata.Default,
				Network:         message.Network,
				Source:          protocol.SourceOrigin,
				SourceData:      sourceData,
			}
		}

		internalTransfer, err := s.handler.Handle(ctx, transaction, transfer)
		if err != nil {
			if !errors.Is(err, contract.ErrorUnknownEvent) {
				logger.Global().Warn(
					"handle crossbell transfer failed",
					zap.Error(err),
					zap.String("worker", s.Name()),
					zap.String("network", message.Network),
					zap.String("address", message.Address),
				)
			}

			continue
		}

		if transfer.Platform != "" {
			internalTransfer.Platform = transfer.Platform
		}

		internalTransfer.Tag = filter.UpdateTag(internalTransfer.Tag, transfer.Tag)
		internalTransfers = append(internalTransfers, *internalTransfer)
	}

	return internalTransfers, nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &service{}
}
