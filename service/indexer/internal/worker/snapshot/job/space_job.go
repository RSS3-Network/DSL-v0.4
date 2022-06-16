package job

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	graphqlx "github.com/naturalselectionlabs/pregod/common/snapshot/graphql"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"

	"github.com/sirupsen/logrus"
)

type SnapshotSpaceJob struct {
	SnapshotJobBase
}

var TraceSpaceJob = "space_job"

const LimitOnce = 1000

func (job *SnapshotSpaceJob) Name() string {
	return "snapshot_space_job"
}

func (job *SnapshotSpaceJob) Spec() string {
	return "* * * * *" // 1 min
}

func (job *SnapshotSpaceJob) Timeout() time.Duration {
	return time.Minute
}

func (job *SnapshotSpaceJob) Run(renewal worker.RenewalFunc) error {
	// nolint:ineffassign // just an initialization
	sleepTime := time.Second

	for {
		pullInfoStatus, err := job.InnerJobRun()
		if err != nil {
			logrus.Errorf("[snapshot space job] run error: %v", err)
		}

		if pullInfoStatus == PullInfoStatusLatest {
			sleepTime = time.Second
		} else {
			sleepTime = time.Minute * 5
		}

		if err = renewal(context.Background(), time.Minute); err != nil {
			return err
		}

		time.Sleep(sleepTime)
	}
}

func (job *SnapshotSpaceJob) InnerJobRun() (PullInfoStatus, error) {
	err := job.Check()
	if err != nil {
		return PullInfoStatusNotLatest, fmt.Errorf("[snapshot space job] check error: %v", err)
	}

	ctx, runSnap := otel.Tracer(TraceSpaceJob).Start(context.Background(), "run")
	defer runSnap.End()

	logrus.Info("[snapshot space job] run")

	// get latest space id
	statusStroge, err := job.GetLastStatusFromCache(ctx)
	if err != nil {
		logrus.Errorf("[snapshot space job] get last status, db error: %v", err)
		statusStroge.Pos = 0
		statusStroge.Status = PullInfoStatusNotLatest
	}

	if err != nil || statusStroge.Pos == 0 {
		statusStroge.Pos, err = job.getSpaceTotalFromDB(ctx)
		if err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot space job] get space total from db, db error: %v", err)
		}
	}

	// get space info from url
	skip := statusStroge.Pos + LimitOnce
	variable := snapshot.GetMultipleSpacesVariable{
		First:          LimitOnce,
		Skip:           graphql.Int(skip),
		OrderBy:        "created",
		OrderDirection: snapshot.OrderDirectionAsc,
	}

	spaces, err := job.SnapshotClient.GetMultipleSpaces(ctx, variable)
	if err != nil {
		return statusStroge.Status, fmt.Errorf("[snapshot space job] get multiple spaces, graphql error: %v", err)
	}

	setInDb := false
	// nolint:gocritic // dont' want change nan things
	if len(spaces) == 0 {
		statusStroge.Status = PullInfoStatusLatest
	} else if len(spaces) < LimitOnce {
		setInDb = true
		statusStroge.Pos = statusStroge.Pos + int32(len(spaces))
		statusStroge.Status = PullInfoStatusLatest
	} else {
		setInDb = true
		statusStroge.Pos = skip
		statusStroge.Status = PullInfoStatusNotLatest
	}

	// set space info in db
	if setInDb {
		if err := job.setSpaceInDB(ctx, spaces); err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot space job] pos[%d], set space in db, db error: %v", statusStroge.Pos, err)
		}
	}

	// set space status in cache and db
	err = job.SetCurrentStatus(ctx, statusStroge)
	if err != nil {
		return statusStroge.Status, fmt.Errorf("[snapshot space job] set current status, db error: %v", err)
	}

	return statusStroge.Status, nil
}

func (job *SnapshotSpaceJob) getSpaceTotalFromDB(ctx context.Context) (int32, error) {
	_, handlerSpan := otel.Tracer(TraceSpaceJob).Start(ctx, "get_space_total_from_db")
	defer handlerSpan.End()

	var count int64

	if err := job.DatabaseClient.
		Model(&model.SnapshotSpace{}).
		Select("id").
		Count(&count).Error; err != nil {
		logrus.Errorf("get space total, db error: %v", err)
		return 0, err
	}

	return int32(count), nil
}

func (job *SnapshotSpaceJob) setSpaceInDB(ctx context.Context, graphqlSpaces []graphqlx.Space) error {
	_, handlerSpan := otel.Tracer(TraceSpaceJob).Start(ctx, "set_space_in_db")
	defer handlerSpan.End()

	spaces := []model.SnapshotSpace{}

	for _, graphqlSpace := range graphqlSpaces {
		metadata, err := json.Marshal(graphqlSpace)
		if err != nil {
			logrus.Warnf("[snapshot space job] marshal space metadata, error: %v", err)
			continue
		}

		space := model.SnapshotSpace{
			ID:       string(graphqlSpace.Id),
			Metadata: metadata,
		}

		spaces = append(spaces, space)
	}

	if err := job.DatabaseClient.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		UpdateAll: true,
	}).Create(&spaces).Error; err != nil {
		return err
	}

	return nil
}
