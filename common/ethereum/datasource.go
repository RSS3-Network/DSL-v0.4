package ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/mrc20"
	"github.com/naturalselectionlabs/pregod/common/logger"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	lop "github.com/samber/lo/parallel"
	"go.uber.org/zap"
)

const (
	Source = "ethereum"

	MaxConcurrency = 200
)

var (
	ErrorUnsupportedEvent = errors.New("unsupported event")
	ErrorUnrelatedEvent   = errors.New("unrelated event")
)

func BuildTransactions(ctx context.Context, message *protocol.Message, transactions []*model.Transaction, ethereumClient *ethclient.Client) ([]*model.Transaction, error) {
	blocks, err := lop.MapWithError(transactions, makeBlockHandlerFunc(ctx, message, ethereumClient), lop.NewOption().WithConcurrency(MaxConcurrency))
	if err != nil {
		return nil, err
	}

	blockMap := make(map[int64]*types.Block)
	for _, block := range blocks {
		blockMap[block.Number().Int64()] = block
	}

	// Error topic/field count mismatch
	transactions, err = lop.MapWithError(transactions, makeTransactionHandlerFunc(ctx, message, ethereumClient, blockMap), lop.NewOption().WithConcurrency(MaxConcurrency))
	if err != nil {
		logger.Global().Error("failed to handle transaction", zap.Error(err))
	}

	return transactions, err
}

func makeBlockHandlerFunc(ctx context.Context, message *protocol.Message, ethereumClient *ethclient.Client) func(transaction *model.Transaction, i int) (*types.Block, error) {
	return func(transaction *model.Transaction, i int) (*types.Block, error) {
		block, err := ethereumClient.BlockByNumber(ctx, big.NewInt(transaction.BlockNumber))
		if err != nil {
			return nil, err
		}

		return block, nil
	}
}

func makeTransactionHandlerFunc(ctx context.Context, message *protocol.Message, ethereumClient *ethclient.Client, blockMap map[int64]*types.Block) func(transaction *model.Transaction, i int) (*model.Transaction, error) {
	return func(transaction *model.Transaction, i int) (*model.Transaction, error) {
		block := blockMap[transaction.BlockNumber]

		transaction.Timestamp = time.Unix(int64(block.Time()), 0)

		for index, blockTransaction := range block.Transactions() {
			if blockTransaction.Hash().String() == transaction.Hash {
				transaction.Index = int64(index)

				break
			}
		}

		internalTransaction, _, err := ethereumClient.TransactionByHash(ctx, common.HexToHash(transaction.Hash))
		if err != nil {
			return nil, err
		}

		var transactionMessage types.Message

		switch internalTransaction.Type() {
		case types.LegacyTxType:
			transactionMessage, err = internalTransaction.AsMessage(types.NewEIP155Signer(internalTransaction.ChainId()), nil)
		case types.DynamicFeeTxType:
			transactionMessage, err = internalTransaction.AsMessage(types.LatestSignerForChainID(internalTransaction.ChainId()), nil)
		default:
			err = ErrorUnsupportedTransactionType
		}

		if err != nil {
			return nil, err
		}

		transaction.AddressFrom = strings.ToLower(transactionMessage.From().String())

		addressTo := AddressGenesis.String()

		if internalTransaction.To() != nil {
			addressTo = internalTransaction.To().String()
		}

		transaction.AddressTo = strings.ToLower(addressTo)

		receipt, err := ethereumClient.TransactionReceipt(ctx, internalTransaction.Hash())
		if err != nil {
			return nil, err
		}

		transactionSuccess := receipt.Status == types.ReceiptStatusSuccessful
		transaction.Success = &transactionSuccess

		transaction.Source = Source

		if transaction.SourceData, err = json.Marshal(&SourceData{
			Transaction: internalTransaction,
			Receipt:     receipt,
		}); err != nil {
			return nil, err
		}

		if transaction.Transfers, err = handleReceipt(ctx, message, transaction, receipt); err != nil {
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
			Source:          Source,
			SourceData:      transaction.SourceData,
			RelatedUrls: []string{
				BuildScanURL(message.Network, transaction.Hash),
			},
		})

		return transaction, nil
	}
}

func handleReceipt(ctx context.Context, message *protocol.Message, transaction *model.Transaction, receipt *types.Receipt) ([]model.Transfer, error) {
	transfers := make([]model.Transfer, 0)

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
	transfer := model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		Index:           int64(log.Index),
		Network:         transaction.Network,
		Metadata:        metadata.Default,
		Source:          Source,
		RelatedUrls: []string{
			BuildScanURL(message.Network, transaction.Hash),
		},
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
	case mrc20.EventHashLogTransfer:
		filterer, err := mrc20.NewMRC20Filterer(log.Address, nil)
		if err != nil {
			return nil, err
		}

		event, err := filterer.ParseLogTransfer(log)
		if err != nil {
			return nil, err
		}

		transfer.AddressFrom = strings.ToLower(event.From.String())
		transfer.AddressTo = strings.ToLower(event.To.String())
	default:
		return nil, ErrorUnsupportedEvent
	}

	address := strings.ToLower(message.Address)

	if !(transfer.AddressTo == address || transfer.AddressFrom == address) {
		return nil, ErrorUnrelatedEvent
	}

	var err error

	if transfer.SourceData, err = json.Marshal(log); err != nil {
		return nil, err
	}

	return &transfer, nil
}
