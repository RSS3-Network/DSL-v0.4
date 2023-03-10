package sound

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	soundContract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/sound"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/sound"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type service struct {
	kuroraClient *kurora.Client
}

var (
	_ crawler.Crawler = (*service)(nil)

	soundLogsCacheKey = "crawler:sound"
)

func New(config *config.Config) (crawler.Crawler, error) {
	kuroraClient, err := kurora.Dial(context.Background(), config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		logrus.Errorf("sound: kuroraClient NewClient error, %v", err)
		return nil, err
	}

	sound.APIKey = config.Sound.APIKey

	return &service{
		kuroraClient: kuroraClient,
	}, nil
}

func (s *service) Name() string {
	return protocol.PlatformSound
}

func (s *service) Run() error {
	ctx := context.Background()

	for {
		transactions, cacheInfo, err := s.getLogs(ctx)
		if err != nil {
			loggerx.Global().Error("sound: getLogs error", zap.Error(err))

			return err
		}

		if len(transactions) == 0 {
			time.Sleep(1 * time.Minute)
			continue
		}

		message := &protocol.Message{Network: protocol.NetworkEthereum}
		if transactions, err = ethereum.BuildTransactions(ctx, message, transactions); err != nil {
			loggerx.Global().Error("sound: build transaction error", zap.Error(err))

			return err
		}

		// handle
		result := s.handle(ctx, transactions)

		// insert db
		err = database.UpsertTransactions(ctx, result, true)
		if err != nil {
			continue
		}

		// set cache
		cache.Global().Set(ctx, soundLogsCacheKey, cacheInfo, 7*24*time.Hour)
	}
}

func (s *service) getLogs(ctx context.Context) ([]*model.Transaction, string, error) {
	tracer := otel.Tracer("sound")
	_, trace := tracer.Start(ctx, "sound:getLogs")
	var internalTransactions []*model.Transaction
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

	query := kurora.DatasetSoundEventQuery{
		TopicFirstList: []common.Hash{
			soundContract.EventHashEditionCreatedV1,
			soundContract.EventHashEditionCreatedV3,
			soundContract.EventHashEditionPurchasedV3,
			soundContract.EventHashEditionPurchasedV5,
		},
		Limit: lo.ToPtr(100),
	}

	cacheInfo, _ := cache.Global().Get(ctx, soundLogsCacheKey).Result()
	if cacheList := strings.Split(cacheInfo, ":"); len(cacheList) == 2 {
		blockNumber, _ := decimal.NewFromString(cacheList[0])
		query.BlockNumberFrom = &blockNumber
		query.Cursor = &cacheList[1]
	}

	loggerx.Global().Info("FetchDatasetLensEvents request", zap.Any("query", query))

	result, err := s.kuroraClient.FetchDatasetSoundEvents(ctx, query)
	if err != nil {
		loggerx.Global().Error("FetchDatasetLensEvents error", zap.Error(err), zap.Any("query", query))

		return nil, "", err
	}

	transactionMap := make(map[string]*model.Transaction)
	for _, event := range result {
		transaction := &model.Transaction{
			Hash:        event.TransactionHash.String(),
			BlockNumber: event.BlockNumber.BigInt().Int64(),
			Index:       int64(event.TransactionIndex),
			Network:     protocol.NetworkEthereum,
			Platform:    protocol.PlatformSound,
			Transfers:   make([]model.Transfer, 0),
			Source:      protocol.SourceKurora,
		}
		transactionMap[transaction.Hash] = transaction
	}

	for _, transaction := range transactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	last, err := lo.Last(result)
	if err != nil {
		return nil, "", err
	}

	cursor := kurora.LogCursor(last.TransactionHash, last.Index)
	cacheInfo = fmt.Sprintf("%v:%v", last.BlockNumber, cursor)

	return internalTransactions, cacheInfo, nil
}

func (s *service) handle(ctx context.Context, transactions []*model.Transaction) []*model.Transaction {
	var mu sync.Mutex
	var result []*model.Transaction
	opt := lop.NewOption().WithConcurrency(10)
	lop.ForEach(transactions, func(transaction *model.Transaction, i int) {
		transferMap := make(map[int64]model.Transfer)
		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)

		// get receipt
		transfers, err := sound.HandleReceipt(ctx, transaction)
		if err != nil {
			loggerx.Global().Error("sound.HandleReceipt error", zap.Error(err), zap.Any("transaction", transaction))

			return
		}

		transaction.Transfers = append(transaction.Transfers, transfers...)

		for _, transfer := range transaction.Transfers {
			if transaction.Tag == "" {
				transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
			}
		}

		mu.Lock()
		result = append(result, transaction)
		mu.Unlock()
	}, opt)

	return result
}
