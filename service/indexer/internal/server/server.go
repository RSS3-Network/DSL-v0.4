package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/alchemy"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/eip1577"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	eth_etl "github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/pregod_etl/ethereum"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource_asset"
	alchemy_asset "github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource_asset/alchemy"
	rabbitmqx "github.com/naturalselectionlabs/pregod/service/indexer/internal/rabbitmq"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/collectible/marketplace"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/collectible/poap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/donation/gitcoin"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/exchange/liquidity"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/exchange/swap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/governance/snapshot"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/mirror"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction/bridge"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/samber/lo"
	"github.com/scylladb/go-set/strset"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Server struct {
	config           *config.Config
	datasources      []datasource.Datasource
	datasourcesAsset []datasource_asset.Datasource
	workers          []worker.Worker
	employer         *shedlock.Employer
}

var _ command.Interface = &Server{}

func (s *Server) Initialize() (err error) {
	var exporter trace.SpanExporter

	if s.config.OpenTelemetry == nil {
		if exporter, err = opentelemetry.DialWithPath(opentelemetry.DefaultPath); err != nil {
			return err
		}
	} else if s.config.OpenTelemetry.Enabled {
		if exporter, err = opentelemetry.DialWithURL(s.config.OpenTelemetry.String()); err != nil {
			return err
		}
	}

	otel.SetTracerProvider(trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("pregod-1-1-indexer"),
			semconv.ServiceVersionKey.String(protocol.Version),
		)),
	))

	databaseClient, err := database.Dial(s.config.Postgres.String(), false)
	if err != nil {
		return err
	}

	database.ReplaceGlobal(databaseClient)

	ethDbClient, err := database.Dial(s.config.EthereumEtl.String(), false)
	if err != nil {
		return err
	}

	database.ReplaceEthDb(ethDbClient)

	redisClient, err := cache.Dial(s.config.Redis)
	if err != nil {
		return err
	}

	cache.ReplaceGlobal(redisClient)

	ipfs.New(s.config.RPC.IPFS.IO)

	err = rabbitmqx.InitializeMQ(s.config.RabbitMQ)
	if err != nil {
		return err
	}

	alchemyDatasource, err := alchemy.New(s.config.RPC)
	if err != nil {
		return err
	}

	blockscoutDatasource, err := blockscout.New()
	if err != nil {
		return err
	}

	s.datasources = []datasource.Datasource{
		alchemyDatasource,
		moralis.New(s.config.Moralis.Key),
		arweave.New(),
		blockscoutDatasource,
		zksync.New(),
		eth_etl.New(),
		eip1577.New(s.employer),
	}

	swapWorker, err := swap.New(s.employer)
	if err != nil {
		return err
	}

	ethereumClientMap, err := ethclientx.Dial(s.config.RPC, protocol.EthclientNetworks)
	if err != nil {
		return err
	}

	for network, client := range ethereumClientMap {
		ethclientx.ReplaceGlobal(network, client)
	}

	s.workers = []worker.Worker{
		liquidity.New(),
		swapWorker,
		bridge.New(),
		marketplace.New(),
		poap.New(),
		mirror.New(),
		gitcoin.New(),
		snapshot.New(),
		crossbell.New(),
		transaction.New(),
	}

	s.employer = shedlock.New()

	for _, internalWorker := range s.workers {
		loggerx.Global().Info("start initializing worker", zap.String("worker", internalWorker.Name()))

		startTime := time.Now()

		if err := internalWorker.Initialize(context.Background()); err != nil {
			return err
		}

		loggerx.Global().Info("initialize worker completion", zap.String("worker", internalWorker.Name()), zap.Duration("duration", time.Since(startTime)))

		if internalWorker.Jobs() == nil {
			continue
		}

		for _, job := range internalWorker.Jobs() {
			if err := s.employer.AddJob(job.Name(), job.Spec(), job.Timeout(), worker.NewCronJob(s.employer, job)); err != nil {
				return err
			}
		}
	}

	// asset
	alchemyAssetDatasource, err := alchemy_asset.New(s.config.RPC)
	if err != nil {
		return err
	}

	s.datasourcesAsset = []datasource_asset.Datasource{
		alchemyAssetDatasource,
	}

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	s.employer.Start()

	defer s.employer.Stop()

	go func() {
		deliveryCh, err := rabbitmqx.GetRabbitmqChannel().Consume(rabbitmqx.GetRabbitmqQueue().Name, "", true, false, false, false, nil)
		if err != nil {
			return
		}

		for delivery := range deliveryCh {
			message := protocol.Message{}
			if err := json.Unmarshal(delivery.Body, &message); err != nil {
				loggerx.Global().Error("failed to unmarshal message", zap.Error(err))

				continue
			}

			go func() {
				if err := s.handle(context.Background(), &message); err != nil {
					loggerx.Global().Error("failed to handle message", zap.Error(err), zap.String("address", message.Address), zap.String("network", message.Network))
				}
			}()
		}
	}()

	go func() {
		deliveryAssetCh, err := rabbitmqx.GetRabbitmqChannel().Consume(rabbitmqx.GetRabbitmqAssetQueue().Name, "", true, false, false, false, nil)
		if err != nil {
			return
		}

		for delivery := range deliveryAssetCh {
			message := protocol.Message{}
			if err := json.Unmarshal(delivery.Body, &message); err != nil {
				loggerx.Global().Error("failed to unmarshal message", zap.Error(err))

				continue
			}

			go func() {
				if err := s.handleAsset(context.Background(), &message); err != nil {
					loggerx.Global().Error("failed to handle asset message", zap.Error(err), zap.String("address", message.Address), zap.String("network", message.Network))
				}
			}()
		}
	}()

	stopchan := make(chan os.Signal, 1)
	signal.Notify(stopchan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-stopchan

	return nil
}

