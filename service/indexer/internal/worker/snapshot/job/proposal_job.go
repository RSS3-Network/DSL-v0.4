package job

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	graphqlx "github.com/naturalselectionlabs/pregod/common/snapshot/graphql"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"
)

type SnapshotProposalJob struct {
	SnapshotJobBase
}

var traceProposalJob = "proposal_job"

func (job *SnapshotProposalJob) Name() string {
	return "snapshot_proposal_job"
}

func (job *SnapshotProposalJob) Spec() string {
	return "* * * * *" // 1 min
}

func (job *SnapshotProposalJob) Timeout() time.Duration {
	return time.Minute
}

func (job *SnapshotProposalJob) Run(renewal worker.RenewalFunc) error {
	// nolint:ineffassign // just an initialization
	sleepTime := time.Second

	for {
		pullInfoStatus, err := job.InnerJobRun()
		if err != nil {
			logrus.Errorf("[snapshot proposal job] run error: %v", err)
		}

		if pullInfoStatus == PullInfoStatusLatest {
			sleepTime = time.Second
		} else {
			sleepTime = time.Minute * 5
		}

		time.Sleep(sleepTime)
	}
}

func (job *SnapshotProposalJob) InnerJobRun() (PullInfoStatus, error) {
	err := job.Check()
	if err != nil {
		return PullInfoStatusNotLatest, fmt.Errorf("[snapshot proposal job] check error: %v", err)
	}

	ctx, runSnap := otel.Tracer(traceProposalJob).Start(context.Background(), "run")
	defer runSnap.End()

	logrus.Info("[snapshot proposal job] run")

	// get latest proposal id
	statusStroge, err := job.GetLastStatusFromCache(ctx)
	if err != nil {
		logrus.Errorf("[snapshot proposal job] get last status, db error: %v", err)
		statusStroge.Pos = 0
		statusStroge.Status = PullInfoStatusNotLatest
	}

	if err != nil || statusStroge.Pos == 0 {
		statusStroge.Pos, err = job.getProposalTotalFromDB(ctx)
		if err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot proposal job] get proposal total from db, db error: %v", err)
		}
	}

	// get proposal info from url
	skip := statusStroge.Pos + job.Limit
	variable := snapshot.GetMultipleProposalsVariable{
		First:          graphql.Int(job.Limit),
		Skip:           graphql.Int(skip),
		OrderBy:        "created",
		OrderDirection: snapshot.OrderDirectionAsc,
	}

	proposals, err := job.SnapshotClient.GetMultipleProposals(ctx, variable)
	if err != nil {
		return statusStroge.Status, fmt.Errorf("[snapshot proposal job] get multiple proposals, graphql error: %v", err)
	}

	setInDb := false
	// nolint:gocritic // dont' want change nan things
	if len(proposals) == 0 {
		statusStroge.Status = PullInfoStatusLatest
	} else if len(proposals) < int(job.Limit) {
		setInDb = true
		statusStroge.Pos = statusStroge.Pos + int32(len(proposals))
		statusStroge.Status = PullInfoStatusLatest
	} else {
		setInDb = true
		statusStroge.Pos = skip
		statusStroge.Status = PullInfoStatusNotLatest
	}

	// set proposal info in db
	if setInDb {
		if err := job.setProposalsInDB(ctx, proposals); err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot proposal job] pos[%d], set proposal in db, db error: %v", statusStroge.Pos, err)
		}
	}

	return PullInfoStatusLatest, nil
}

func (job *SnapshotProposalJob) getProposalTotalFromDB(ctx context.Context) (int32, error) {
	_, handlerSpan := otel.Tracer(traceProposalJob).Start(ctx, "get_proposal_total_from_db")
	defer handlerSpan.End()

	var count int64

	if err := job.DatabaseClient.
		Model(&model.SnapshotProposal{}).
		Select("id").
		Count(&count).Error; err != nil {
		logrus.Errorf("get proposal total, db error: %v", err)
		return 0, err
	}

	return int32(count), nil
}

func (job *SnapshotProposalJob) setProposalsInDB(ctx context.Context, graphqlproposals []graphqlx.Proposal) error {
	_, handlerSpan := otel.Tracer(traceProposalJob).Start(ctx, "set_proposal_in_db")
	defer handlerSpan.End()

	proposals := []model.SnapshotProposal{}

	for _, graphqlproposal := range graphqlproposals {
		metadata, err := json.Marshal(graphqlproposal)
		if err != nil {
			logrus.Warnf("[snapshot proposal job] marshal proposal metadata, error: %v", err)
			continue
		}

		proposal := model.SnapshotProposal{
			ID:       string(graphqlproposal.Id),
			SpaceID:  string(graphqlproposal.Id),
			Metadata: metadata,
		}

		proposals = append(proposals, proposal)
	}

	if err := job.DatabaseClient.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		UpdateAll: true,
	}).Create(&proposals).Error; err != nil {
		return err
	}

	return nil
}
