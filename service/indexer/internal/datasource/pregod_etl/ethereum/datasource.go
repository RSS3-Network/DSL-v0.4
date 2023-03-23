package ethereum

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/kurora/constant"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	eth_etl "github.com/naturalselectionlabs/pregod/common/datasource/pregod_etl/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	kuroraClient *kurora.Client
}

func (d *Datasource) Name() string {
	return protocol.SourcePregodETL
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

		if internalTransaction.BlockNumber < message.BlockNumber && !strings.EqualFold(internalTransaction.AddressFrom, message.Address) &&
			!allowlist.AllowList.Contains(internalTransaction.AddressFrom) && !allowlist.AllowList.Contains(internalTransaction.AddressTo) {
			continue
		}

		transactions = append(transactions, &internalTransaction)
	}

	// if transactions, err = ethereum.BuildTransactions(ctx, message, transactions); err != nil {
	//	 loggerx.Global().Error("failed to build transactions", zap.Error(err))
	//
	//	 return nil, err
	// }

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

func (d *Datasource) getAssestTransactionAndLogs(ctx context.Context, message *protocol.Message, parameter eth_etl.GetAssetTransfersParameter) (*eth_etl.GetAssetTransfersResult, error) {
	result := new(eth_etl.GetAssetTransfersResult)
	var mu sync.Mutex
	var eg errgroup.Group
	eg.Go(func() error {
		txns, err := d.kuroraClient.FetchEthereumTransactions(ctx, constant.NetworkEthereum, kurora.EthereumTransactionQuery{
			From:            lo.ToPtr[common.Address](common.BytesToAddress([]byte(parameter.FromAddress))),
			BlockNumberFrom: lo.ToPtr[decimal.Decimal](decimal.New(int64(parameter.FromBlock), 0)),
			Limit:           lo.ToPtr[int](eth_etl.DefaultTransactionLimit),
		})
		if err != nil {
			loggerx.Global().Error("failed to fetch ethereum transactions", zap.Error(err))
			return err
		}
		mu.Lock()
		for i := range txns {
			result.Transfers = append(result.Transfers, eth_etl.Transfer{
				BlockNum: txns[i].BlockNumber.IntPart(),
				Hash:     txns[i].Hash,
				From:     txns[i].From,
				To:       (*txns[i].To)[:],
				Value:    float64(txns[i].Value.IntPart()),
			})
		}
		mu.Unlock()

		return nil
	})

	eg.Go(func() error {
		ethereumLogs, err := d.kuroraClient.FetchEthereumLogs(ctx, constant.NetworkEthereum, kurora.EthereumLogQuery{
			BlockNumberFrom: lo.ToPtr[decimal.Decimal](decimal.New(int64(parameter.FromBlock), 0)),
			TopicFirst:      lo.ToPtr[common.Hash](erc20.EventHashTransfer),
			TopicThird:      lo.ToPtr[common.Hash](common.BytesToHash([]byte(parameter.FromAddress))),
			Limit:           lo.ToPtr[int](eth_etl.DefaultTransferLimit),
		})
		if err != nil {
			loggerx.Global().Error("failed to fetch ethereum logs", zap.Error(err))
			return err
		}
		mu.Lock()
		for i := range ethereumLogs {
			value, _ := hexutil.DecodeUint64(ethereumLogs[i].Data.String())
			result.Transfers = append(result.Transfers, eth_etl.Transfer{
				BlockNum: ethereumLogs[i].BlockNumber.IntPart(),
				Hash:     ethereumLogs[i].TransactionHash,
				From:     common.BytesToAddress(ethereumLogs[i].TopicSecond.Bytes()),
				To:       ethereumLogs[i].TopicThird.Bytes(),
				Value:    float64(value),
			})
		}
		mu.Unlock()
		return nil
	})

	sort.Slice(result.Transfers, func(i, j int) bool {
		return result.Transfers[i].BlockNum < result.Transfers[j].BlockNum
	})

	return result, eg.Wait()
}

func (d *Datasource) getAssetTransactionHashes(ctx context.Context, message *protocol.Message, parameter eth_etl.GetAssetTransfersParameter) ([]model.Transaction, error) {
	tracer := otel.Tracer("datasource_etl")
	_, trace := tracer.Start(ctx, "datasource_etl:Handle")
	internalTransactions := make([]model.Transaction, 0)

	result, err := d.getAssestTransactionAndLogs(ctx, message, parameter)
	// Deprecated result, err := eth_etl.GetAssetTransfers(context.Background(), parameter)
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
			AddressFrom: strings.ToLower(transfer.From.Hex()),
		}

		if len(transfer.To) > 0 {
			transaction.AddressTo = strings.ToLower(common.BytesToAddress(transfer.To).Hex())
		}

		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func New(ctx context.Context, endpoint string) (datasource.Datasource, error) {
	internalDatasource := &Datasource{}
	var err error
	internalDatasource.kuroraClient, err = kurora.Dial(ctx, endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error("failed to connect to kurora", zap.Error(err), zap.String("endpoint", endpoint))
		return nil, err
	}
	return internalDatasource, nil
}