func (s *Server) handle(ctx context.Context, message *protocol.Message) (err error) {
	lockKey := fmt.Sprintf("indexer:%v:%v", message.Address, message.Network)

	if !s.employer.DoLock(lockKey, 2*time.Minute) {
		return fmt.Errorf("%v lock", lockKey)
	}

	cctx, cancel := context.WithCancel(context.Background())
	go func(cctx context.Context) {
		for {
			time.Sleep(time.Second)
			if err := s.employer.Renewal(cctx, lockKey, time.Minute); err != nil {
				return
			}
		}
	}(cctx)

	var transactions []model.Transaction
	defer func() {
		cancel()
		s.employer.UnLock(lockKey)

		transfers := 0

		for _, transaction := range transactions {
			transfers += len(transaction.Transfers)
		}

		loggerx.Global().Info("indexed data completion", zap.String("address", message.Address), zap.String("network", message.Network), zap.Int("transactions", len(transactions)), zap.Int("transfers", transfers))

		// upsert address status
		go s.upsertAddress(ctx, model.Address{
			Address: message.Address,
		})
	}()

	// convert address to lowercase
	message.Address = strings.ToLower(message.Address)
	tracer := otel.Tracer("indexer")

	ctx, handlerSpan := tracer.Start(ctx, "indexer:handler")

	handlerSpan.SetAttributes(
		attribute.String("network", message.Network),
	)

	defer handlerSpan.End()

	loggerx.Global().Info("start indexing data", zap.String("address", message.Address), zap.String("network", message.Network))

	// Get the time of the latest data for this address and network
	var result struct {
		Timestamp   time.Time `gorm:"column:timestamp"`
		BlockNumber int64     `gorm:"column:block_number"`
	}

	if !message.Reindex {
		if err := database.Global().
			Model((*model.Transaction)(nil)).
			Select("COALESCE(timestamp, 'epoch'::timestamp) AS timestamp, COALESCE(block_number, 0) AS block_number").
			Where("owner = ?", message.Address).
			Where("network = ?", message.Network).
			Order("timestamp DESC").
			Limit(1).
			First(&result).
			Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	message.Timestamp = result.Timestamp
	message.BlockNumber = result.BlockNumber

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, ds := range s.datasources {
		wg.Add(1)
		go func(message *protocol.Message, datasource datasource.Datasource) {
			defer wg.Done()
			for _, network := range datasource.Networks() {
				if network == message.Network {
					loggerx.Global().Info("start datasource", zap.String("datasource", datasource.Name()), zap.String("address", message.Address))
					startTime := time.Now()

					// handle
					internalTransactions, err := datasource.Handle(ctx, message)

					// log
					loggerx.Global().Info("datasource completion", zap.String("datasource", datasource.Name()), zap.String("address", message.Address), zap.Duration("duration", time.Since(startTime)))

					// Avoid blocking indexed workers
					if err != nil {
						loggerx.Global().Error("datasource handle failed", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address), zap.String("datasource", datasource.Name()))

						continue
					}
					mu.Lock()
					transactions = append(transactions, internalTransactions...)
					mu.Unlock()
				}
			}
		}(message, ds)

	}
	wg.Wait()

	transactionsMap := getTransactionsMap(transactions)

	return s.handleWorkers(ctx, message, transactions, transactionsMap)
}

