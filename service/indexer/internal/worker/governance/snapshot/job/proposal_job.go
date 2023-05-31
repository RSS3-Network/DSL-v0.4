package job

import (
	"context"
	"encoding/json"
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

type SnapshotProposalJob struct {
	SnapshotJobBase
}

func (job *SnapshotProposalJob) Name() string {
	return "snapshot_proposal_job"
}

func (job *SnapshotProposalJob) Spec() string {
	return "* * * * *" // 1 min
}

func (job *SnapshotProposalJob) Timeout() time.Duration {
	return time.Minute * 2
}

func (job *SnapshotProposalJob) Run(renewal worker.RenewalFunc) error {
	// nolint:ineffassign // just an initialization
	sleepTime := time.Second

	logrus.Info("[snapshot proposal job] run")

	for {
		pullInfoStatus, err := job.InnerJobRun()
		if err != nil {
			logrus.Errorf("[snapshot proposal job] run error: %v", err)
		}

		if pullInfoStatus == PullInfoStatusLatest {
			sleepTime = job.LowUpdateTime
		} else {
			sleepTime = job.HighUpdateTime
		}

		if err = renewal(context.Background(), time.Minute*2); err != nil {
			return err
		}

		time.Sleep(sleepTime)
	}
}

func (job *SnapshotProposalJob) InnerJobRun() (status PullInfoStatus, err error) {
	err = job.Check()
	if err != nil {
		return PullInfoStatusNotLatest, fmt.Errorf("[snapshot proposal job] check error: %v", err)
	}

	tracer := otel.Tracer("snapshot_proposal_job")
	ctx, trace := tracer.Start(context.Background(), "snapshot_proposal_job:InnerJobRun")

	defer func() { opentelemetry.Log(trace, nil, status, err) }()

	created, err := job.getLatestProposalFromDB(ctx)
	if err != nil {
		status = PullInfoStatusNotLatest
		return status, fmt.Errorf("[snapshot proposal job] get latest proposal total from db, db error: %v", err)
	}

	// get proposal info from url
	variable := snapshot.GetMultipleProposalsVariable{
		First: graphql.Int(job.Limit),
		Where: graphqlx.ProposalWhere{
			CreatedGte: graphql.Int(created),
		},
		OrderBy:        "created",
		OrderDirection: snapshot.OrderDirectionAsc,
	}

	proposals, err := job.SnapshotClient.GetMultipleProposals(ctx, variable)
	if err != nil {
		status = PullInfoStatusNotLatest
		return status, fmt.Errorf("[snapshot proposal job] get multiple proposals, graphql error: %v", err)
	}

	if len(proposals) > 0 {
		if err := job.setProposalsInDB(ctx, proposals); err != nil {
			status = PullInfoStatusNotLatest
			return status, fmt.Errorf("[snapshot proposal job] created[%d], set proposal in db, db error: %v", created, err)
		}
		logrus.Infof("[snapshot proposal job] pull created [%d]", created)

		// upsert spaces
		spaceIdMap := make(map[graphql.String]struct{}, 0)
		spaceIds := make([]graphql.String, 0)

		for _, proposal := range proposals {
			if _, ok := spaceIdMap[proposal.Space.Id]; !ok {
				spaceIds = append(spaceIds, proposal.Space.Id)
			}
			spaceIdMap[proposal.Space.Id] = struct{}{}
		}

		if len(spaceIds) > 0 {
			err = job.upsertSpacesInDB(ctx, spaceIds)
			if err != nil {
				status = PullInfoStatusNotLatest
				return status, fmt.Errorf("[snapshot proposal job] get multiple spaces, graphql error: %v", err)
			}
		}
	}

	// nolint:gocritic // dont' want change nan things
	if len(proposals) == 0 {
		status = PullInfoStatusLatest
	} else if len(proposals) < int(job.Limit) {
		status = PullInfoStatusLatest
	} else {
		status = PullInfoStatusNotLatest
	}

	return status, nil
}

func (job *SnapshotProposalJob) getLatestProposalFromDB(ctx context.Context) (int32, error) {
	var proposal governance.SnapshotProposal

	if err := database.Global().
		Model(&governance.SnapshotProposal{}).
		Order("date_created DESC").
		Limit(1).
		Take(&proposal).Error; err != nil {
		logrus.Errorf("get latest proposal, db error: %v", err)
		return 0, err
	}

	return int32(proposal.DateCreated.Unix()), nil
}

func (job *SnapshotProposalJob) setProposalsInDB(ctx context.Context, graphqlProposals []graphqlx.Proposal) (err error) {
	tracer := otel.Tracer("snapshot_proposal_job")
	_, trace := tracer.Start(ctx, "snapshot_proposal_job:setProposalsInDB")

	defer func() { opentelemetry.Log(trace, graphqlProposals, nil, err) }()

	var proposals []governance.SnapshotProposal

	for _, graphqlProposal := range graphqlProposals {
		metadata, err := json.Marshal(graphqlProposal)
		if err != nil {
			logrus.Warnf("[snapshot proposal job] marshal proposal metadata, error: %v", err)
			continue
		}

		proposal := governance.SnapshotProposal{
			ID:          string(graphqlProposal.Id),
			SpaceID:     string(graphqlProposal.Space.Id),
			Author:      strings.ToLower(string(graphqlProposal.Author)),
			DateCreated: time.Unix(int64(graphqlProposal.Created), 0),
			Metadata:    metadata,
		}

		proposals = append(proposals, proposal)
	}

	if err := database.Global().Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		UpdateAll: true,
	}).Create(&proposals).Error; err != nil {
		return err
	}

	return nil
}
