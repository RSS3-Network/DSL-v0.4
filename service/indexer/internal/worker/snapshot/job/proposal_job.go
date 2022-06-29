package job

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"

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

func (job *SnapshotProposalJob) InnerJobRun() (PullInfoStatus, error) {
	err := job.Check()
	if err != nil {
		return PullInfoStatusNotLatest, fmt.Errorf("[snapshot proposal job] check error: %v", err)
	}

	tracer := otel.Tracer("snapshot_proposal_job")
	ctx, trace := tracer.Start(context.Background(), "snapshot_proposal_job:InnerJobRun")

	defer trace.End()

	var statusStroge StatusStroge

	// get latest proposal id
	if job.RedisClient != nil {
		statusStroge, err = job.GetLastStatusFromCache(ctx)
		if err != nil {
			logrus.Errorf("[snapshot proposal job] get last status, db error: %v", err)
			statusStroge.Pos = 0
			statusStroge.Status = PullInfoStatusNotLatest
		}
	}

	if job.RedisClient == nil || err != nil {
		statusStroge.Pos, err = job.getProposalTotalFromDB(ctx)
		if err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot proposal job] get proposal total from db, db error: %v", err)
		}
	}

	// get proposal info from url
	skip := statusStroge.Pos
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

	if len(proposals) > 0 {
		if err := job.setProposalsInDB(ctx, proposals); err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot proposal job] pos[%d], set proposal in db, db error: %v", statusStroge.Pos, err)
		}
		logrus.Infof("[snapshot proposal job] pull skip [%d]", statusStroge.Pos)
	}

	skip = statusStroge.Pos + job.Limit
	// nolint:gocritic // dont' want change nan things
	if len(proposals) == 0 {
		statusStroge.Status = PullInfoStatusLatest
	} else if len(proposals) < int(job.Limit) {
		statusStroge.Pos = statusStroge.Pos + int32(len(proposals))
		statusStroge.Status = PullInfoStatusLatest
	} else {
		statusStroge.Pos = skip
		statusStroge.Status = PullInfoStatusNotLatest
	}

	// set space status in cache and db
	if job.RedisClient != nil {
		err = job.SetCurrentStatus(ctx, statusStroge)
		if err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot proposal job] set current status, db error: %v", err)
		}
	}

	return statusStroge.Status, nil
}

func (job *SnapshotProposalJob) getProposalTotalFromDB(ctx context.Context) (int32, error) {
	tracer := otel.Tracer("snapshot_proposal_job")
	_, trace := tracer.Start(ctx, "snapshot_proposal_job:getProposalTotalFromDB")

	defer trace.End()

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
	tracer := otel.Tracer("snapshot_proposal_job")
	_, trace := tracer.Start(ctx, "snapshot_proposal_job:setProposalsInDB")

	defer trace.End()

	proposals := []model.SnapshotProposal{}

	for _, graphqlproposal := range graphqlproposals {
		metadata, err := json.Marshal(graphqlproposal)
		if err != nil {
			logrus.Warnf("[snapshot proposal job] marshal proposal metadata, error: %v", err)
			continue
		}

		proposal := model.SnapshotProposal{
			ID:          string(graphqlproposal.Id),
			SpaceID:     string(graphqlproposal.Space.Id),
			Author:      strings.ToLower(string(graphqlproposal.Author)),
			DateCreated: time.Unix(int64(graphqlproposal.Created), 0),
			Metadata:    metadata,
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
