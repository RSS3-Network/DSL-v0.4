package lens

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/datasource/pregod_etl"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/lens_comm"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/scylladb/go-set/strset"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

var (
	_ crawler.Crawler = (*service)(nil)

	lensLogsCacheKey = "crawler:lens:%v" // eventHash -> blocknumber:cursor
)

type service struct {
	config           *config.Config
	ethClient        *ethclient.Client
	etlClient        *pregod_etl.Client
	commWorkerClient *lens_comm.Client
}

func New(config *config.Config) crawler.Crawler {

}

func (s *service) Run() error {
	var ctx = context.Background()
	var wg sync.WaitGroup
	for eventHash, contractAddress := range lens.SupportLensEvents {
		wg.Add(1)
		go func(eventHash common.Hash, contractAddress common.Address) {
			defer wg.Done()

			for {
				// get lens logs
				transactions, err := s.getLensLogs(ctx, eventHash, contractAddress)
				if err != nil {
					loggerx.Global().Error("[lens] Run: GetLensLogs error", zap.Error(err))

					return
				}

				if len(transactions) == 0 {
					time.Sleep(10 * time.Minute)

					continue
				}

				// build transaction
				message := &protocol.Message{
					Network: protocol.NetworkPolygon,
				}
				if transactions, err = ethereum.BuildTransactions(ctx, message, transactions, s.ethClient); err != nil {
					loggerx.Global().Error("failed to build transactions", zap.Error(err))

					return
				}

				// lens worker
				internalTransactions := s.getInternalTransaction(ctx, transactions)

				// insert db
				err = s.upsertTransactions(ctx, internalTransactions)
				if err != nil {
					return
				}
			}
		}(eventHash, contractAddress)
	}
	wg.Wait()
}

func (s *service) getLensLogs(ctx context.Context, eventHash common.Hash, contractAddress common.Address) ([]*model.Transaction, error) {
	tracer := otel.Tracer("lens")
	_, trace := tracer.Start(ctx, "len:GetLensLogs")
	var internalTransactions = []*model.Transaction{}
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

	parameter := pregod_etl.GetLogsRequest{
		Address:    contractAddress.String(),
		TopicFirst: eventHash.String(),
	}

	cacheKey := fmt.Sprintf(lensLogsCacheKey, eventHash.String())
	cacheInfo, _ := cache.Global().Get(ctx, cacheKey).Result()
	if cacheList := strings.Split(cacheInfo, ":"); len(cacheList) == 2 {
		parameter.BlockNumberFrom, _ = strconv.ParseInt(cacheList[1], 10, 64)
		parameter.Cursor = cacheList[2]
	}

	// get logs from pregod_etl
	result, err := s.etlClient.GetLogs(ctx, "lens", parameter)
	if err != nil {
		loggerx.Global().Error("[lens] GetLensLogs: etlClient GetLogs error", zap.Error(err))

		return nil, err
	}
	if len(result.Result) == 0 {
		return nil, nil
	}

	for _, transfer := range result.Result {
		// get profile by profileId
		profileId, err := hexutil.DecodeBig(transfer.TopicSecond)
		if err != nil {
			loggerx.Global().Error("[lens] GetLensLogs: hexutil.Decode error", zap.Error(err))

			continue
		}
		profile, err := s.GetProfileById(ctx, profileId)
		if err != nil {
			loggerx.Global().Error("[lens] GetLensLogs: GetProfileById error", zap.Error(err))

			continue
		}

		transaction := &model.Transaction{
			BlockNumber: transfer.BlockNumber.BigInt().Int64(),
			Hash:        transfer.TransactionHash,
			Index:       transfer.TransactionIndex.BigInt().Int64(),
			Network:     protocol.NetworkPolygon,
			Source:      protocol.SourcePregodETL,
			Platform:    protocol.PlatformLens,
			Transfers:   make([]model.Transfer, 0),
			Owner:       strings.ToLower(profile.FollowNFT.String()),
		}

		internalTransactions = append(internalTransactions, transaction)
	}

	// set cache
	cacheInfo = fmt.Sprintf("%v:%v", result.Result[len(result.Result)-1].BlockNumber, result.Cursor)
	cache.Global().Set(ctx, cacheKey, cacheInfo, 7*24*time.Hour)

	return internalTransactions, nil
}