func (s *Server) handleAsset(ctx context.Context, message *protocol.Message) (err error) {
	lockKey := fmt.Sprintf("indexer_asset:%v:%v", message.Address, message.Network)

	if !s.employer.DoLock(lockKey, 2*time.Minute) {
		return fmt.Errorf("%v lock", lockKey)
	}

	cctx, cancel := context.WithCancel(context.Background())
	go func(cctx context.Context) {
		for {
			time.Sleep(time.Second)
			if err := s.employer.Renewal(cctx, lockKey, time.Minute); err != nil {
				return
			}
		}
	}(cctx)

	defer s.employer.UnLock(lockKey)
	defer cancel()

	// convert address to lowercase
	message.Address = strings.ToLower(message.Address)
	tracer := otel.Tracer("indexer")

	ctx, handlerSpan := tracer.Start(ctx, "indexer:handleAsset")

	handlerSpan.SetAttributes(
		attribute.String("network", message.Network),
	)

	defer handlerSpan.End()

	loggerx.Global().Info("start indexing asset data", zap.String("address", message.Address), zap.String("network", message.Network))

	// Get data from datasources
	var assets []model.Asset

	for _, datasource := range s.datasourcesAsset {
		for _, network := range datasource.Networks() {
			if network == message.Network {
				internalAssets, err := datasource.Handle(ctx, message)
				// Avoid blocking indexed workers
				if err != nil {
					loggerx.Global().Error("datasource handle failed", zap.Error(err))
					continue
				}

				assets = append(assets, internalAssets...)
			}
		}
	}

	// set db
	if err := database.Global().
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).
		Create(assets).Error; err != nil {
		return err
	}

	return nil
}

func getTransactionsMap(transactions []model.Transaction) map[string]model.Transaction {
	transactionsMap := make(map[string]model.Transaction)

	for _, t := range transactions {
		transactionsMap[t.Hash] = t
	}

	return transactionsMap
}

func transactionsMap2Array(transactionsMap map[string]model.Transaction) []model.Transaction {
	transactions := make([]model.Transaction, 0)

	for _, t := range transactionsMap {
		transactions = append(transactions, t)
	}

	return transactions
}

func (s *Server) upsertTransactions(ctx context.Context, message *protocol.Message, tx *gorm.DB, transactions []model.Transaction) (err error) {
	tracer := otel.Tracer("indexer")
	_, span := tracer.Start(ctx, "indexer:upsertTransactions")

	defer opentelemetry.Log(span, len(transactions), nil, err)

	dbChunkSize := 800

	var (
		transfers           []model.Transfer
		updatedTransactions []model.Transaction
	)

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
			// handle unsupported Unicode escape sequence
			if bytes.Contains(transfer.Metadata, []byte(`\u0000`)) {
				transfer.Metadata = bytes.ReplaceAll(transfer.Metadata, []byte(`\u0000`), []byte{})
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
		if err = tx.
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			Create(ts).Error; err != nil {
			loggerx.Global().Error("failed to upsert transactions", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address))

			tx.Rollback()

			return err
		}
	}

	for _, ts := range lo.Chunk(transfers, dbChunkSize) {
		if err = tx.
			Clauses(clause.OnConflict{
				UpdateAll: true,
				DoUpdates: clause.AssignmentColumns([]string{"metadata"}),
			}).
			Create(ts).Error; err != nil {
			loggerx.Global().Error("failed to upsert transfers", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address))

			tx.Rollback()

			return err
		}
	}

	return tx.Commit().Error
}

