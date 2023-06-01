package job

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hasura/go-graphql-client"
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

	created, err := job.getLatestVoteFromDB(ctx)
	if err != nil {
		status = PullInfoStatusNotLatest
		return status, fmt.Errorf("[snapshot vote job] get latest vote total from db, db error: %v", err)
	}

	variable := snapshot.GetMultipleVotesVariable{
		First: graphql.Int(job.Limit),
		Where: graphqlx.VoteWhere{
			CreatedGte: graphql.Int(created),
		},
		OrderBy:        "created",
		OrderDirection: snapshot.OrderDirectionAsc,
	}

	votes, err := job.SnapshotClient.GetMultipleVotes(ctx, variable)
	if err != nil {
		status = PullInfoStatusNotLatest
		return status, fmt.Errorf("[snapshot vote job] get multiple votes, graphql error: %v", err)
	}

	// set vote info in db
	if len(votes) > 0 {
		if err := job.setVoteInDB(ctx, votes); err != nil {
			status = PullInfoStatusNotLatest
			return status, fmt.Errorf("[snapshot vote job] created[%d], set vote in db, db error: %v", created, err)
		}
		logrus.Infof("[snapshot vote job] pull created [%d]", created)

		// upsert spaces
		spaceIdMap := make(map[graphql.String]struct{}, 0)
		spaceIds := make([]graphql.String, 0)

		for _, vote := range votes {
			if _, ok := spaceIdMap[vote.Space.Id]; !ok {
				spaceIds = append(spaceIds, vote.Space.Id)
			}
			spaceIdMap[vote.Space.Id] = struct{}{}
		}

		if len(spaceIds) > 0 {
			err = job.upsertSpacesInDB(ctx, spaceIds)
			if err != nil {
				status = PullInfoStatusNotLatest
				return status, fmt.Errorf("[snapshot vote job] get multiple spaces, graphql error: %v", err)
			}
		}
	}

	// nolint:gocritic // dont' want change nan things
	if len(votes) == 0 {
		status = PullInfoStatusLatest
	} else if len(votes) < int(job.Limit) {
		status = PullInfoStatusLatest
	} else {
		status = PullInfoStatusNotLatest
	}

	return status, nil
}

func (job *SnapshotVoteJob) getLatestVoteFromDB(ctx context.Context) (int32, error) {
	var vote governance.SnapshotVote

	if err := database.Global().
		Model(&governance.SnapshotVote{}).
		Order("date_created DESC").
		Limit(1).
		Take(&vote).Error; err != nil {
		logrus.Errorf("get latest vote, db error: %v", err)
		return 0, err
	}

	return int32(vote.DateCreated.Unix()), nil
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
