package lens

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/shopspring/decimal"

	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"

	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	lens_comm "github.com/naturalselectionlabs/pregod/common/worker/lens"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

var (
	_ crawler.Crawler = (*service)(nil)

	lensLogsCacheKey = "crawler:lens"
)

type service struct {
	config           *config.Config
	kuroraClient     *kurora.Client
	commWorkerClient *lens_comm.Client
}

func New(config *config.Config) crawler.Crawler {
	crawler := &service{
		config:           config,
		commWorkerClient: lens_comm.New(),
	}

	return crawler
}

func (s *service) Name() string {
	return protocol.PlatformLens
}

func (s *service) Run() error {
	ctx := context.Background()

	var err error
	s.kuroraClient, err = kurora.Dial(context.Background(), s.config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		logrus.Errorf("[lens] kuroraClient NewClient error, %v", err)
		return err
	}

	for {
		// get lens logs
		transactions, err := s.getLensLogs(ctx)
		if err != nil {
			loggerx.Global().Error("failed to query lens", zap.Error(err))

			time.Sleep(10 * time.Minute)
			continue
		}

		if len(transactions) == 0 {
			time.Sleep(10 * time.Minute)

			continue
		}

		// deduplicate data
		// transactions, err = database.DeduplicateTransactions(ctx, transactions)
		// if err != nil || len(transactions) == 0 {
		// 	continue
		// }

		// build transaction
		message := &protocol.Message{
			Network: protocol.NetworkPolygon,
		}
		if transactions, err = ethereum.BuildTransactions(ctx, message, transactions); err != nil {
			logrus.Error("failed to build transactions, ", err)

			continue
		}

		// lens worker
		internalTransactions := s.getInternalTransaction(ctx, transactions)

		// insert db
		err = database.UpsertTransactions(ctx, internalTransactions, false)
		if err != nil {
			continue
		}
	}
}

func (s *service) getLensLogs(ctx context.Context) ([]*model.Transaction, error) {
	tracer := otel.Tracer("lens")
	_, trace := tracer.Start(ctx, "len:GetLensLogs")
	var internalTransactions []*model.Transaction
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

	query := kurora.DatasetLensEventQuery{
		Address: &lens.HubProxyContractAddress,
	}

	cacheInfo, _ := cache.Global().Get(ctx, lensLogsCacheKey).Result()
	if cacheList := strings.Split(cacheInfo, ":"); len(cacheList) == 2 {
		blockNumber, _ := decimal.NewFromString(cacheList[0])
		query.BlockNumberFrom = &blockNumber
		query.Cursor = &cacheList[1]
	}

	loggerx.Global().Info("FetchDatasetLensEvents request", zap.Any("query", query))

	// get logs from pregod_etl
	result, err := s.kuroraClient.FetchDatasetLensEvents(ctx, query)
	if err != nil {
		loggerx.Global().Error("FetchDatasetLensEvents error", zap.Error(err), zap.Any("query", query))
		return nil, err
	}

	transactionMap := make(map[string]*model.Transaction)
	for _, transfer := range result {
		transaction := &model.Transaction{
			BlockNumber: transfer.BlockNumber.BigInt().Int64(),
			Hash:        transfer.TransactionHash.String(),
			Index:       int64(transfer.TransactionIndex),
			Network:     protocol.NetworkPolygon,
			Platform:    protocol.PlatformLens,
			Transfers:   make([]model.Transfer, 0),
			Source:      protocol.SourceKurora,
		}

		transactionMap[transaction.Hash] = transaction
	}

	for _, transaction := range transactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	// set cache
	if len(result) == 0 {
		return internalTransactions, nil
	}

	cursor := kurora.LogCursor(result[len(result)-1].TransactionHash, result[len(result)-1].Index)
	cacheInfo = fmt.Sprintf("%v:%v", result[len(result)-1].BlockNumber, cursor)
	cache.Global().Set(ctx, lensLogsCacheKey, cacheInfo, 7*24*time.Hour)

	return internalTransactions, nil
}

func (s *service) getInternalTransaction(ctx context.Context, transactions []*model.Transaction) []*model.Transaction {
	var mu sync.Mutex
	var internalTransactions []*model.Transaction
	opt := lop.NewOption().WithConcurrency(10)
	lop.ForEach(transactions, func(transaction *model.Transaction, i int) {
		transferMap := make(map[int64]model.Transfer)
		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)
		transaction.Transfers = append(transaction.Transfers, transferMap[protocol.IndexVirtual])

		// get receipt
		internalTransfers, err := s.commWorkerClient.HandleReceipt(ctx, transaction)
		if err != nil {
			logrus.Errorf("[lens worker] handleReceipt error, %v", err)

			return
		}

		transaction.Transfers = append(transaction.Transfers, internalTransfers...)

		for _, transfer := range transaction.Transfers {
			if transaction.Tag == "" {
				transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
			}
		}

		mu.Lock()
		internalTransactions = append(internalTransactions, transaction)
		mu.Unlock()
	}, opt)

	return internalTransactions
}
