package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	bridge "github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/arbitrum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/optimism"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/polygon"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shopspring/decimal"
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

func (w *Worker) fillTransactionMetadata(transaction model.Transaction, transfer model.Transfer) model.Transaction {
	transaction.Owner = transfer.AddressFrom
	transaction.Platform = transfer.Platform
	transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)

	return transaction
}

func (w *Worker) buildTransfer(ctx context.Context, transaction model.Transaction, log types.Log, from, to common.Address, platform string, chainID uint64, tokenAddress *common.Address, tokenValue *big.Int, transferType string) (*model.Transfer, error) {
	var (
		tokenMetadata *metadata.Token
		err           error
	)

	if tokenAddress == nil {
		tokenMetadata, err = w.tokenClient.NatvieToMetadata(ctx, transaction.Network)
		if err != nil {
			return nil, fmt.Errorf("failed to get token metadata: %w", err)
		}
	} else {
		tokenMetadata, err = w.tokenClient.ERC20ToMetadata(ctx, transaction.Network, tokenAddress.String())
		if err != nil {
			return nil, fmt.Errorf("failed to get token metadata: %w", err)
		}
	}

	internalTokenValue := decimal.NewFromBigInt(tokenValue, 0)
	tokenMetadata.Value = &internalTokenValue
	tokenDisplay := internalTokenValue.Shift(-int32(tokenMetadata.Decimals))
	tokenMetadata.ValueDisplay = &tokenDisplay

	network, exists := networkMap[chainID]
	if !exists {
		return nil, fmt.Errorf("unsupported chain id: %d", chainID)
	}

	metadataRaw, err := json.Marshal(bridge.Bridge{
		Action:        transferType,
		TargetNetwork: network,
		Token:         *tokenMetadata,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Tag:             filter.TagTransaction,
		Type:            filter.TransactionBridge,
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(from.String()),
		AddressTo:       strings.ToLower(to.String()),
		Metadata:        metadataRaw,
		Network:         transaction.Network,
		Platform:        platform,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func New() worker.Worker {
	return &Worker{
		tokenClient: token.New(),
	}
}
