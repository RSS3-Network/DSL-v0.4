package snapshot

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/snapshot/job"
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
	//internalTransfers := make([]model.Transfer, 0)

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

// TODO:add cache for snapshot

func getSnapshotVoteInfo() {
	// from db
	
}

func getSnapshotSpaceInfo() {

}

func getSnapshotProposalInfo() {

}

func New(
	databaseClient *gorm.DB,
	redisClient *redis.Client) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		redisClient:    redisClient,
		snapshotClient: snapshot.NewClient(),
	}
}
