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
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/alchemy"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource_asset"
	alchemy_asset "github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource_asset/alchemy"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/trigger"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/trigger/ens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/collectible/marketplace"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/collectible/poap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/donation/gitcoin"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/exchange/swap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/governance/snapshot"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/mirror"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction"
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

var _ command.Interface = &Server{}

type Server struct {
	config             *config.Config
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	rabbitmqQueue      rabbitmq.Queue
	rabbitmqAssetQueue rabbitmq.Queue
	databaseClient     *gorm.DB
	redisClient        *redis.Client
	datasources        []datasource.Datasource
	datasourcesAsset   []datasource_asset.Datasource
	workers            []worker.Worker
	triggers           []trigger.Trigger
	employer           *shedlock.Employer
}

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

	s.databaseClient, err = database.Dial(s.config.Postgres.String(), true)
	if err != nil {
		return err
	}

	database.ReplaceGlobal(s.databaseClient)

	if s.redisClient, err = cache.Dial(s.config.Redis); err != nil {
		return err
	}

	err = s.InitializeMQ()
	if err != nil {
		return err
	}

	alchemyDatasource, err := alchemy.New(s.config.RPC)
	if err != nil {
		return err
	}

	blockscoutDatasource, err := blockscout.New(s.config.RPC)
	if err != nil {
		return err
	}

	s.datasources = []datasource.Datasource{
		alchemyDatasource, moralis.New(s.config.Moralis.Key, s.config.RPC), arweave.New(), blockscoutDatasource, zksync.New(),
	}

	swapWorker, err := swap.New(s.config.RPC, s.employer, s.databaseClient)
	if err != nil {
		return err
	}

	ethereumClientMap, err := ethereum.New(s.config.RPC)
	if err != nil {
		return err
	}

	s.triggers = []trigger.Trigger{
		ens.New(),
	}

	s.workers = []worker.Worker{
		swapWorker,
		marketplace.New(s.databaseClient, ethereumClientMap),
		poap.New(ethereumClientMap),
		mirror.New(),
		gitcoin.New(s.databaseClient, s.redisClient, ethereumClientMap),
		snapshot.New(s.databaseClient, s.redisClient),
		transaction.New(s.databaseClient, ethereumClientMap),
	}

	s.employer = shedlock.New(s.redisClient)

	for _, internalWorker := range s.workers {
		if err := internalWorker.Initialize(context.Background()); err != nil {
			return err
		}

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
	alchemyAssetDatasource, err := alchemy_asset.New(s.config.RPC, ethereumClientMap)
	if err != nil {
		return err
	}

	s.datasourcesAsset = []datasource_asset.Datasource{
		alchemyAssetDatasource,
	}

	return nil
}