func (s *Server) handleWorkers(ctx context.Context, message *protocol.Message, transactions []model.Transaction, transactionsMap map[string]model.Transaction) (err error) {
	tracer := otel.Tracer("indexer")
	ctx, span := tracer.Start(ctx, "indexer:handleWorkers")

	defer opentelemetry.Log(span, message, transactions, err)

	// Using workers to clean data
	for _, worker := range s.workers {
		for _, network := range worker.Networks() {
			if network == message.Network {
				// log
				loggerx.Global().Info("start worker", zap.String("worker", worker.Name()), zap.String("address", message.Address))
				startTime := time.Now()

				internalTransactions, err := worker.Handle(ctx, message, transactions)

				// log
				loggerx.Global().Info("worker completion", zap.String("worker", worker.Name()), zap.String("address", message.Address), zap.Duration("duration", time.Since(startTime)))

				if err != nil {
					loggerx.Global().Error("worker handle failed", zap.Error(err), zap.String("worker", worker.Name()), zap.String("network", network))

					continue
				}

				if len(internalTransactions) == 0 {
					continue
				}

				newTransactionMap := getTransactionsMap(internalTransactions)
				for _, t := range newTransactionMap {
					transactionsMap[t.Hash] = t
				}

				transactions = transactionsMap2Array(transactionsMap)
			}
		}
	}

	// Open a database transaction
	tx := database.Global().WithContext(ctx).Begin()

	// Delete data from this address and reindex it
	if message.Reindex {
		var hashes []string

		// TODO Use the owner to replace hashes field
		// Get all hashes of this address on this network
		if err := tx.
			Model((*model.Transaction)(nil)).
			Where("network = ? AND owner = ?", message.Network, message.Address).
			Pluck("hash", &hashes).
			Error; err != nil {
			return err
		}

		if err := tx.Where("network = ? AND hash IN (SELECT * FROM UNNEST(?::TEXT[]))", message.Network, pq.Array(hashes)).Delete(&model.Transaction{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Where("network = ? AND transaction_hash IN (SELECT * FROM UNNEST(?::TEXT[]))", message.Network, pq.Array(hashes)).Delete(&model.Transfer{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return s.upsertTransactions(ctx, message, tx, transactions)
}

func (s *Server) upsertAddress(ctx context.Context, address model.Address) {
	for _, network := range protocol.SupportNetworks {
		key := fmt.Sprintf("indexer:%v:%v", address.Address, network)
		n, err := cache.Global().Exists(ctx, key).Result()
		switch {
		case err != nil:
			return
		case n == 1: // exists
			address.IndexingNetworks = append(address.IndexingNetworks, network)
		default: // not exists (unlocked)
			address.DoneNetworks = append(address.DoneNetworks, network)
		}
	}

	if len(address.IndexingNetworks) == 0 {
		address.Status = true
	}

	if err := database.Global().
		Model(model.Address{}).
		Clauses(clause.OnConflict{
			UpdateAll: true,
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).
		Create(map[string]interface{}{
			"address":           address.Address,
			"status":            address.Status,
			"done_networks":     address.DoneNetworks,
			"indexing_networks": address.IndexingNetworks,
		}).Error; err != nil {
		loggerx.Global().Error("failed to upsert address", zap.Error(err), zap.String("address", address.Address))
	}
	s.publishRefreshMessage(ctx, address)
}

func (s *Server) publishRefreshMessage(ctx context.Context, address model.Address) {
	tracer := otel.Tracer("publishRefreshMessage")
	_, rabbitmqSnap := tracer.Start(ctx, "rabbitmq")

	defer rabbitmqSnap.End()
	// refresh or new address
	address.UpdatedAt = time.Now().Add(-1 * time.Second)
	messageData, err := json.Marshal(&protocol.RefreshMessage{
		Address: address,
	})
	if err != nil {
		return
	}
	if err := rabbitmqx.GetRabbitmqChannel().Publish(protocol.ExchangeRefresh, "", false, false, rabbitmq.Publishing{
		ContentType: protocol.ContentTypeJSON,
		Body:        messageData,
	}); err != nil {
		loggerx.Global().Error("failed to publish refresh message to mq", zap.Error(err), zap.String("address", address.Address))
	}
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
