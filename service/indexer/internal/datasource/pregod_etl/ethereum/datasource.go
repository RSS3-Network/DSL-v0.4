package ethereum

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	eth_etl "github.com/naturalselectionlabs/pregod/common/datasource/pregod_etl/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const (
	Source = "ethereum-etl"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct{}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	tracer := otel.Tracer("datasource_etl")
	ctx, trace := tracer.Start(ctx, "datasource_etl:Handle")

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

		transactions = append(transactions, &internalTransaction)
	}

	if transactions, err = ethereum.BuildTransactions(ctx, message, transactions); err != nil {
		loggerx.Global().Error("failed to build transactions", zap.Error(err))

		return nil, err
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
	tracer := otel.Tracer("datasource_etl")
	ctx, trace := tracer.Start(ctx, "datasource_etl:getAllAssetTransferHashes")
	internalTransactionMap := make(map[string]model.Transaction)
	var err error
	defer func() { opentelemetry.Log(trace, message, internalTransactionMap, err) }()

	parameter := eth_etl.GetAssetTransfersParameter{
		FromAddress: strings.ToLower(message.Address),
		FromBlock:   message.BlockNumber,
	}

	// Get the transactions sent from this address
	internalTransactions, err := d.getAssetTransactionHashes(ctx, message, parameter)
	if err != nil {
		return nil, err
	}

	for _, transaction := range internalTransactions {
		internalTransactionMap[transaction.Hash] = transaction
	}

	return internalTransactionMap, nil
}

func (d *Datasource) getAssetTransactionHashes(ctx context.Context, message *protocol.Message, parameter eth_etl.GetAssetTransfersParameter) ([]model.Transaction, error) {
	tracer := otel.Tracer("datasource_etl")
	_, trace := tracer.Start(ctx, "datasource_etl:Handle")
	internalTransactions := make([]model.Transaction, 0)

	result, err := eth_etl.GetAssetTransfers(context.Background(), parameter)
	defer func() { opentelemetry.Log(trace, message, len(internalTransactions), err) }()

	if err != nil {
		loggerx.Global().Error("failed to get ethereum data from ETL", zap.Error(err))

		return nil, err
	}

	if len(result.Transfers) == 0 {
		return internalTransactions, nil
	}

	for _, transfer := range result.Transfers {

		transaction := model.Transaction{
			BlockNumber: transfer.BlockNum,
			Hash:        transfer.Hash.Hex(),
			Network:     message.Network,
			Source:      d.Name(),
			Transfers:   make([]model.Transfer, 0),
		}

		if strings.EqualFold(transfer.Category, eth_etl.CategoryExternal) {
			transaction.AddressFrom = strings.ToLower(transfer.From.Hex())
			if len(transfer.To) > 0 {
				transaction.AddressTo = strings.ToLower(common.BytesToAddress(transfer.To).Hex())
			}
		}

		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func New() datasource.Datasource {
	return &Datasource{}
}
