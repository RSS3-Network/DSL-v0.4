package ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/kurora/common/client/ethereum"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/gitcoin"
	mrc202 "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/mrc20"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"

	"go.uber.org/zap"
)

const (
	MaxConcurrency        = 200
	BaseNetworkSystemType = 126
)

var (
	ErrorUnsupportedEvent = errors.New("unsupported event")
	ErrorUnrelatedEvent   = errors.New("unrelated event")
)

func BuildTransactions(ctx context.Context, message *protocol.Message, transactions []*model.Transaction) ([]*model.Transaction, error) {
	tracer := otel.Tracer("ethereum")
	ctx, span := tracer.Start(ctx, "ethereum:BuildTransactions")

	var err error

	defer opentelemetry.Log(span, message, transactions, err)

	result := make([]*model.Transaction, 0)

	for _, transactionsSplit := range lo.Chunk(transactions, 2*MaxConcurrency) {
		blocks, err := lop.MapWithError(transactionsSplit, makeBlockHandlerFunc(ctx, message), lop.NewOption().WithConcurrency(MaxConcurrency))
		if err != nil {
			loggerx.Global().Error("failed to handle blocks", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address))

			return nil, err
		}

		blocks = lo.Filter(blocks, func(block *ethereum.Block, i int) bool {
			return block != nil && block.Number != nil
		})

		blockMap := make(map[int64]*ethereum.Block)
		for _, block := range blocks {
			blockMap[block.Number.Int64()] = block
		}

		// Error topic/field count mismatch
		transactionsSplit, err = lop.MapWithError(transactionsSplit, makeTransactionHandlerFunc(ctx, message, blockMap), lop.NewOption().WithConcurrency(MaxConcurrency))
		if err != nil {
			loggerx.Global().Error("failed to handle transactions", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address))

			return nil, err
		}

		if transactionsSplit != nil {
			result = append(result, lo.Filter(transactionsSplit, func(transaction *model.Transaction, i int) bool {
				return transaction != nil
			})...)
		}
	}

	return result, nil
}

func makeBlockHandlerFunc(ctx context.Context, message *protocol.Message) func(transaction *model.Transaction, i int) (*ethereum.Block, error) {
	return func(transaction *model.Transaction, i int) (*ethereum.Block, error) {
		ethereumClient, err := ethclientx.GlobalX(message.Network)
		if err != nil {
			return nil, fmt.Errorf("get global ethereum client failed: %w", err)
		}

		block, err := ethereumClient.BlockByNumber(ctx, big.NewInt(transaction.BlockNumber))
		if err != nil {
			zap.L().Error("failed to get block and start to check receipt", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address), zap.Int64("block_number", transaction.BlockNumber))

			receipt, err := ethereumClient.ReceiptByHash(ctx, common.HexToHash(transaction.Hash))
			if err != nil {
				zap.L().Error("failed to get receipt", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address), zap.Int64("block_number", transaction.BlockNumber))

				return nil, err
			}

			if block, err = ethereumClient.BlockByNumber(ctx, receipt.BlockNumber); err != nil {
				zap.L().Error("failed to get block finally!!", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address), zap.Int64("block_number", transaction.BlockNumber))

				return nil, err
			}

			transaction.BlockNumber = block.Number.Int64()
		}

		if block == nil || block.Number == nil {
			loggerx.Global().Error("block is nil", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address), zap.Int64("block_number", transaction.BlockNumber))
		}
		return block, nil
	}
}

