package alchemy

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/naturalselectionlabs/pregod/common/alchemy"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	lop "github.com/samber/lo/parallel"
)

const (
	Source = "alchemy"

	MaxConcurrency = 200
)

var (
	ErrorUnsupportedNetwork       = errors.New("unsupported network")
	ErrorUnsupportedEvent         = errors.New("unsupported event")
	ErrorUnrelatedEvent           = errors.New("unrelated event")
	ErrorFailedToParseBlockNumber = errors.New("failed to parse block number")
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	ethereumClientMap map[string]*ethclient.Client // QuickNode
	rpcClientMap      map[string]*rpc.Client       // Alchemy
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	ethereumClient, exist := d.ethereumClientMap[message.Network]
	if !exist {
		return nil, ErrorUnsupportedNetwork
	}

	transactionMap, err := d.getAllAssetTransferHashes(ctx, message)
	if err != nil {
		return nil, err
	}

	// Get all block and transaction data
	transactions := make([]*model.Transaction, 0)
	for _, transaction := range transactionMap {
		internalTransaction := transaction

		transactions = append(transactions, &internalTransaction)
	}

	blocks, err := lop.MapWithError(transactions, d.handleBlockFunc(ctx, message, ethereumClient), lop.NewOption().WithConcurrency(MaxConcurrency))
	if err != nil {
		return nil, err
	}

	blockMap := make(map[int64]*types.Block)
	for _, block := range blocks {
		blockMap[block.Number().Int64()] = block
	}

	transactions, err = lop.MapWithError(transactions, d.handleTransactionFunc(ctx, message, ethereumClient, blockMap), lop.NewOption().WithConcurrency(MaxConcurrency))

	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		if transaction != nil {
			internalTransactions = append(internalTransactions, *transaction)
		}
	}

	return internalTransactions, nil
}

func (d *Datasource) handleBlockFunc(ctx context.Context, message *protocol.Message, ethereumClient *ethclient.Client) func(transaction *model.Transaction, i int) (*types.Block, error) {
	return func(transaction *model.Transaction, i int) (*types.Block, error) {
		block, err := ethereumClient.BlockByNumber(ctx, big.NewInt(transaction.BlockNumber))

		if err != nil {
			return nil, err
		}

		return block, nil
	}
}

func (d *Datasource) handleTransactionFunc(ctx context.Context, message *protocol.Message, ethereumClient *ethclient.Client, blockMap map[int64]*types.Block) func(transaction *model.Transaction, i int) (*model.Transaction, error) {
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
			err = ethereum.ErrorUnsupportedTransactionType
		}

		if err != nil {
			return nil, err
		}

		transaction.AddressFrom = strings.ToLower(transactionMessage.From().String())
		transaction.AddressTo = strings.ToLower(transactionMessage.To().String())

		receipt, err := ethereumClient.TransactionReceipt(ctx, internalTransaction.Hash())
		if err != nil {
			return nil, err
		}

		transactionSuccess := receipt.Status == types.ReceiptStatusSuccessful
		transaction.Success = &transactionSuccess

		if transaction.SourceData, err = json.Marshal(&ethereum.SourceData{
			Transaction: internalTransaction,
			Receipt:     receipt,
		}); err != nil {
			return nil, err
		}

		if transaction.Transfers, err = d.handleReceipt(ctx, message, transaction, receipt); err != nil {
			return nil, err
		}

		return transaction, nil
	}
}

