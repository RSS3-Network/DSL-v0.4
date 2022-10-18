package bridge

import (
	"context"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/arbitrum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/optimism"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/polygon"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.uber.org/zap"
)

var _ worker.Worker = (*Worker)(nil)

type Worker struct {
	tokenClient *token.Client
}

func (w *Worker) Name() string {
	return "bridge"
}

func (w *Worker) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
	}
}

func (w *Worker) Initialize(ctx context.Context) error {
	return nil
}

func (w *Worker) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		var (
			internalTransaction *model.Transaction
			err                 error
		)

		// Polygon Bridge
		if transaction.Network == protocol.NetworkEthereum && (strings.EqualFold(transaction.AddressTo, polygon.AddressBridge.String()) || strings.EqualFold(transaction.AddressFrom, polygon.AddressBridge.String())) {
			if internalTransaction, err = w.handlePolygon(ctx, transaction); err != nil {
				zap.L().Error("failed to handle polygon transaction", zap.Error(err))

				continue
			}
		}

		// Optimism Bridge
		if transaction.Network == protocol.NetworkEthereum && (strings.EqualFold(transaction.AddressTo, optimism.AddressGateway.String()) || strings.EqualFold(transaction.AddressFrom, optimism.AddressGateway.String())) {
			if internalTransaction, err = w.handleOptimism(ctx, transaction); err != nil {
				zap.L().Error("failed to handle optimism transaction", zap.Error(err))

				continue
			}
		}

		// Arbitrum One Bridge
		if transaction.Network == protocol.NetworkEthereum && (strings.EqualFold(transaction.AddressTo, arbitrum.AddressInboxOne.String()) || strings.EqualFold(transaction.AddressFrom, arbitrum.AddressInboxOne.String())) {
			if internalTransaction, err = w.handleArbitrum(ctx, transaction); err != nil {
				zap.L().Error("failed to handle arbitrum transaction", zap.Error(err))

				continue
			}
		}

		// Arbitrum Nova Bridge
		if transaction.Network == protocol.NetworkEthereum && (strings.EqualFold(transaction.AddressTo, arbitrum.AddressInboxNova.String()) || strings.EqualFold(transaction.AddressFrom, arbitrum.AddressInboxNova.String())) {
			if internalTransaction, err = w.handleArbitrum(ctx, transaction); err != nil {
				zap.L().Error("failed to handle arbitrum transaction", zap.Error(err))

				continue
			}
		}

		if internalTransaction != nil {
			internalTransactions = append(internalTransactions, *internalTransaction)
		}
	}

	return internalTransactions, nil
}

func (w *Worker) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &Worker{}
}
