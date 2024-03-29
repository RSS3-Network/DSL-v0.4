package alchemy

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/alchemy"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var (
	ErrorUnsupportedNetwork       = errors.New("unsupported network")
	ErrorFailedToParseBlockNumber = errors.New("failed to parse block number")
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	rpcClientMap map[string]*alchemy.Client // Alchemy
}

func (d *Datasource) Name() string {
	return protocol.SourceAlchemy
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkPolygon,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	tracer := otel.Tracer("datasource_alchemy")
	ctx, trace := tracer.Start(ctx, "datasource_alchemy:Handle")

	transactions := make([]*model.Transaction, 0)
	var err error
	defer func() { opentelemetry.Log(trace, message, len(transactions), err) }()

	transactionMap, err := d.getAllAssetTransferHashes(ctx, message)
	if err != nil {
		loggerx.Global().Error("failed to get all asset transfer hashes", zap.Error(err))

		return nil, err
	}

	// Get all block and transaction data
	for _, transaction := range transactionMap {
		internalTransaction := transaction

		if internalTransaction.BlockNumber < message.BlockNumber {
			continue
		}

		if internalTransaction.AddressFrom != "" && !strings.EqualFold(internalTransaction.AddressFrom, message.Address) && !allowlist.AllowList.Contains(internalTransaction.AddressFrom) {
			continue
		}

		transactions = append(transactions, &internalTransaction)
	}

	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		if transaction != nil {
			internalTransactions = append(internalTransactions, *transaction)
		}
	}

	return internalTransactions, nil
}

func (d *Datasource) getAllAssetTransferHashes(ctx context.Context, message *protocol.Message) (map[string]model.Transaction, error) {
	tracer := otel.Tracer("datasource_alchemy")
	ctx, trace := tracer.Start(ctx, "datasource_alchemy:getAllAssetTransferHashes")
	internalTransactionMap := make(map[string]model.Transaction)
	var err error
	defer func() { opentelemetry.Log(trace, message, internalTransactionMap, err) }()

	category := []string{
		alchemy.CategoryExternal, alchemy.CategoryERC20, alchemy.CategoryERC721, alchemy.CategoryERC1155,
	}

	parameter := alchemy.GetAssetTransfersParameter{
		FromAddress: strings.ToLower(message.Address),
		MaxCount:    hexutil.EncodeBig(big.NewInt(alchemy.MaxCount)),
		Category:    category,
		FromBlock:   hexutil.EncodeUint64(uint64(message.BlockNumber)),
		Order:       "desc",
	}

	// Get the transactions sent from this address
	internalTransactions, err := d.getAssetTransactionHashes(ctx, message, parameter)
	if err != nil {
		return nil, err
	}

	for _, transaction := range internalTransactions {
		internalTransactionMap[transaction.Hash] = transaction
	}

	// Find POAP transactions from XDAI
	if message.Network == protocol.NetworkXDAI {
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
	}

	return internalTransactionMap, nil
}

func (d *Datasource) getAssetTransactionHashes(ctx context.Context, message *protocol.Message, parameter alchemy.GetAssetTransfersParameter) ([]model.Transaction, error) {
	tracer := otel.Tracer("datasource_alchemy")
	_, trace := tracer.Start(ctx, "datasource_alchemy:Handle")
	internalTransactions := make([]model.Transaction, 0)

	var err error
	defer func() { opentelemetry.Log(trace, message, len(internalTransactions), err) }()

	alchemyClient, exist := d.rpcClientMap[message.Network]
	if !exist {
		return nil, ErrorUnsupportedNetwork
	}

	for {
		var result *alchemy.GetAssetTransfersResult

		result, err = alchemyClient.GetAssetTransfers(context.Background(), parameter)
		if err != nil {
			loggerx.Global().Error("failed to get asset transfers", zap.Error(err))

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

			transaction := model.Transaction{
				BlockNumber: blockNumber.Int64(),
				Hash:        transfer.Hash,
				Network:     message.Network,
				Source:      d.Name(),
				Transfers:   make([]model.Transfer, 0),
			}

			if transfer.Category == alchemy.CategoryExternal {
				transaction.AddressFrom = transfer.From
				transaction.AddressTo = transfer.To
			}

			internalTransactions = append(internalTransactions, transaction)
		}

		if result.PageKey == "" {
			break
		}

		if len(internalTransactions) > datasource.DatasourceLimit {
			break
		}

		parameter.PageKey = result.PageKey
	}

	return internalTransactions, nil
}

func New(config *configx.RPC) (datasource.Datasource, error) {
	internalDatasource := Datasource{
		rpcClientMap: map[string]*alchemy.Client{},
	}

	var err error

	if internalDatasource.rpcClientMap[protocol.NetworkEthereum], err = alchemy.NewClient(protocol.NetworkEthereum, config.Alchemy.Ethereum); err != nil {
		return nil, err
	}
	if internalDatasource.rpcClientMap[protocol.NetworkPolygon], err = alchemy.NewClient(protocol.NetworkPolygon, config.Alchemy.Polygon); err != nil {
		return nil, err
	}

	return &internalDatasource, nil
}
