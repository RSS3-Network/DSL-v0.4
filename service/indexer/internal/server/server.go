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

	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource_asset"
	alchemy_asset "github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource_asset/alchemy"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/mq"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/trigger"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/trigger/ens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/exchange/swap"
	"github.com/samber/lo"
	"github.com/scylladb/go-set/strset"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Server struct{}

var _ command.Interface = &Server{}

func (s *Server) Initialize() (err error) {
	if err := database.Dial(config.ConfigIndexer.Postgres.String(), true); err != nil {
		return err
	}

	if err := cache.Dial(config.ConfigIndexer.Redis); err != nil {
		return err
	}

	ipfs.New(config.ConfigIndexer.RPC.IPFS.IO)

	if err := mq.Initialize(); err != nil {
		return err
	}

	if err := datasource.Initialize(); err != nil {
		return err
	}

	swapWorker, err := swap.New(config.ConfigIndexer.RPC, s.employer)
	if err != nil {
		return err
	}

	ethereumClientMap, err := ethereum.New(config.ConfigIndexer.RPC)
	if err != nil {
		return err
	}

	s.triggers = []trigger.Trigger{
		ens.New(),
	}

	// s.workers = []internalModel.Worker{
	// 	liquidity.New(ethereumClientMap),
	// 	swapWorker,
	// 	marketplace.New(ethereumClientMap),
	// 	poap.New(ethereumClientMap),
	// 	mirror.New(),
	// 	gitcoin.New(ethereumClientMap),
	// 	snapshot.New(),
	// 	lens_worker.New(ethereumClientMap),
	// 	transaction.New(ethereumClientMap),
	// }

	shedlock.New()

	// for _, internalWorker := range s.workers {
	// 	loggerx.Global().Info("start initializing worker", zap.String("worker", internalWorker.Name()))

	// 	startTime := time.Now()

	// 	if err := internalWorker.Initialize(context.Background()); err != nil {
	// 		return err
	// 	}

	// 	loggerx.Global().Info("initialize worker completion", zap.String("worker", internalWorker.Name()), zap.Duration("duration", time.Since(startTime)))

	// 	if internalWorker.Jobs() == nil {
	// 		continue
	// 	}

	// 	for _, job := range internalWorker.Jobs() {
	// 		if err := s.employer.AddJob(job.Name(), job.Spec(), job.Timeout(), worker.NewCronJob(s.employer, job)); err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// asset
	alchemyAssetDatasource, err := alchemy_asset.New(config.ConfigIndexer.RPC, ethereumClientMap)
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
		deliveryCh, err := s.rabbitmqChannel.Consume(s.rabbitmqQueue.Name, "", true, false, false, false, nil)
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
		deliveryAssetCh, err := s.rabbitmqChannel.Consume(s.rabbitmqAssetQueue.Name, "", true, false, false, false, nil)
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

	defer cancel()

	var transactions []model.Transaction
	defer func() {
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

	// Ignore triggers
	if !message.IgnoreTrigger {
		if err := s.executeTriggers(ctx, message); err != nil {
			zap.L().Error("failed to execute trigger", zap.Error(err), zap.String("address", message.Address), zap.String("network", message.Network))
		}
	}

	// Ignore notes
	if message.IgnoreNote {
		return nil
	}

	// Open a database transaction
	tx := database.Global().WithContext(ctx).Begin()

	for _, datasource := range s.datasources {
		for _, network := range datasource.Networks() {
			if network == message.Network {
				// Get the time of the latest data for this address and network
				var result struct {
					Timestamp   time.Time `gorm:"column:timestamp"`
					BlockNumber int64     `gorm:"column:block_number"`
				}

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

				if err := tx.
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
					loggerx.Global().Error("datasource handle failed", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address), zap.String("datasource", datasource.Name()))

					continue
				}

				transactions = append(transactions, internalTransactions...)
			}
		}
	}

	transactionsMap := getTransactionsMap(transactions)

	return s.handleWorkers(ctx, message, tx, transactions, transactionsMap)
}

func (s *Server) executeTriggers(ctx context.Context, message *protocol.Message) error {
	// Get data from trigger
	for _, internalTrigger := range s.triggers {
		for _, network := range internalTrigger.Networks() {
			if message.Network == network {
				go func(internalTrigger trigger.Trigger) {
					if err := internalTrigger.Handle(ctx, message); err != nil {
						loggerx.Global().Error("failed to handle trigger", zap.Error(err), zap.String("address", message.Address), zap.String("network", message.Network))
					}
				}(internalTrigger)

				break
			}
		}
	}

	return nil
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


func (s *Server) handleWorkers(ctx context.Context, message *protocol.Message, tx *gorm.DB, transactions []model.Transaction, transactionsMap map[string]model.Transaction) (err error) {
	tracer := otel.Tracer("indexer")
	ctx, span := tracer.Start(ctx, "indexer:handleWorkers")

	defer opentelemetry.Log(span, message, transactions, err)

	// Using workers to clean data
	for _, worker := range s.workers {
		for _, network := range worker.Networks() {
			if network == message.Network {
				internalTransactions, err := worker.Handle(ctx, message, transactions)
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

	return s.upsertTransactions(ctx, message, tx, transactions)
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