func (d *Datasource) handleReceipt(ctx context.Context, message *protocol.Message, transaction *model.Transaction, receipt *types.Receipt) ([]model.Transfer, error) {
	transfers := make([]model.Transfer, 0)

	for _, log := range receipt.Logs {
		transfer, err := d.handleLog(ctx, message, transaction, *log)
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

func (d *Datasource) handleLog(ctx context.Context, message *protocol.Message, transaction *model.Transaction, log types.Log) (*model.Transfer, error) {
	transfer := model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		Index:           int64(log.Index),
		Network:         transaction.Network,
		Source:          d.Name(),
	}

	switch log.Topics[0] {
	case erc20.EventHashTransfer, erc721.EventHashTransfer:
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
	default:
		return nil, ErrorUnsupportedEvent
	}

	if !(transfer.AddressTo == message.Address || transfer.AddressFrom == message.Address) {
		return nil, ErrorUnrelatedEvent
	}

	var err error

	if transfer.SourceData, err = json.Marshal(log); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (d *Datasource) getAllAssetTransferHashes(ctx context.Context, message *protocol.Message) (map[string]model.Transaction, error) {
	category := []string{
		alchemy.CategoryExternal, alchemy.CategoryERC20, alchemy.CategoryERC721, alchemy.CategoryERC1155,
	}

	// Alchemy doesn't support querying Polygon's internal transactions
	if message.Network == protocol.NetworkEthereum {
		category = append(category, alchemy.CategoryInternal)
	}

	parameter := alchemy.GetAssetTransfersParameter{
		FromAddress: strings.ToLower(message.Address),
		MaxCount:    hexutil.EncodeBig(big.NewInt(alchemy.MaxCount)),
		Category:    category,
	}

	internalTransactionMap := make(map[string]model.Transaction)

	// Get the transactions sent from this address
	internalTransactions, err := d.getAssetTransactionHashes(ctx, message, parameter)
	if err != nil {
		return nil, err
	}

	for _, transaction := range internalTransactions {
		internalTransactionMap[transaction.Hash] = transaction
	}

	// Get the transactions received at this address
	parameter.FromAddress = ""
	parameter.ToAddress = strings.ToLower(message.Address)

	internalTransactions, err = d.getAssetTransactionHashes(ctx, message, parameter)
	if err != nil {
		return nil, err
	}

	// Transaction hash de-duplication
	for _, transaction := range internalTransactions {
		internalTransactionMap[transaction.Hash] = transaction
	}

	return internalTransactionMap, nil
}

func (d *Datasource) getAssetTransactionHashes(ctx context.Context, message *protocol.Message, parameter alchemy.GetAssetTransfersParameter) ([]model.Transaction, error) {
	rpcClient, exist := d.rpcClientMap[message.Network]
	if !exist {
		return nil, ErrorUnsupportedNetwork
	}

	internalTransactions := make([]model.Transaction, 0)

	for {
		result, err := alchemy.GetAssetTransfers(context.Background(), rpcClient, parameter)
		if err != nil {
			return nil, err
		}

		if len(result.Transfers) == 0 {
			break
		}

		for _, transfer := range result.Transfers {
			blockNumber := big.NewInt(0)

			if _, ok := blockNumber.SetString(transfer.BlockNum, 0); !ok {
				return nil, ErrorFailedToParseBlockNumber
			}

			internalTransactions = append(internalTransactions, model.Transaction{
				BlockNumber: blockNumber.Int64(),
				Hash:        transfer.Hash,
				Network:     message.Network,
				Source:      d.Name(),
				Transfers:   make([]model.Transfer, 0),
			})
		}

		if result.PageKey == "" {
			break
		}

		parameter.PageKey = result.PageKey
	}

	return internalTransactions, nil
}

func New(config *configx.RPC) (datasource.Datasource, error) {
	internalDatasource := Datasource{
		ethereumClientMap: map[string]*ethclient.Client{},
		rpcClientMap:      map[string]*rpc.Client{},
	}

	var err error

	if internalDatasource.ethereumClientMap[protocol.NetworkEthereum], err = ethclient.Dial(config.General.Ethereum.HTTP); err != nil {
		return nil, err
	}

	if internalDatasource.ethereumClientMap[protocol.NetworkPolygon], err = ethclient.Dial(config.General.Polygon.HTTP); err != nil {
		return nil, err
	}

	if internalDatasource.rpcClientMap[protocol.NetworkEthereum], err = rpc.Dial(config.Alchemy.Ethereum.HTTP); err != nil {
		return nil, err
	}
	if internalDatasource.rpcClientMap[protocol.NetworkPolygon], err = rpc.Dial(config.Alchemy.Polygon.HTTP); err != nil {
		return nil, err
	}

	return &internalDatasource, nil
}
