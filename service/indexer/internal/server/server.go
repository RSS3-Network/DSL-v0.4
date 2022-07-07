package server

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/datasource/nft"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	poapworker "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/collectible/poap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/donation/gitcoin"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/exchange/swap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/governance/snapshot"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell"
	lensworker "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/lens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/mirror"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/alchemy"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction/coinmarketcap"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
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
	databaseClient     *gorm.DB
	redisClient        *redis.Client
	datasources        []datasource.Datasource
	workers            []worker.Worker
	indexedWorker      []worker.Worker
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
			semconv.ServiceNameKey.String("pregod2-indexer"),
			semconv.ServiceVersionKey.String("v1.0.0"),
		)),
	))

	nft.Initialize(s.config.Infura.ProjectID)

	s.databaseClient, err = database.Dial(s.config.Postgres.String(), true)
	if err != nil {
		return err
	}

	if s.redisClient, err = cache.Dial(s.config.Redis); err != nil {
		return err
	}

	coinmarketcap.Init(s.config.CoinMarketCap.APIKey)

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

	alchemyDatasource, err := alchemy.New(s.config.RPC)
	if err != nil {
		return err
	}

	s.datasources = []datasource.Datasource{
		alchemyDatasource, moralis.New(s.config.Moralis.Key, s.config.Infura.ProjectID), arweave.New(), blockscout.New(), zksync.New(),
	}

	s.indexedWorker = []worker.Worker{
		lensworker.New(s.databaseClient),
		snapshot.New(s.databaseClient, s.redisClient),
	}

	s.workers = []worker.Worker{
		transaction.New(s.databaseClient),
		swap.New(s.employer, s.databaseClient),
		mirror.New(),
		poapworker.New(),
		gitcoin.New(s.databaseClient, s.redisClient),
		crossbell.New(),
		snapshot.New(s.databaseClient, s.redisClient),
		lensworker.New(s.databaseClient),
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

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	s.employer.Start()

	defer s.employer.Stop()

	deliveryCh, err := s.rabbitmqChannel.Consume(s.rabbitmqQueue.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	for delivery := range deliveryCh {
		message := protocol.Message{}
		if err := json.Unmarshal(delivery.Body, &message); err != nil {
			logrus.Errorln(err)

			continue
		}

		go func() {
			if err := s.handle(context.Background(), &message); err != nil {
				logrus.Errorf("message.Address: %v, message.Network: %v, err: %v", message.Address, message.Network, err)
			}
		}()
	}

	return nil
}

func (s *Server) handle(ctx context.Context, message *protocol.Message) (err error) {
	// convert address to lowercase
	message.Address = strings.ToLower(message.Address)
	tracer := otel.Tracer("indexer")

	ctx, handlerSpan := tracer.Start(ctx, "indexer:handler")

	handlerSpan.SetAttributes(
		attribute.String("network", message.Network),
	)

	defer handlerSpan.End()

	logrus.Info(message.Address, " ", message.Network)

	// Get data from datasources
	var transactions []model.Transaction

	var datasourceError error

	for _, datasource := range s.datasources {
		for _, network := range datasource.Networks() {
			if network == message.Network {
				// Get the time of the latest data for this address and network
				var timestamp time.Time

				if err := s.databaseClient.
					Model((*model.Transaction)(nil)).
					Select("COALESCE(timestamp, 'epoch'::timestamp) AS timestamp").
					Where(map[string]interface{}{
						"address_from": message.Address,
						"network":      message.Network,
					}).
					Or(map[string]interface{}{
						"address_to": message.Address,
						"network":    message.Network,
					}).
					Order("timestamp DESC").
					Limit(1).
					Pluck("timestamp", &timestamp).
					Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}

				message.Timestamp = timestamp

				internalTransactions, err := datasource.Handle(ctx, message)
				// Avoid blocking indexed workers
				if err != nil {
					logger.Global().Error("datasource handle failed", zap.Error(err))

					datasourceError = err

					transactions = append(transactions, model.Transaction{
						AddressFrom: message.Address,
						AddressTo:   message.Address,
						Network:     network,
						Transfers: []model.Transfer{
							{
								Index:       0,
								AddressFrom: message.Address,
								AddressTo:   message.Address,
								Network:     message.Network,
							},
						},
					})

					continue
				}

				transactions = append(transactions, internalTransactions...)
			}
		}
	}

	transactionsMap := getTransactionsMap(transactions)

	if datasourceError != nil {
		if s.handleIndexedWorkers(ctx, message, transactions, transactionsMap) != nil {
			return err
		}

		return datasourceError
	}

	return s.handleWorkers(ctx, message, transactions, transactionsMap)
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

func (s *Server) upsertTransactions(ctx context.Context, transactions []model.Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	if err := s.databaseClient.
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).
		Create(transactions).Error; err != nil {
		return err
	}

	for _, transaction := range transactions {
		if len(transaction.Transfers) == 0 {
			continue
		}

		if err := s.databaseClient.
			Clauses(clause.OnConflict{
				UpdateAll: true,
				DoUpdates: clause.AssignmentColumns([]string{"metadata"}),
			}).
			Create(transaction.Transfers).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) handleWorkers(ctx context.Context, message *protocol.Message, transactions []model.Transaction, transactionsMap map[string]model.Transaction) error {
	// Using workers to clean data
	for _, worker := range s.workers {
		for _, network := range worker.Networks() {
			if network == message.Network {
				internalTransactions, err := worker.Handle(ctx, message, transactions)
				if err != nil {
					logrus.Error(worker.Name(), message.Network, err)

					return err
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

	return s.upsertTransactions(ctx, transactions)
}

func (s *Server) handleIndexedWorkers(ctx context.Context, message *protocol.Message, transactions []model.Transaction, transactionsMap map[string]model.Transaction) error {
	for _, worker := range s.indexedWorker {
		for _, network := range worker.Networks() {
			if network == message.Network {
				internalTransactions, err := worker.Handle(ctx, message, transactions)
				if err != nil {
					logrus.Error(worker.Name(), message.Network, err)

					return err
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

	return s.upsertTransactions(ctx, transactions)
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
