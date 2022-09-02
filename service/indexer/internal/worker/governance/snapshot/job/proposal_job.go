package job

import (
	"context"
	"encoding/json"
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
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/internalModel"
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

func (job *SnapshotProposalJob) Run(renewal internalModel.RenewalFunc) error {
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

	var statusStorage StatusStorage

	// get latest proposal id
	if cache.Global() != nil {
		statusStorage, err = job.GetLastStatusFromCache(ctx)
		if err != nil {
			logrus.Errorf("[snapshot proposal job] get last status, db error: %v", err)
			statusStorage.Pos = 0
			statusStorage.Status = PullInfoStatusNotLatest
		}
	}

	if cache.Global() == nil || err != nil {
		statusStorage.Pos, err = job.getProposalTotalFromDB(ctx)
		if err != nil {
			return statusStorage.Status, fmt.Errorf("[snapshot proposal job] get proposal total from db, db error: %v", err)
		}
	}

	// get proposal info from url
	skip := statusStorage.Pos
	variable := snapshot.GetMultipleProposalsVariable{
		First:          graphql.Int(job.Limit),
		Skip:           graphql.Int(skip),
		OrderBy:        "created",
		OrderDirection: snapshot.OrderDirectionAsc,
	}

	proposals, err := job.SnapshotClient.GetMultipleProposals(ctx, variable)
	if err != nil {
		return statusStorage.Status, fmt.Errorf("[snapshot proposal job] get multiple proposals, graphql error: %v", err)
	}

	if len(proposals) > 0 {
		if err := job.setProposalsInDB(ctx, proposals); err != nil {
			return statusStorage.Status, fmt.Errorf("[snapshot proposal job] pos[%d], set proposal in db, db error: %v", statusStorage.Pos, err)
		}
		logrus.Infof("[snapshot proposal job] pull skip [%d]", statusStorage.Pos)
	}

	skip = statusStorage.Pos + job.Limit
	// nolint:gocritic // dont' want change nan things
	if len(proposals) == 0 {
		statusStorage.Status = PullInfoStatusLatest
	} else if len(proposals) < int(job.Limit) {
		statusStorage.Pos = statusStorage.Pos + int32(len(proposals))
		statusStorage.Status = PullInfoStatusLatest
	} else {
		statusStorage.Pos = skip
		statusStorage.Status = PullInfoStatusNotLatest
	}

	// set space status in cache and db
	if cache.Global() != nil {
		err = job.SetCurrentStatus(ctx, statusStorage)
		if err != nil {
			return statusStorage.Status, fmt.Errorf("[snapshot proposal job] set current status, db error: %v", err)
		}
	}

	return statusStorage.Status, nil
}

func (job *SnapshotProposalJob) getProposalTotalFromDB(ctx context.Context) (int32, error) {
	var count int64

	if err := database.Global().
		Model(&governance.SnapshotProposal{}).
		Select("id").
		Count(&count).Error; err != nil {
		logrus.Errorf("get proposal total, db error: %v", err)
		return 0, err
	}

	return int32(count), nil
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
