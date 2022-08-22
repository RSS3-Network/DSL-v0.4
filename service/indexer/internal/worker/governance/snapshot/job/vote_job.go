package job

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/governance"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/snapshot"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/snapshot/graphql"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"
)

type SnapshotVoteJob struct {
	SnapshotJobBase
}

func (job *SnapshotVoteJob) Name() string {
	return "snapshot_vote_job"
}

func (job *SnapshotVoteJob) Spec() string {
	return "* * * * *" // 1 min
}

func (job *SnapshotVoteJob) Timeout() time.Duration {
	return 10 * time.Minute
}

func (job *SnapshotVoteJob) Run(renewal worker.RenewalFunc) error {
	// nolint:ineffassign // just an initialization
	sleepTime := time.Second

	logrus.Info("[snapshot vote job] run")

	for {
		pullInfoStatus, err := job.InnerJobRun()
		if err != nil {
			logrus.Errorf("[snapshot vote job] run error: %v", err)
		}

		if pullInfoStatus == PullInfoStatusLatest {
			sleepTime = job.LowUpdateTime
		} else {
			sleepTime = job.HighUpdateTime
		}

		if err = renewal(context.Background(), job.Timeout()); err != nil {
			return err
		}

		time.Sleep(sleepTime)
	}
}

func (job *SnapshotVoteJob) InnerJobRun() (status PullInfoStatus, err error) {
	err = job.Check()
	if err != nil {
		return PullInfoStatusNotLatest, fmt.Errorf("[snapshot vote job] check error: %v", err)
	}

	tracer := otel.Tracer("snapshot_vote_job")
	ctx, trace := tracer.Start(context.Background(), "snapshot_vote_job:InnerJobRun")

	defer func() { opentelemetry.Log(trace, nil, status, err) }()

	var statusStroge StatusStroge

	// get latest vote id
	if cache.Global() != nil {
		statusStroge, err = job.GetLastStatusFromCache(ctx)
		if err != nil {
			logrus.Errorf("[snapshot vote job] get last status, db error: %v", err)
			statusStroge.Pos = 0
			statusStroge.Status = PullInfoStatusNotLatest
		}
	}

	if cache.Global() == nil || err != nil {
		statusStroge.Pos, err = job.getVoteTotalFromDB(ctx)
		if err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot vote job] get vote total from db, db error: %v", err)
		}
	}

	// get vote info from url
	skip := statusStroge.Pos

	variable := snapshot.GetMultipleVotesVariable{
		First:          graphql.Int(job.Limit),
		Skip:           graphql.Int(skip),
		OrderBy:        "created",
		OrderDirection: snapshot.OrderDirectionAsc,
	}

	votes, err := job.SnapshotClient.GetMultipleVotes(ctx, variable)
	if err != nil {
		return statusStroge.Status, fmt.Errorf("[snapshot vote job] get multiple votes, graphql error: %v", err)
	}

	// set vote info in db
	if len(votes) > 0 {
		if err := job.setVoteInDB(ctx, votes); err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot vote job] pos[%d], set vote in db, db error: %v", statusStroge.Pos, err)
		}
		logrus.Infof("[snapshot vote job] pull skip [%d]", statusStroge.Pos)
	}

	skip = statusStroge.Pos + job.Limit
	// nolint:gocritic // dont' want change nan things
	if len(votes) == 0 {
		statusStroge.Status = PullInfoStatusLatest
	} else if len(votes) < int(job.Limit) {
		statusStroge.Pos = statusStroge.Pos + int32(len(votes))
		statusStroge.Status = PullInfoStatusLatest
	} else {
		statusStroge.Pos = skip
		statusStroge.Status = PullInfoStatusNotLatest
	}

	// set vote status in cache and db
	if cache.Global() != nil {
		err = job.SetCurrentStatus(ctx, statusStroge)
		if err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot vote job] set current status, db error: %v", err)
		}
	}

	return statusStroge.Status, nil
}

func (job *SnapshotVoteJob) getVoteTotalFromDB(ctx context.Context) (int32, error) {
	var count int64

	if err := database.Global().
		Model(&governance.SnapshotVote{}).
		Select("id").
		Count(&count).Error; err != nil {
		logrus.Errorf("get vote total, db error: %v", err)
		return 0, err
	}

	return int32(count), nil
}

func (job *SnapshotVoteJob) setVoteInDB(ctx context.Context, graphqlVotes []graphqlx.Vote) (err error) {
	tracer := otel.Tracer("snapshot_vote_job")
	_, trace := tracer.Start(ctx, "snapshot_vote_job:setVoteInDB")

	defer func() { opentelemetry.Log(trace, graphqlVotes, nil, err) }()

	votes := []governance.SnapshotVote{}

	for _, graphqlVote := range graphqlVotes {

		vote := governance.SnapshotVote{
			ID:          string(graphqlVote.Id),
			Voter:       strings.ToLower(string(graphqlVote.Voter)),
			Choice:      graphqlVote.Choice,
			ProposalID:  string(graphqlVote.Proposal.Id),
			SpaceID:     string(graphqlVote.Space.Id),
			DateCreated: time.Unix(int64(graphqlVote.Created), 0),
		}

		votes = append(votes, vote)
	}

	if err := database.Global().Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		UpdateAll: true,
	}).Create(&votes).Error; err != nil {
		return err
	}

	return nil
}
