package job

import (
	"context"
	"fmt"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	graphqlx "github.com/naturalselectionlabs/pregod/common/snapshot/graphql"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"
)

type SnapshotVoteJob struct {
	SnapshotJobBase
}

var traceVoteJob = "vote_job"

func (job *SnapshotVoteJob) Name() string {
	return "snapshot_vote_job"
}

func (job *SnapshotVoteJob) Spec() string {
	return "* * * * *" // 1 min
}

func (job *SnapshotVoteJob) Timeout() time.Duration {
	return time.Minute
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

		if err = renewal(context.Background(), time.Minute); err != nil {
			return err
		}

		time.Sleep(sleepTime)
	}
}

func (job *SnapshotVoteJob) InnerJobRun() (PullInfoStatus, error) {
	err := job.Check()
	if err != nil {
		return PullInfoStatusNotLatest, fmt.Errorf("[snapshot vote job] check error: %v", err)
	}

	ctx, runSnap := otel.Tracer(traceVoteJob).Start(context.Background(), "run")
	defer runSnap.End()

	// get latest vote id
	statusStroge, err := job.GetLastStatusFromCache(ctx)
	if err != nil {
		logrus.Errorf("[snapshot vote job] get last status, db error: %v", err)
		statusStroge.Pos = 0
		statusStroge.Status = PullInfoStatusNotLatest
	}

	if err != nil {
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
	err = job.SetCurrentStatus(ctx, statusStroge)
	if err != nil {
		return statusStroge.Status, fmt.Errorf("[snapshot vote job] set current status, db error: %v", err)
	}

	return statusStroge.Status, nil
}

func (job *SnapshotVoteJob) getVoteTotalFromDB(ctx context.Context) (int32, error) {
	_, handlerSpan := otel.Tracer(traceVoteJob).Start(ctx, "get_vote_total_from_db")
	defer handlerSpan.End()

	var count int64

	if err := job.DatabaseClient.
		Model(&model.SnapshotVote{}).
		Select("id").
		Count(&count).Error; err != nil {
		logrus.Errorf("get vote total, db error: %v", err)
		return 0, err
	}

	return int32(count), nil
}

func (job *SnapshotVoteJob) setVoteInDB(ctx context.Context, graphqlVotes []graphqlx.Vote) error {
	_, handlerSpan := otel.Tracer(traceVoteJob).Start(ctx, "set_vote_in_db")
	defer handlerSpan.End()

	votes := []model.SnapshotVote{}

	for _, graphqlVote := range graphqlVotes {

		vote := model.SnapshotVote{
			ID:          string(graphqlVote.Id),
			Voter:       string(graphqlVote.Voter),
			Choice:      graphqlVote.Choice,
			ProposalID:  string(graphqlVote.Proposal.Id),
			SpaceID:     string(graphqlVote.Space.Id),
			DateCreated: time.Unix(int64(graphqlVote.Created), 0),
		}

		votes = append(votes, vote)
	}

	if err := job.DatabaseClient.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		UpdateAll: true,
	}).Create(&votes).Error; err != nil {
		return err
	}

	return nil
}