func makeTransactionHandlerFunc(ctx context.Context, message *protocol.Message, blockMap map[int64]*ethereum.Block) func(transaction *model.Transaction, i int) (*model.Transaction, error) {
	return func(transaction *model.Transaction, i int) (*model.Transaction, error) {
		block, exists := blockMap[transaction.BlockNumber]
		if !exists {
			return nil, nil
		}

		transaction.Timestamp = time.Unix(int64(block.Timestamp), 0)

		// Find transaction in block
		internalTransaction, exists := lo.Find(block.Transactions, func(blockTransaction *ethereum.Transaction) bool {
			return strings.EqualFold(blockTransaction.Hash.String(), transaction.Hash)
		})
		if !exists {
			return nil, nil
		}

		// No signature verification
		transaction.AddressFrom = strings.ToLower(internalTransaction.From.String())
		if internalTransaction.To == nil {
			transaction.AddressTo = strings.ToLower(AddressGenesis.String())
		} else {
			transaction.AddressTo = strings.ToLower(internalTransaction.To.String())
		}

		// Default owner
		transaction.Owner = transaction.AddressFrom

		// crawler message address is nil
		if transaction.Source != protocol.SourceKurora && transaction.AddressFrom != "" && message.Address != "" && !strings.EqualFold(transaction.AddressFrom, message.Address) &&
			!allowlist.AllowList.Contains(transaction.AddressFrom) && !allowlist.AllowList.Contains(transaction.AddressTo) {
			return nil, nil
		}

		ethereumClient, err := ethclientx.GlobalX(message.Network)
		if err != nil {
			return nil, fmt.Errorf("get global ethereum client failed: %w", err)
		}

		receipt, err := ethereumClient.ReceiptByHash(ctx, internalTransaction.Hash)
		if err != nil {
			loggerx.Global().Error("failed to get transaction receipt", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address), zap.String("hash", transaction.Hash))

			return nil, err
		}

		// Compatible with Arbitrum Transaction types
		// https://github.com/OffchainLabs/go-ethereum/blob/01cc043469b25ca03a296be256aae2082c161bd7/core/types/transaction.go#L43-L55
		if lo.Contains([]uint64{
			100, 101, 102, 104, 105, 106, 120,
		}, internalTransaction.Type) {
			internalTransaction.Type = types.LegacyTxType
		}

		switch internalTransaction.Type {
		case types.LegacyTxType, types.AccessListTxType:
			fee := decimal.NewFromBigInt(internalTransaction.GasPrice, 0).Mul(decimal.NewFromInt(int64(receipt.GasUsed))).Shift(-18)
			transaction.Fee = &fee
		case types.DynamicFeeTxType:
			if block.BaseFee == nil {
				block.BaseFee = big.NewInt(0)
				loggerx.Global().Error("block baseFee is nil", zap.String("network", message.Network), zap.String("address", message.Address), zap.String("hash", transaction.Hash))
			}

			if internalTransaction.GasTipCap == nil {
				internalTransaction.GasTipCap = big.NewInt(0)
				loggerx.Global().Error("gasTipCap is nil", zap.String("network", message.Network), zap.String("address", message.Address), zap.String("hash", transaction.Hash))
			}

			fee := (decimal.NewFromBigInt(block.BaseFee, 0).Add(decimal.NewFromBigInt(internalTransaction.GasTipCap, 0))).Mul(decimal.NewFromInt(int64(receipt.GasUsed))).Shift(-18)
			transaction.Fee = &fee
		case BaseNetworkSystemType:
			if transaction.Network != protocol.NetworkBase {
				// L2
				return nil, fmt.Errorf("unsupported transaction type %d: %s", internalTransaction.Type, internalTransaction.Hash)
			}

			return nil, nil
		default:
			// L2
			return nil, fmt.Errorf("unsupported transaction type %d: %s", internalTransaction.Type, internalTransaction.Hash)
		}

		// Transaction compatibility conversion for Arbitrum
		raw, err := json.Marshal(internalTransaction)
		if err != nil {
			return nil, fmt.Errorf("marshal transaction failed: %w", err)
		}

		var originTransaction types.Transaction
		if err := json.Unmarshal(raw, &originTransaction); err != nil {
			return nil, fmt.Errorf("unmarshal transaction failed: %w", err)
		}

		raw, err = json.Marshal(receipt)
		if err != nil {
			return nil, fmt.Errorf("marshal receipt failed: %w", err)
		}

		var originReceipt types.Receipt
		if err := json.Unmarshal(raw, &originReceipt); err != nil {
			return nil, fmt.Errorf("unmarshal transaction failed: %w", err)
		}

		// Handle receipt status
		transactionSuccess := receipt.Status == types.ReceiptStatusSuccessful
		transaction.Success = &transactionSuccess

		// Build source data
		sourceData := SourceData{
			Transaction: &originTransaction,
			Receipt:     &originReceipt,
		}

		if transaction.SourceData, err = json.Marshal(sourceData); err != nil {
			return nil, fmt.Errorf("marshal source data failed: %w", err)
		}

		if transaction.Transfers, err = handleReceipt(ctx, message, transaction, &originReceipt); err != nil {
			loggerx.Global().Error("failed to handle receipt", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address))

			return nil, err
		}

		transaction.Transfers = append(transaction.Transfers, model.Transfer{
			// This is a virtual transfer
			TransactionHash: transaction.Hash,
			Timestamp:       transaction.Timestamp,
			Index:           protocol.IndexVirtual,
			AddressFrom:     transaction.AddressFrom,
			AddressTo:       transaction.AddressTo,
			Metadata:        metadata.Default,
			Network:         message.Network,
			Source:          transaction.Source,
			SourceData:      transaction.SourceData,
			RelatedUrls: []string{
				BuildScanURL(message.Network, transaction.Hash),
			},
		})

		return transaction, nil
	}
}

