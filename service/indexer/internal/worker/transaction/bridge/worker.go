package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	bridge "github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/arbitrum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/hop"
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
		protocol.NetworkBinanceSmartChain,
		protocol.NetworkPolygon,
		protocol.NetworkXDAI,
		protocol.NetworkCrossbell,
		protocol.NetworkOptimism,
		protocol.NetworkAvalanche,
		protocol.NetworkCelo,
		protocol.NetworkFantom,
	}
}

func (w *Worker) Initialize(ctx context.Context) error {
	return nil
}

func (w *Worker) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	var (
		internalTransactions = make([]model.Transaction, 0)
		locker               sync.Mutex
	)

	waitGroup := sync.WaitGroup{}

	for _, transaction := range transactions {
		platform, exists := platformMap[common.HexToAddress(transaction.AddressTo)]
		if !exists {
			continue
		}

		var (
			internalTransaction = transaction
			internalTransfer    *model.Transfer
			err                 error
		)

		internalTransaction.Transfers = make([]model.Transfer, 0)

		waitGroup.Add(1)

		go func() {
			var sourceData ethereum.SourceData
			if err := json.Unmarshal(internalTransaction.SourceData, &sourceData); err != nil {
				zap.L().Warn("unmarshal source data", zap.Error(err), zap.String("hash", internalTransaction.Hash))

				return
			}

			for _, log := range sourceData.Receipt.Logs {
				switch log.Topics[0] {
				case hop.EventTransferSentToL2:
					internalTransfer, err = w.handleHopTransferSentToL2(ctx, internalTransaction, *log)
				case hop.EventTransferSent:
					internalTransfer, err = w.handleHopTransferSent(ctx, internalTransaction, *log)
				case arbitrum.EventHashMessageDelivered:
					internalTransfer, err = w.handleArbitrumHashMessageDelivered(ctx, internalTransaction, *log)
				case optimism.EventHashETHDepositInitiated:
					internalTransfer, err = w.handleOptimismETHDepositInitiated(ctx, internalTransaction, *log)
				case optimism.EventHashERC20DepositInitiated:
					internalTransfer, err = w.handleOptimismERC20DepositInitiated(ctx, internalTransaction, *log)
				case optimism.EventHashETHWithdrawalFinalized:
					internalTransfer, err = w.handleOptimismETHWithdrawalFinalized(ctx, internalTransaction, *log)
				case optimism.EventHashERC20WithdrawalFinalized:
					internalTransfer, err = w.handleOptimismERC20WithdrawalFinalized(ctx, internalTransaction, *log)
				case polygon.EventHashLockedEther:
					internalTransfer, err = w.handlePolygonLockedEther(ctx, internalTransaction, *log)
				case polygon.EventHashLockedERC20:
					internalTransfer, err = w.handlePolygonLockedERC20(ctx, internalTransaction, *log)
				case polygon.EventHashLockedMintableERC20:
					internalTransfer, err = w.handlePolygonLockedMintableERC20(ctx, internalTransaction, *log)
				case polygon.EventHashExitedEther:
					internalTransfer, err = w.handlePolygonExitedEther(ctx, internalTransaction, *log)
				default:
					continue
				}

				if err != nil {
					zap.L().Warn("handle transfer", zap.Error(err), zap.String("hash", internalTransaction.Hash))

					return
				}

				w.buildTransactionMetadata(&internalTransaction, *internalTransfer, platform)
				internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
			}

			if len(internalTransaction.Transfers) > 0 {
				locker.Lock()
				defer locker.Unlock()

				internalTransactions = append(internalTransactions, internalTransaction)
			}

			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	return internalTransactions, nil
}

func (w *Worker) buildTransactionMetadata(transaction *model.Transaction, transfer model.Transfer, platform string) {
	transaction.Owner = transfer.AddressFrom
	transaction.Platform = platform
	transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
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

func (w *Worker) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &Worker{
		tokenClient: token.New(),
	}
}
