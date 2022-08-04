package blockscout

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"go.opentelemetry.io/otel"
)

const (
	Source = "blockscout"
)

var _ datasource.Datasource = (*Datasource)(nil)

type Datasource struct {
	ethereumClientMap   map[string]*ethclient.Client
	blockscoutClientMap map[string]*blockscout.Client
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkEthereumClassic, protocol.NetworkXDAI,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("datasource_blockscout")
	_, trace := tracer.Start(ctx, "datasource_blockscout:Handle")

	defer opentelemetry.Log(trace, message, transactions, err)

	ethereumClient, exist := d.ethereumClientMap[message.Network]
	if !exist {
		return nil, errors.New("no ethereum client for network")
	}

	transactionMap := make(map[string]*model.Transaction)

	nativeTransactions, err := d.handleTransactions(ctx, message)
	if err != nil {
		return nil, err
	}

	for _, transaction := range nativeTransactions {
		internalTransaction := transaction

		transactionMap[transaction.Hash] = &internalTransaction
	}

	tokenTransfers, err := d.handleTokenTransfers(ctx, message)
	if err != nil {
		return nil, err
	}

	for _, transfer := range tokenTransfers {
		transactionMap[transfer.TransactionHash] = &model.Transaction{
			BlockNumber: transfer.BlockNumber.Int64(),
			Hash:        transfer.TransactionHash,
			Network:     message.Network,
			Source:      d.Name(),
		}
	}

	unindexedTransactions := make([]*model.Transaction, 0)

	for _, transaction := range transactionMap {
		if transaction.AddressFrom != "" && !strings.EqualFold(transaction.AddressFrom, message.Address) && !allowlist.Allow(transaction.AddressFrom) {
			continue
		}

		unindexedTransactions = append(unindexedTransactions, transaction)
	}

	indexedTransactions, err := ethereum.BuildTransactions(ctx, message, unindexedTransactions, ethereumClient)
	if err != nil {
		return nil, err
	}

	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range indexedTransactions {
		if transaction.Timestamp.Before(message.Timestamp) {
			continue
		}

		internalTransactions = append(internalTransactions, *transaction)
	}

	return internalTransactions, nil
}

func (d *Datasource) handleTransactions(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("datasource_blockscout")
	_, span := tracer.Start(ctx, "datasource_blockscout:handleTransactions")

	defer opentelemetry.Log(span, message, transactions, err)

	transactions = make([]model.Transaction, 0)

	// Use a different Client for different networks
	blockscoutClient := d.blockscoutClientMap[message.Network]

	internalTransactions, _, err := blockscoutClient.GetTransactionList(ctx, common.HexToAddress(message.Address), &blockscout.GetTransactionListOption{})
	if err != nil {
		return nil, err
	}

	for _, internalTransaction := range internalTransactions {
		sourceData, err := json.Marshal(internalTransaction)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, model.Transaction{
			Hash:        internalTransaction.Hash,
			BlockNumber: internalTransaction.BlockNumber.IntPart(),
			AddressFrom: strings.ToLower(internalTransaction.From),
			AddressTo:   strings.ToLower(internalTransaction.To),
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
		})
	}

	return transactions, err
}

func (d *Datasource) handleTokenTransfers(ctx context.Context, message *protocol.Message) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("datasource_blockscout")
	_, span := tracer.Start(ctx, "datasource_blockscout:handleTokenTransfers")

	defer opentelemetry.Log(span, message, transfers, err)

	// Use a different Client for different networks
	blockscoutClient := d.blockscoutClientMap[message.Network]

	internalTokenTransfers, _, err := blockscoutClient.GetTokenTransactionList(ctx, common.HexToAddress(message.Address), &blockscout.GetTokenTransactionListOption{})
	if err != nil {
		return nil, err
	}

	for _, internalTokenTransfer := range internalTokenTransfers {
		sourceData, err := json.Marshal(internalTokenTransfer)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, model.Transfer{
			TransactionHash: internalTokenTransfer.Hash,
			BlockNumber:     internalTokenTransfer.BlockNumber.BigInt(),
			Network:         message.Network,
			Source:          d.Name(),
			SourceData:      sourceData,
		})
	}

	return transfers, nil
}

func New(config *configx.RPC) (datasource.Datasource, error) {
	ethereumClientEthereum, err := ethclient.Dial(config.General.Ethereum.HTTP)
	if err != nil {
		return nil, err
	}

	ethereumClientPolygon, err := ethclient.Dial(config.General.Polygon.HTTP)
	if err != nil {
		return nil, err
	}

	ethereumClientBinanceSmartChain, err := ethclient.Dial(config.General.BinanceSmartChain.HTTP)
	if err != nil {
		return nil, err
	}

	ethereumClientXDAI, err := ethclient.Dial(config.General.XDAI.HTTP)
	if err != nil {
		return nil, err
	}

	ethereumClientCrossbell, err := ethclient.Dial(config.General.Crossbell.HTTP)
	if err != nil {
		return nil, err
	}

	return &Datasource{
		ethereumClientMap: map[string]*ethclient.Client{
			protocol.NetworkEthereum:          ethereumClientEthereum,
			protocol.NetworkPolygon:           ethereumClientPolygon,
			protocol.NetworkBinanceSmartChain: ethereumClientBinanceSmartChain,
			protocol.NetworkCrossbell:         ethereumClientCrossbell,
			protocol.NetworkXDAI:              ethereumClientXDAI,
		},
		blockscoutClientMap: map[string]*blockscout.Client{
			protocol.NetworkEthereum:        blockscout.New(blockscout.EndpointDefault, blockscout.NetworkEthereum),
			protocol.NetworkEthereumClassic: blockscout.New(blockscout.EndpointDefault, blockscout.NetworkEthereumClassic),
			protocol.NetworkXDAI:            blockscout.New(blockscout.EndpointDefault, blockscout.NetworkXDAI),
			protocol.NetworkCrossbell:       blockscout.New(blockscout.EndpointCrossbell, blockscout.NetworkCrossbell),
		},
	}, nil
}