func handleReceipt(ctx context.Context, message *protocol.Message, transaction *model.Transaction, receipt *types.Receipt) ([]model.Transfer, error) {
	tracer := otel.Tracer("ethereum")
	ctx, trace := tracer.Start(ctx, "ethereum:handleReceipt")
	transfers := make([]model.Transfer, 0)
	var err error
	defer func() { opentelemetry.Log(trace, message, transfers, err) }()

	for _, log := range receipt.Logs {
		transfer, err := handleLog(ctx, message, transaction, *log)
		if err != nil {
			if errors.Is(err, ErrorUnsupportedEvent) || errors.Is(err, ErrorUnrelatedEvent) {
				continue
			}

			return nil, err
		}

		transfers = append(transfers, *transfer)
	}

	return transfers, nil
}

func handleLog(ctx context.Context, message *protocol.Message, transaction *model.Transaction, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("ethereum")
	_, trace := tracer.Start(ctx, "ethereum:handleLog")

	transfer := model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		Index:           int64(log.Index),
		Network:         transaction.Network,
		Metadata:        metadata.Default,
		Source:          transaction.Source,
		RelatedUrls: []string{
			BuildScanURL(message.Network, transaction.Hash),
		},
	}

	var err error
	defer opentelemetry.Log(trace, message, transfer, err)

	if len(log.Topics) == 0 {
		return nil, ErrorUnsupportedEvent
	}

	switch log.Topics[0] {
	case erc20.EventHashTransfer, erc721.EventHashTransfer:
		switch len(log.Topics) {
		case 3:
			filterer, err := erc20.NewERC20Filterer(log.Address, nil)
			if err != nil {
				return nil, err
			}

			event, err := filterer.ParseTransfer(log)
			if err != nil {
				return nil, err
			}

			transfer.AddressFrom = strings.ToLower(event.From.String())
			transfer.AddressTo = strings.ToLower(event.To.String())
		case 4:
			filterer, err := erc721.NewERC721Filterer(log.Address, nil)
			if err != nil {
				return nil, err
			}

			event, err := filterer.ParseTransfer(log)
			if err != nil {
				return nil, err
			}

			// remove burn when from is not tx owner
			if event.To == ethereum.AddressGenesis && !strings.EqualFold(event.From.String(), transaction.Owner) {
				return nil, ErrorUnsupportedEvent
			}

			transfer.AddressFrom = strings.ToLower(event.From.String())
			transfer.AddressTo = strings.ToLower(event.To.String())
		default:
			return nil, ErrorUnsupportedEvent
		}
	case erc1155.EventHashTransferSingle:
		filterer, err := erc1155.NewERC1155Filterer(log.Address, nil)
		if err != nil {
			return nil, err
		}

		event, err := filterer.ParseTransferSingle(log)
		if err != nil {
			return nil, err
		}

		transfer.AddressFrom = strings.ToLower(event.From.String())
		transfer.AddressTo = strings.ToLower(event.To.String())
	case erc1155.EventHashTransferBatch:
		filterer, err := erc1155.NewERC1155Filterer(log.Address, nil)
		if err != nil {
			return nil, err
		}

		event, err := filterer.ParseTransferBatch(log)
		if err != nil {
			return nil, err
		}

		transfer.AddressFrom = strings.ToLower(event.From.String())
		transfer.AddressTo = strings.ToLower(event.To.String())
	case mrc202.EventHashLogTransfer:
		filterer, err := mrc202.NewMRC20Filterer(log.Address, nil)
		if err != nil {
			return nil, err
		}

		event, err := filterer.ParseLogTransfer(log)
		if err != nil {
			return nil, err
		}

		transfer.AddressFrom = strings.ToLower(event.From.String())
		transfer.AddressTo = strings.ToLower(event.To.String())
	case gitcoin.EventHashDonationSent:
		flterer, err := gitcoin.NewGitcoinFilterer(log.Address, nil)
		if err != nil {
			return nil, err
		}

		event, err := flterer.ParseDonationSent(log)
		if err != nil {
			return nil, err
		}

		transfer.AddressFrom = strings.ToLower(event.Donor.String())
		transfer.AddressTo = strings.ToLower(event.Dest.String())
	case erc20.EventHashApproval: // ERC-20 and ERC-721
		switch {
		case len(log.Topics) == 3:
			filterer, err := erc20.NewERC20Filterer(log.Address, nil)
			if err != nil {
				return nil, fmt.Errorf("create erc20 filterer: %w", err)
			}

			event, err := filterer.ParseApproval(log)
			if err != nil {
				return nil, fmt.Errorf("parse approval event: %w", err)
			}

			transfer.AddressFrom = strings.ToLower(event.Owner.String())
			transfer.AddressTo = strings.ToLower(event.Spender.String())
		case len(log.Topics) == 4:
			filterer, err := erc721.NewERC721Filterer(log.Address, nil)
			if err != nil {
				return nil, fmt.Errorf("create erc20 filterer: %w", err)
			}

			event, err := filterer.ParseApproval(log)
			if err != nil {
				return nil, fmt.Errorf("parse approval event: %w", err)
			}

			transfer.AddressFrom = strings.ToLower(event.Owner.String())
			transfer.AddressTo = strings.ToLower(event.Spender.String())
		}
	case erc721.EventHashApprovalForAll: // ERC-721 and ERC-1155
		filterer, err := erc721.NewERC721Filterer(log.Address, nil)
		if err != nil {
			return nil, fmt.Errorf("create erc721 filterer: %w", err)
		}

		event, err := filterer.ParseApprovalForAll(log)
		if err != nil {
			return nil, fmt.Errorf("parse approval for all event: %w", err)
		}

		transfer.AddressFrom = strings.ToLower(event.Owner.String())
		transfer.AddressTo = strings.ToLower(event.Operator.String())
	default:
		return nil, ErrorUnsupportedEvent
	}

	address := strings.ToLower(message.Address)

	// Ignore transfers that are not related to message.address
	if len(address) > 0 && transaction.AddressFrom != address && transfer.AddressTo != address && transfer.AddressFrom != address {
		return nil, ErrorUnrelatedEvent
	}

	if transfer.SourceData, err = json.Marshal(log); err != nil {
		return nil, err
	}

	return &transfer, nil
}
