package server

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/nft"
	"github.com/naturalselectionlabs/pregod/common/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/lens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/snapshot"

	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/gitcoin"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/mirror"
	poapworker "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/poap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/swap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token/coinmarketcap"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
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
			semconv.ServiceNameKey.String("pregod-indexer"),
			semconv.ServiceVersionKey.String("v0.5.0"),
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
		"", false, false, false, false, nil,
	); err != nil {
		return err
	}

	if err := s.rabbitmqChannel.QueueBind(
		s.rabbitmqQueue.Name, "", protocol.ExchangeJob, false, nil,
	); err != nil {
		return err
	}

	s.datasources = []datasource.Datasource{
		moralis.New(s.config.Moralis.Key), arweave.New(), blockscout.New(), zksync.New(), lens.New(s.databaseClient),
	}

	s.workers = []worker.Worker{
		token.New(s.databaseClient),
		swap.New(s.employer, s.databaseClient),
		mirror.New(),
		poapworker.New(),
		gitcoin.New(s.databaseClient, s.redisClient),
		snapshot.New(s.databaseClient, s.redisClient),
		crossbell.New(),
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

	deliveryCh, err := s.rabbitmqChannel.Consume(s.rabbitmqQueue.Name, "", false, false, false, false, nil)
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
				logrus.Errorln(err)
			}
		}()
	}

	return nil
}

func (s *Server) handle(ctx context.Context, message *protocol.Message) (err error) {
	logrus.Info(message.Address, message.Network)

	// Get data from datasources
	var transactions []model.Transaction

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
				if err != nil {
					return err
				}

				transactions = append(transactions, internalTransactions...)
			}
		}
	}

	if err := s.upsertTransactions(ctx, transactions); err != nil {
		return err
	}

	// Using workers to clean data
	for _, worker := range s.workers {
		for _, network := range worker.Networks() {
			if network == message.Network {
				internalTransactions, err := worker.Handle(ctx, message, transactions)
				if err != nil {
					return err
				}

				if err := s.upsertTransactions(ctx, internalTransactions); err != nil {
					return err
				}
				// if no replacement here, transfers may be edited by more than one workers
				// but previous edits will be lost
				transactions = replaceTransactions(transactions, internalTransactions)
			}
		}
	}

	return nil
}

// if same hash exists in newTransactions, replace it
func replaceTransactions(oldTransactions, newTransactions []model.Transaction) []model.Transaction {
	newMap := map[string]model.Transaction{}
	for _, t := range newTransactions {
		newMap[t.Hash] = t
	}

	resultMap := map[string]model.Transaction{}
	for _, oldT := range oldTransactions {
		if newT, exists := newMap[oldT.Hash]; exists {
			resultMap[oldT.Hash] = newT
		} else {
			resultMap[oldT.Hash] = oldT
		}
	}

	result := []model.Transaction{}
	for _, t := range resultMap {
		result = append(result, t)
	}
	return result
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
			return nil
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

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