func (s *service) getProfileById(ctx context.Context, profileId *big.Int) (contract.DataTypesProfileStruct, error) {
	tracer := otel.Tracer("len")
	_, trace := tracer.Start(ctx, "len:GetProfileById")
	var profile contract.DataTypesProfileStruct
	var err error
	defer func() { opentelemetry.Log(trace, profileId, profile, err) }()

	// rpc
	iLensHub, err := contract.NewILensHub(lens.HubProxyContractAddress, s.ethClient)
	if err != nil {
		loggerx.Global().Error("[len] NewILensHub error", zap.Error(err))

		return profile, err
	}

	profile, err = iLensHub.GetProfile(&bind.CallOpts{}, profileId)
	if err != nil {
		loggerx.Global().Error("[len] iLensHub DefaultProfile error", zap.Error(err))

		return profile, err
	}

	return profile, nil
}

func (s *service) getInternalTransaction(ctx context.Context, transactions []*model.Transaction) []model.Transaction {
	var mu sync.Mutex
	internalTransactions := []model.Transaction{}
	opt := lop.NewOption().WithConcurrency(ethereum.RPCMaxConcurrency)
	lop.ForEach(transactions, func(transaction *model.Transaction, i int) {
		addressTo := common.HexToAddress(transaction.AddressTo)
		if addressTo != lens.HubProxyContractAddress && addressTo != lens.ProfileProxyContractAddress {
			return
		}

		transferMap := make(map[int64]model.Transfer)
		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)
		transaction.Transfers = append(transaction.Transfers, transferMap[protocol.IndexVirtual])

		// get receipt
		internalTransfers, err := s.commWorkerClient.HandleReceipt(ctx, &transaction)
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
		internalTransactions = append(internalTransactions, *transaction)
		mu.Unlock()
	}, opt)

	return internalTransactions
}

func (s *service) upsertTransactions(ctx context.Context, transactions []model.Transaction) error {
	dbChunkSize := 800
	transfers := []model.Transfer{}
	updatedTransactions := []model.Transaction{}

	for _, transaction := range transactions {
		addresses := strset.New(transaction.AddressFrom, transaction.AddressTo)

		// Ignore empty transactions
		internalTransfers := make([]model.Transfer, 0)

		for _, transfer := range transaction.Transfers {
			if bytes.Equal(transfer.Metadata, metadata.Default) {
				continue
			}

			internalTransfers = append(internalTransfers, transfer)
		}

		if len(internalTransfers) == 0 {
			continue
		}

		// Handle all transfers
		for _, transfer := range transaction.Transfers {
			// Ignore empty transfer
			if bytes.Equal(transfer.Metadata, metadata.Default) {
				continue
			}

			transfers = append(transfers, transfer)

			if transfer.AddressFrom != "" && transfer.AddressFrom != ethereum.AddressGenesis.String() {
				addresses.Add(transfer.AddressFrom)
			}

			if transfer.AddressTo != "" && transfer.AddressTo != ethereum.AddressGenesis.String() {
				addresses.Add(transfer.AddressTo)
			}
		}

		transaction.Addresses = addresses.List()
		updatedTransactions = append(updatedTransactions, transaction)
	}

	for _, ts := range lo.Chunk(updatedTransactions, dbChunkSize) {
		if err := c.databaseClient.
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			Create(ts).Error; err != nil {
			loggerx.Global().Error("failed to upsert transactions", zap.Error(err))

			return err
		}
	}

	for _, ts := range lo.Chunk(transfers, dbChunkSize) {
		if err := c.databaseClient.
			Clauses(clause.OnConflict{
				UpdateAll: true,
				DoUpdates: clause.AssignmentColumns([]string{"metadata"}),
			}).
			Create(ts).Error; err != nil {
			loggerx.Global().Error("failed to upsert transfers", zap.Error(err))

			return err
		}
	}
}
