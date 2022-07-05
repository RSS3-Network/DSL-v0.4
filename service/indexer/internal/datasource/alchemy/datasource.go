package alchemy

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/naturalselectionlabs/pregod/common/alchemy"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
)

const (
	Source = "alchemy"
)

var (
	ErrorUnsupportedNetwork       = errors.New("unsupported network")
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
	category := []string{
		alchemy.CategoryExternal, alchemy.CategoryERC20, alchemy.CategoryERC721, alchemy.CategoryERC1155,
	}

	// Alchemy doesn't support querying Polygon's internal transactions
	if message.Network == protocol.NetworkEthereum {
		category = append(category, alchemy.CategoryInternal)
	}

	parameter := alchemy.GetAssetTransfersParameter{
		ToAddress: strings.ToLower(message.Address),
		MaxCount:  hexutil.EncodeBig(big.NewInt(alchemy.MaxCount)),
		Category:  category,
	}

	internalTransactionMap := make(map[string]model.Transaction)

	// Get the transactions sent from this address
	internalTransactions, err := d.getAssetTransfers(ctx, message, parameter)
	if err != nil {
		return nil, err
	}

	for _, transaction := range internalTransactions {
		internalTransactionMap[transaction.Hash] = transaction
	}

	// Get the transactions received at this address
	parameter.FromAddress = ""
	parameter.ToBlock = strings.ToLower(message.Address)

	internalTransactions, err = d.getAssetTransfers(ctx, message, parameter)
	if err != nil {
		return nil, err
	}

	// Transaction hash de-duplication
	for _, transaction := range internalTransactions {
		internalTransactionMap[transaction.Hash] = transaction
	}

	ethereumClient, exist := d.ethereumClientMap[message.Network]
	if !exist {
		return nil, ErrorUnsupportedNetwork
	}

	blockMap := make(map[int64]*types.Block)

	// Get transaction data from the chain
	for _, transaction := range internalTransactionMap {
		block, exist := blockMap[transaction.BlockNumber]
		if !exist {
			block, err = ethereumClient.BlockByNumber(ctx, big.NewInt(transaction.BlockNumber))
			if err != nil {
				return nil, err
			}

			blockMap[transaction.BlockNumber] = block
		}

		internalTransaction, _, err := ethereumClient.TransactionByHash(ctx, common.HexToHash(transaction.Hash))
		if err != nil {
			return nil, err
		}

		receipt, err := ethereumClient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
		if err != nil {
			return nil, err
		}

		transactionSourceData, err := ethereum.ConvertTransaction(*internalTransaction, *block)
		if err != nil {
			return nil, err
		}

		sourceData, err := json.Marshal(ethereum.SourceData{
			Transaction: transactionSourceData,
			Receipt:     receipt,
		})
		if err != nil {
			return nil, err
		}

		transaction.SourceData = sourceData
		internalTransactionMap[transaction.Hash] = transaction
	}

	internalTransactions = make([]model.Transaction, 0)

	for _, transaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (d *Datasource) getAssetTransfers(ctx context.Context, message *protocol.Message, parameter alchemy.GetAssetTransfersParameter) ([]model.Transaction, error) {
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
