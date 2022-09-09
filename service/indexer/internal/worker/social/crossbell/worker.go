package crossbell

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/handler"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const (
	Name     = "crossbell"
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
	opt := lop.NewOption().WithConcurrency(ethereum.RPCMaxConcurrency)

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

		// Get the raw data directly via Ethereum RPC node
		receipt, err := s.ethereumClient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
		if err != nil {
			logrus.Errorf("[crossbell worker] ethereumClient TransactionReceipt error, %v", err)
			return
		}

		internalTransfers, err := s.handleReceipt(ctx, message, transaction, receipt, transferMap)
		if err != nil {
			logrus.Errorf("[crossbell worker] handleReceipt error, %v", err)
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

		transaction.Owner = transaction.AddressFrom

		internalTransactions = append(internalTransactions, transaction)
	}, opt)

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
				loggerx.Global().Warn(
					"handle crossbell transfer failed",
					zap.Error(err),
					zap.String("worker", s.Name()),
					zap.String("network", message.Network),
					zap.String("address", message.Address),
					zap.String("transaction_hash", transaction.Hash),
				)
			}

			continue
		}

		internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(transfer.Tag, internalTransfer.Tag, transfer.Type, internalTransfer.Type)
		if internalTransfer.Tag == transfer.Tag {
			internalTransfer.Platform = transfer.Platform
		}

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
