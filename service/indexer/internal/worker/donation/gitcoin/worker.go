package gitcoin

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/donation"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/donation/gitcoin/job"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	Name                   = protocol.PlatfromGitcoin
	ContractAddressPolygon = "0xb99080b9407436ebb2b8fe56d45ffa47e9bb8877"
	ContractAddressEth     = "0x7d655c57f71464b6f83811c55d84009cd9f5221c"
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
	return Name
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

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (internalTransactions []model.Transaction, err error) {
	tracer := otel.Tracer("gitcoin_worker")
	_, trace := tracer.Start(ctx, "gitcoin_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, internalTransactions, err) }()

	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		if transaction.AddressTo != ContractAddressPolygon && transaction.AddressTo != ContractAddressEth {
			internalTransactionMap[transaction.Hash] = transaction

			continue
		}

		for _, transfer := range transaction.Transfers {
			transfer, err = s.handleGitcoin(ctx, transfer)
			if err != nil {
				logrus.Errorf("[gitcoin worker] Handle: handleGitcoin error, %v", err)
			}

			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Tag, value.Type = filter.UpdateTagAndType(transfer.Tag, value.Tag, transfer.Type, value.Type)
			value.Transfers = append(value.Transfers, transfer)

			if value.Tag == filter.TagDonation {
				value.Platform = Name
			}

			internalTransactionMap[transaction.Hash] = value
		}
	}

	internalTransactions = make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, internalTransaction)
	}

	return internalTransactions, nil
}

func (s *service) handleGitcoin(ctx context.Context, transfer model.Transfer) (data model.Transfer, err error) {
	tracer := otel.Tracer("gitcoin_worker")
	_, trace := tracer.Start(ctx, "gitcoin_worker:handleGitcoin")

	defer func() { opentelemetry.Log(trace, transfer, data, err) }()

	if transfer.AddressTo == "" ||
		transfer.AddressTo == "0x0" ||
		transfer.AddressTo == "0x0000000000000000000000000000000000000000" {
		return transfer, err
	}

	projectStr, err := s.redisClient.HGet(ctx, s.gitcoinProjectCacheKey, transfer.AddressTo).Result()
	if err != nil || len(projectStr) == 0 {
		return transfer, err
	}

	project := &donation.GitcoinProject{}
	if err := json.Unmarshal([]byte(projectStr), &project); err != nil {
		logrus.Errorf("[gitcoin handle] json unmarshal project error: %v", err)
		return transfer, err
	}

	transfer.RelatedUrls = append(transfer.RelatedUrls, "https://gitcoin.co/grants"+strconv.Itoa(project.ID)+"/"+project.Slug)

	// format metadata
	var metadataModel metadata.Metadata

	if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
		logrus.Errorf("[gitcoin handle] json unmarshal transfer metadata error: %v", err)
		return transfer, err
	}

	metadataModel.Gitcoin = &metadata.Gitcoin{
		ID:          project.ID,
		Slug:        project.Slug,
		Title:       project.Title,
		Description: project.Description,
		Logo:        project.Logo,
	}

	metadata, err := json.Marshal(metadataModel)
	if err != nil {
		return transfer, err
	}

	transfer.Metadata = metadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagDonation, transfer.Tag, filter.DonationDonate, transfer.Type)

	if transfer.Tag == filter.TagDonation {
		transfer.Platform = Name
	}

	return transfer, nil
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
