package job_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/cache"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/snapshot/job"
	"gorm.io/gorm"
	"testing"
)

var (
	databaseClient *gorm.DB
	redisClient    *redis.Client
	snapshotClient *snapshot.Client
)

func init() {
	var err error

	databaseClient, err = database.Dial("postgres://postgres:password@127.0.0.1:5432/pregod2", true)
	if err != nil {
		panic(err)
	}

	redisClient, err = cache.Dial(&configx.Redis{
		Addr:       "127.0.0.1:6379",
		Password:   "",
		DB:         0,
		TLSEnabled: false,
	})
	if err != nil {
		panic(err)
	}

	snapshotClient = snapshot.NewClient()
}

func TestSnapshotSpaceInnerJob(t *testing.T) {
	count := 0

	for {
		if count > 2 {
			break
		}

		spaceJob := job.SnapshotSpaceJob{
			SnapshotJobBase: job.SnapshotJobBase{
				Name:           "snapshot_space_job",
				DatabaseClient: databaseClient,
				RedisClient:    redisClient,
				SnapshotClient: snapshotClient,
				Limit:          100,
			},
		}

		if _, err := spaceJob.InnerJobRun(); err != nil {
			panic(err)
		}

		count = count + 1
	}
}

func TestSnapshotProposalInnerJob(t *testing.T) {
	count := 0

	for {
		if count > 2 {
			break
		}

		proposalJob := job.SnapshotProposalJob{
			SnapshotJobBase: job.SnapshotJobBase{
				Name:           "snapshot_proposal_job",
				DatabaseClient: databaseClient,
				RedisClient:    redisClient,
				SnapshotClient: snapshotClient,
				Limit:          100,
			},
		}

		if _, err := proposalJob.InnerJobRun(); err != nil {
			panic(err)
		}

		count = count + 1
	}
}

func TestSnapshotVoteInnerJob(t *testing.T) {
	count := 0

	for {
		if count > 2 {
			break
		}
		voteJob := job.SnapshotVoteJob{
			SnapshotJobBase: job.SnapshotJobBase{
				Name:           "snapshot_vote_job",
				DatabaseClient: databaseClient,
				RedisClient:    redisClient,
				SnapshotClient: snapshotClient,
				Limit:          100,
			},
		}

		if _, err := voteJob.InnerJobRun(); err != nil {
			t.Error(err)
		}

		count = count + 1
	}
}
