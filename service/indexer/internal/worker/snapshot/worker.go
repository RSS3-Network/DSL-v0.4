package snapshot

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/snapshot/job"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
)

type service struct {
	databaseClient *gorm.DB
	redisClient    *redis.Client
	snapshotClient *snapshot.Client
}

var snapShotWorker = "snapshot_worker"

func (s *service) Name() string {
	return "snapshot"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkEthereumClassic,
		protocol.NetworkBinanceSmartChain,
		protocol.NetworkPolygon,
		protocol.NetworkZkSync,
		protocol.NetworkXDAI,
		protocol.NetworkArbitrum,
		protocol.NetworkOptimism,
		protocol.NetworkFantom,
		protocol.NetworkAvalanche,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	transactions := make([]model.Transaction, 0)

	// filter network

	// get snapshot list

	// set snapshot in Transfers

	return transfers, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&job.SnapshotSpaceJob{
			SnapshotJobBase: job.SnapshotJobBase{
				Name:           "snapshot_space_job",
				DatabaseClient: s.databaseClient,
				RedisClient:    s.redisClient,
				SnapshotClient: s.snapshotClient,
				Limit:          1000,
			},
		},
		&job.SnapshotProposalJob{
			SnapshotJobBase: job.SnapshotJobBase{
				Name:           "snapshot_proposal_job",
				DatabaseClient: s.databaseClient,
				RedisClient:    s.redisClient,
				SnapshotClient: s.snapshotClient,
				Limit:          2000,
			},
		},
		&job.SnapshotVoteJob{
			SnapshotJobBase: job.SnapshotJobBase{
				Name:           "snapshot_vote_job",
				DatabaseClient: s.databaseClient,
				RedisClient:    s.redisClient,
				SnapshotClient: s.snapshotClient,
				Limit:          100000,
			},
		},
	}
}

func (s *service) getSnapshotVoteInfo(ctx context.Context, address string) ([]model.SnapshotVote, error) {
	_, handlerSpan := otel.Tracer(snapShotWorker).Start(ctx, "get_snapshot_vote_info")
	defer handlerSpan.End()

	var snapshotVotes []model.SnapshotVote

	// TODO:from redis

	// from db
	if err := s.databaseClient.
		Model(&model.SnapshotVote{}).
		Where("voter = ?", address).
		Find(&snapshotVotes).Error; err != nil {
		return nil, err
	}

	return snapshotVotes, nil
}

func (s *service) getSnapshotProposalInfo(ctx context.Context, proposals []string) ([]model.SnapshotProposal, error) {
	_, handlerSpan := otel.Tracer(snapShotWorker).Start(ctx, "get_snapshot_proposal_info")
	defer handlerSpan.End()

	var snapshotProposals []model.SnapshotProposal

	// TODO:from redis

	// from db
	if err := s.databaseClient.
		Model(&model.SnapshotProposal{}).
		Where("proposal_id IN (?)", proposals).
		Find(&snapshotProposals).Error; err != nil {
		return nil, err
	}

	return snapshotProposals, nil
}

func (s *service) getSnapshotSpaceInfo(ctx context.Context, spaces []string) ([]model.SnapshotSpace, error) {
	_, handlerSpan := otel.Tracer(snapShotWorker).Start(ctx, "get_snapshot_space_info")
	defer handlerSpan.End()

	var snapshotSpaces []model.SnapshotSpace

	// TODO:from redis

	// from db
	if err := s.databaseClient.
		Model(&model.SnapshotSpace{}).
		Where("space in (?)", spaces).
		Find(&snapshotSpaces).Error; err != nil {
		return nil, err
	}

	return snapshotSpaces, nil
}

// TODO
func (s *service) setSnapshotVoteInfoInCache(ctx context.Context) {}

// TODO
func (s *service) setSnapshotProposalInfoInCache(ctx context.Context) {}

// TODO
func (s *service) setSnapshotSpaceInfoInCache(ctx context.Context) {}

func New(
	databaseClient *gorm.DB,
	redisClient *redis.Client) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		redisClient:    redisClient,
		snapshotClient: snapshot.NewClient(),
	}
}