func (s *Server) InitializeMQ() (err error) {
	s.rabbitmqConnection, err = rabbitmq.Dial(s.config.RabbitMQ.String())
	if err != nil {
		return err
	}

	s.rabbitmqChannel, err = s.rabbitmqConnection.Channel()
	if err != nil {
		return err
	}

	if err := s.rabbitmqChannel.ExchangeDeclare(
		protocol.ExchangeJob, "direct", true, false, false, false, nil,
	); err != nil {
		return err
	}

	if s.rabbitmqQueue, err = s.rabbitmqChannel.QueueDeclare(
		protocol.IndexerWorkQueue, false, false, false, false, nil,
	); err != nil {
		return err
	}

	if err := s.rabbitmqChannel.QueueBind(
		s.rabbitmqQueue.Name, protocol.IndexerWorkRoutingKey, protocol.ExchangeJob, false, nil,
	); err != nil {
		return err
	}

	// asset mq
	if s.rabbitmqAssetQueue, err = s.rabbitmqChannel.QueueDeclare(
		protocol.IndexerAssetQueue, false, false, false, false, nil,
	); err != nil {
		return err
	}

	if err := s.rabbitmqChannel.QueueBind(
		s.rabbitmqAssetQueue.Name, protocol.IndexerAssetRoutingKey, protocol.ExchangeJob, false, nil,
	); err != nil {
		return err
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
		deliveryCh, err := s.rabbitmqChannel.Consume(s.rabbitmqQueue.Name, "", true, false, false, false, nil)
		if err != nil {
			return
		}

		for delivery := range deliveryCh {
			message := protocol.Message{}
			if err := json.Unmarshal(delivery.Body, &message); err != nil {
				logger.Global().Error("failed to unmarshal message", zap.Error(err))

				continue
			}

			go func() {
				if err := s.handle(context.Background(), &message); err != nil {
					logger.Global().Error("failed to handle message", zap.Error(err), zap.String("address", message.Address), zap.String("network", message.Network))
				}
			}()
		}
	}()

	go func() {
		deliveryAssetCh, err := s.rabbitmqChannel.Consume(s.rabbitmqAssetQueue.Name, "", true, false, false, false, nil)
		if err != nil {
			return
		}

		for delivery := range deliveryAssetCh {
			message := protocol.Message{}
			if err := json.Unmarshal(delivery.Body, &message); err != nil {
				logger.Global().Error("failed to unmarshal message", zap.Error(err))

				continue
			}

			go func() {
				if err := s.handleAsset(context.Background(), &message); err != nil {
					logger.Global().Error("failed to handle asset message", zap.Error(err), zap.String("address", message.Address), zap.String("network", message.Network))
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

	defer s.employer.UnLock(lockKey)
	defer cancel()

	// convert address to lowercase
	message.Address = strings.ToLower(message.Address)
	tracer := otel.Tracer("indexer")

	ctx, handlerSpan := tracer.Start(ctx, "indexer:handler")

	handlerSpan.SetAttributes(
		attribute.String("network", message.Network),
	)

	defer handlerSpan.End()

	logger.Global().Info("start indexing data", zap.String("address", message.Address), zap.String("network", message.Network))

	// upsert address status
	go s.upsertAddress(model.Address{
		Address: message.Address,
		Network: message.Network,
		Status:  model.AddressStatusRunning,
	})

	// Get data from trigger
	for _, internalTrigger := range s.triggers {
		for _, network := range internalTrigger.Networks() {
			if message.Network == network {
				go func(internalTrigger trigger.Trigger) {
					if err := internalTrigger.Handle(ctx, message); err != nil {
						logger.Global().Error("failed to handle trigger", zap.Error(err), zap.String("address", message.Address), zap.String("network", message.Network))
					}
				}(internalTrigger)

				break
			}
		}
	}

	// Get data from datasources
	var transactions []model.Transaction

	for _, datasource := range s.datasources {
		for _, network := range datasource.Networks() {
			if network == message.Network {
				// Get the time of the latest data for this address and network
				var result struct {
					Timestamp   time.Time `gorm:"column:timestamp"`
					BlockNumber int64     `gorm:"column:block_number"`
				}

				if err := s.databaseClient.
					Model((*model.Transaction)(nil)).
					Select("COALESCE(timestamp, 'epoch'::timestamp) AS timestamp, COALESCE(block_number, 0) AS block_number").
					Where("owner = ?", message.Address).
					Where("network = ?", message.Network).
					Order("timestamp DESC").
					Limit(1).
					First(&result).
					Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					continue
				}

				message.Timestamp = result.Timestamp
				message.BlockNumber = result.BlockNumber

				internalTransactions, err := datasource.Handle(ctx, message)
				// Avoid blocking indexed workers
				if err != nil {
					logger.Global().Error("datasource handle failed", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address), zap.String("datasource", datasource.Name()))

					continue
				}

				transactions = append(transactions, internalTransactions...)
			}
		}
	}

	transactionsMap := getTransactionsMap(transactions)

	defer func() {
		transfers := 0

		for _, transaction := range transactions {
			transfers += len(transaction.Transfers)
		}

		logger.Global().Info("indexed data completion", zap.String("address", message.Address), zap.String("network", message.Network), zap.Int("transactions", len(transactions)), zap.Int("transfers", transfers))

		// upsert address status
		go s.upsertAddress(model.Address{
			Address: message.Address,
			Network: message.Network,
			Status:  model.AddressStatusDone,
		})
	}()

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

	logger.Global().Info("start indexing asset data", zap.String("address", message.Address), zap.String("network", message.Network))

	// Get data from datasources
	var assets []model.Asset

	for _, datasource := range s.datasourcesAsset {
		for _, network := range datasource.Networks() {
			if network == message.Network {
				internalAssets, err := datasource.Handle(ctx, message)
				// Avoid blocking indexed workers
				if err != nil {
					logger.Global().Error("datasource handle failed", zap.Error(err))
					continue
				}

				assets = append(assets, internalAssets...)
			}
		}
	}

	// set db
	if err := s.databaseClient.
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

func (s *Server) upsertTransactions(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (err error) {
	tracer := otel.Tracer("indexer")
	_, span := tracer.Start(ctx, "indexer:upsertTransactions")

	defer opentelemetry.Log(span, len(transactions), nil, err)

	tx := s.databaseClient.Begin()

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
			logger.Global().Error("failed to upsert transactions", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address))

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
			logger.Global().Error("failed to upsert transfers", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address))

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
				internalTransactions, err := worker.Handle(ctx, message, transactions)
				if err != nil {
					logger.Global().Error("worker handle failed", zap.Error(err), zap.String("worker", worker.Name()), zap.String("network", network))

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

	return s.upsertTransactions(ctx, message, transactions)
}

func (s *Server) upsertAddress(address model.Address) error {
	if err := s.databaseClient.
		Clauses(clause.OnConflict{
			UpdateAll: true,
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).
		Create(&address).Error; err != nil {
		logger.Global().Error("failed to upsert address", zap.Error(err), zap.String("network", address.Network), zap.String("address", address.Address))
		return err
	}
	return nil
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
