package lens

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/datasource/pregod_etl"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
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

	lensLogsCacheKey = "crawler:lens:%v" // eventHash -> blocknumber:cursor
)

type service struct {
	config           *config.Config
	etlClient        *pregod_etl.Client
	commWorkerClient *lens_comm.Client
}

func New(config *config.Config) crawler.Crawler {
	var err error

	crawler := &service{config: config}

	if crawler.etlClient, err = pregod_etl.NewClient(protocol.NetworkPolygon, config.RPC.PregodETL.Polygon.HTTP); err != nil {
		logrus.Error("[lens] pregod_etl.NewClient error, ", err)
		return nil
	}

	crawler.commWorkerClient = lens_comm.New()

	return crawler
}

func (s *service) Name() string {
	return protocol.PlatformLens
}

func (s *service) Run() error {
	ctx := context.Background()
	var wg sync.WaitGroup
	for eventHash, contractAddress := range lens.SupportLensEvents {
		wg.Add(1)
		go func(eventHash common.Hash, contractAddress common.Address) {
			defer wg.Done()

			for {
				// get lens logs
				transactions, err := s.getLensLogs(ctx, eventHash, contractAddress)
				if err != nil {
					logrus.Error("[lens] Run: GetLensLogs error, ", err)

					return
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
					Network: protocol.NetworkPolygon,
				}
				if transactions, err = ethereum.BuildTransactions(ctx, message, transactions); err != nil {
					logrus.Error("failed to build transactions, ", err)

					continue
				}

				// lens worker
				internalTransactions := s.getInternalTransaction(ctx, transactions)

				// insert db
				err = database.UpsertTransactions(ctx, internalTransactions)
				if err != nil {
					continue
				}
			}
		}(eventHash, contractAddress)
	}
	wg.Wait()
	return nil
}

func (s *service) getLensLogs(ctx context.Context, eventHash common.Hash, contractAddress common.Address) ([]*model.Transaction, error) {
	tracer := otel.Tracer("lens")
	_, trace := tracer.Start(ctx, "len:GetLensLogs")
	var internalTransactions []*model.Transaction
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

	parameter := pregod_etl.GetLogsRequest{
		Address:    contractAddress.String(),
		TopicFirst: eventHash.String(),
		Limit:      100,
		Order:      "asc",
	}

	cacheKey := fmt.Sprintf(lensLogsCacheKey, eventHash.String())
	cacheInfo, _ := cache.Global().Get(ctx, cacheKey).Result()
	if cacheList := strings.Split(cacheInfo, ":"); len(cacheList) == 2 {
		parameter.BlockNumberFrom, _ = strconv.ParseInt(cacheList[0], 10, 64)
		parameter.Cursor = cacheList[1]
	}

	// get logs from pregod_etl
	result, err := s.etlClient.GetLogs(ctx, "lens", parameter)
	if err != nil {
		logrus.Error("[lens] GetLensLogs: etlClient GetLogs error, ", err)

		return nil, err
	}

	for _, transfer := range result.Result {
		// get profile by profileId
		profileId := big.NewInt(0).SetBytes(common.FromHex(transfer.TopicSecond))
		profile, err := s.getLensOwnerAddressById(ctx, profileId)
		if err != nil {
			logrus.Errorf("[lens] GetLensLogs: getLensAddressById error, %v, profileId: %v", err, profileId)

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
			Owner:       strings.ToLower(profile.String()),
		}

		internalTransactions = append(internalTransactions, transaction)
	}

	// set cache
	if len(result.Result) == 0 {
		return internalTransactions, nil
	}

	cacheInfo = fmt.Sprintf("%v:%v", result.Result[len(result.Result)-1].BlockNumber, result.Cursor)
	cache.Global().Set(ctx, cacheKey, cacheInfo, 7*24*time.Hour)

	return internalTransactions, nil
}

func (s *service) getLensOwnerAddressById(ctx context.Context, profileId *big.Int) (common.Address, error) {
	tracer := otel.Tracer("len")
	_, trace := tracer.Start(ctx, "len:getLensAddressById")
	var owner common.Address
	var err error
	defer func() { opentelemetry.Log(trace, profileId, owner, err) }()

	// rpc
	ethclient, err := ethclientx.Global(protocol.NetworkPolygon)
	if err != nil {
		return owner, err
	}

	lensHub, err := contract.NewHub(lens.HubProxyContractAddress, ethclient)
	if err != nil {
		logrus.Error("[len] NewILensHub error, ", err)

		return owner, err
	}

	owner, err = lensHub.OwnerOf(&bind.CallOpts{}, profileId)
	if err != nil {
		logrus.Error("[len] iLensHub GetFollowNFT error, ", err)

		return owner, err
	}

	return owner, nil
}

func (s *service) getInternalTransaction(ctx context.Context, transactions []*model.Transaction) []*model.Transaction {
	var mu sync.Mutex
	var internalTransactions []*model.Transaction
	opt := lop.NewOption().WithConcurrency(10)
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
