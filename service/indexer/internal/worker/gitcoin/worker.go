package gitcoin

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/naturalselectionlabs/pregod/common/constant"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/gitcoin/job"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

var _ worker.Worker = (*service)(nil)

type service struct {
	databaseClient         *gorm.DB
	redisClient            *redis.Client
	gitcoinProjectCacheKey string
}

func New(databaseClient *gorm.DB, redisClient *redis.Client) worker.Worker {
	return &service{
		databaseClient:         databaseClient,
		redisClient:            redisClient,
		gitcoinProjectCacheKey: "gitcoin_project",
	}
}

func (s *service) Name() string {
	return "gitcoin"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
		protocol.NetworkZkSync,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_gitcoin")

	_, handlerSpan := tracer.Start(ctx, "handler")

	defer handlerSpan.End()

	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			projectStr, err := s.redisClient.HGet(ctx, s.gitcoinProjectCacheKey, transfer.AddressTo).Result()
			if err != nil || len(projectStr) == 0 {
				continue
			}

			project := &model.GitcoinProject{}
			if err := json.Unmarshal([]byte(projectStr), &project); err != nil {
				logrus.Errorf("[gitcoin handle] json unmarshal project error: %v", err)
				continue
			}

			transfer.Type = "donate"
			transfer.Tags = append(transfer.Tags, constant.TransferTagGitcoin.String(), constant.TransferTagDonation.String())
			transfer.RelatedUrls = append(transfer.RelatedUrls, utils.GetTxHashURL(transfer.Network, transfer.TransactionHash), "https://gitcoin.co/grants"+strconv.Itoa(project.Id)+"/"+project.Slug)

			// format metadata
			var metadataModel metadata.Metadata

			if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
				logrus.Errorf("[gitcoin handle] json unmarshal transfer metadata error: %v", err)
				continue
			}

			metadataModel.Gitcoin = &metadata.Gitcoin{
				Id:          project.Id,
				Slug:        project.Slug,
				Title:       project.Title,
				Description: project.Description,
				Logo:        project.Logo,
			}

			metadata, err := json.Marshal(metadataModel)
			if err != nil {
				continue
			}

			transfer.Metadata = metadata

			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Transfers = append(value.Transfers, transfer)
			internalTransactionMap[transaction.Hash] = value
		}
	}

	internalTransactions := make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, internalTransaction)
	}

	return internalTransactions, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&job.GitcoinProjectJob{
			DatabaseClient:         s.databaseClient,
			RedisClient:            s.redisClient,
			GitcoinProjectCacheKey: s.gitcoinProjectCacheKey,
		},
	}
}
