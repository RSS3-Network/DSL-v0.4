package crossbell

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/kurora/constant"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	contract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/crossbell"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var (
	_                 crawler.Crawler = (*service)(nil)
	crossbellCacheKey                 = "crawler:crossbell"
)

type service struct {
	config          *config.Config
	kuroraClient    *kurora.Client
	crossbellClient *crossbell.Client
}

func New(conf *config.Config) crawler.Crawler {
	return &service{
		config: conf,
	}
}

func (s *service) Name() string {
	return protocol.PlatformCrossbell
}

func (s *service) Network() string {
	return protocol.NetworkCrossbell
}

func (s *service) Run() error {
	var err error
	ctx := context.Background()

	s.kuroraClient, err = kurora.Dial(ctx, s.config.RPC.PregodETL.KuroraV2.HTTP, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error("crossbell: kurora.Dial error", zap.Error(err), zap.String("endpoint", s.config.RPC.PregodETL.KuroraV2.HTTP))

		return err
	}

	s.crossbellClient, err = crossbell.New()
	if err != nil {
		loggerx.Global().Error("crossbell: crossbell.New error", zap.Error(err))

		return err
	}

	for {
		transactions, err := s.GetKuroraLogs(ctx)
		if err != nil {
			loggerx.Global().Error("crossbell: GetKuroraLogs error", zap.Error(err), zap.String("endpoint", s.config.RPC.PregodETL.KuroraV2.HTTP))
			return err
		}

		if len(transactions) == 0 {
			time.Sleep(10 * time.Minute)
			continue
		}

		// deduplicate data
		transactions, err = database.DeduplicateTransactions(ctx, transactions)
		if err != nil || len(transactions) == 0 {
			continue
		}

		// build transaction
		message := &protocol.Message{
			Network: s.Network(),
		}
		if transactions, err = ethereum.BuildTransactions(ctx, message, transactions); err != nil {
			loggerx.Global().Error("failed to build transactions", zap.Error(err))

			continue
		}

		// crossbell worker
		internalTransactions := s.getInternalTransaction(ctx, transactions)

		// insert db
		err = database.UpsertTransactions(ctx, internalTransactions)
		if err != nil {
			continue
		}
	}

	return nil
}

func (s *service) GetKuroraLogs(ctx context.Context) ([]*model.Transaction, error) {
	tracer := otel.Tracer("crossbell")
	_, trace := tracer.Start(ctx, "crossbell:GetKuroraLogs")
	var internalTransactions []*model.Transaction
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

	cursor, _ := cache.Global().Get(ctx, crossbellCacheKey).Result()
	query := kurora.EthereumLogQuery{
		AddressList: []common.Address{
			contract.AddressCharacter,
			contract.AddressLinkList,
		},
		Cursor: &cursor,
	}

	if len(cursor) > 0 {
		query.Cursor = &cursor
	}

	result, err := s.kuroraClient.FetchEthereumLogs(ctx, constant.NetworkCrossbell, query)
	if err != nil {
		loggerx.Global().Error("kuroraClient FetchEthereumLogs error", zap.Error(err))
		return nil, err
	}

	transactionMap := make(map[string]*model.Transaction)
	for _, log := range result {
		if _, exists := transactionMap[log.TransactionHash.String()]; !exists {
			transactionMap[log.TransactionHash.String()] = &model.Transaction{
				BlockNumber: log.BlockNumber.BigInt().Int64(),
				Hash:        log.TransactionHash.String(),
				Index:       int64(log.TransactionIndex),
				Network:     s.Network(),
				Source:      protocol.SourcePregodETL,
				Platform:    s.Name(),
				Transfers:   make([]model.Transfer, 0),
			}
		}

		transfers := transactionMap[log.TransactionHash.String()].Transfers
		transfers = append(transfers, model.Transfer{
			Timestamp:   log.Timestamp,
			BlockNumber: log.BlockNumber.BigInt(),
			Index:       int64(log.Index),
			Metadata:    metadata.Default,
			Network:     s.Network(),
			Platform:    s.Name(),
			Source:      protocol.SourcePregodETL,
		})
	}

	for _, tx := range transactionMap {
		internalTransactions = append(internalTransactions, tx)
	}

	// set cache
	last, err := lo.Last(result)
	if err == nil {
		cursor := kurora.LogCursor(last.TransactionHash, last.Index)
		cache.Global().Set(ctx, crossbellCacheKey, cursor, 7*24*time.Hour)
	}

	return internalTransactions, nil
}

func (s *service) getInternalTransaction(ctx context.Context, transactions []*model.Transaction) []*model.Transaction {
	var mu sync.Mutex
	var internalTransactions []*model.Transaction
	opt := lop.NewOption().WithConcurrency(10)
	lop.ForEach(transactions, func(transaction *model.Transaction, i int) {
		addressTo := common.HexToAddress(transaction.AddressTo)
		if addressTo != contract.AddressCharacter && addressTo != contract.AddressLinkList {
			return
		}

		transferMap := make(map[int64]model.Transfer)
		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)
		transaction.Transfers = append(transaction.Transfers, transferMap[protocol.IndexVirtual])

		var sourceData ethereum.SourceData
		if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
			zap.L().Error("failed to unmarshal source data", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

			return
		}

		message := protocol.Message{
			Network: s.Network(),
		}

		internalTransfers, err := s.crossbellClient.HandleReceipt(ctx, &message, transaction, sourceData.Receipt, transferMap)
		if err != nil {
			zap.L().Error("failed to handle receipt", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

			return
		}

		transaction.Transfers = append(transaction.Transfers, internalTransfers...)

		for _, transfer := range transaction.Transfers {
			transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
		}

		mu.Lock()
		internalTransactions = append(internalTransactions, transaction)
		mu.Unlock()
	}, opt)

	return internalTransactions
}
