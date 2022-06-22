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

var traceSpaceJob = "space_job"

func (job *SnapshotSpaceJob) Name() string {
	return "snapshot_space_job"
}

func (job *SnapshotSpaceJob) Spec() string {
	return "* * * * *" // 1 min
}

func (job *SnapshotSpaceJob) Timeout() time.Duration {
	return time.Minute * 2
}

func (job *SnapshotSpaceJob) Run(renewal worker.RenewalFunc) error {
	// nolint:ineffassign // just an initialization
	sleepTime := time.Second

	logrus.Info("[snapshot space job] run")

	for {
		pullInfoStatus, err := job.InnerJobRun()
		if err != nil {
			logrus.Errorf("[snapshot space job] run error: %v", err)
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

func (job *SnapshotSpaceJob) InnerJobRun() (PullInfoStatus, error) {
	err := job.Check()
	if err != nil {
		return PullInfoStatusNotLatest, fmt.Errorf("[snapshot space job] check error: %v", err)
	}

	ctx, runSnap := otel.Tracer(traceSpaceJob).Start(context.Background(), "run")
	defer runSnap.End()

	var statusStroge StatusStroge

	// get latest space id
	if job.RedisClient != nil {
		statusStroge, err = job.GetLastStatusFromCache(ctx)
		if err != nil {
			logrus.Errorf("[snapshot space job] get last status, db error: %v", err)
			statusStroge.Pos = 0
			statusStroge.Status = PullInfoStatusNotLatest
		}
	}

	if job.RedisClient == nil || err != nil {
		statusStroge.Pos, err = job.getSpaceTotalFromDB(ctx)
		if err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot space job] get space total from db, db error: %v", err)
		}
	}

	// get space info from url
	skip := statusStroge.Pos
	variable := snapshot.GetMultipleSpacesVariable{
		First:          graphql.Int(job.Limit),
		Skip:           graphql.Int(skip),
		OrderBy:        "created",
		OrderDirection: snapshot.OrderDirectionAsc,
	}

	spaces, err := job.SnapshotClient.GetMultipleSpaces(ctx, variable)
	if err != nil {
		return statusStroge.Status, fmt.Errorf("[snapshot space job] get multiple spaces, graphql error: %v", err)
	}

	// set space info in db
	if len(spaces) > 0 {
		if err := job.setSpaceInDB(ctx, spaces); err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot space job] pos[%d], set space in db, db error: %v", statusStroge.Pos, err)
		}
		logrus.Infof("[snapshot space job] pull skip [%d]", statusStroge.Pos)
	}

	skip = statusStroge.Pos + job.Limit
	// nolint:gocritic // dont' want change nan things
	if len(spaces) == 0 {
		statusStroge.Status = PullInfoStatusLatest
	} else if len(spaces) < int(job.Limit) {
		statusStroge.Pos = statusStroge.Pos + int32(len(spaces))
		statusStroge.Status = PullInfoStatusLatest
	} else {
		statusStroge.Pos = skip
		statusStroge.Status = PullInfoStatusNotLatest
	}

	// set space status in cache and db
	if job.RedisClient != nil {
		err = job.SetCurrentStatus(ctx, statusStroge)
		if err != nil {
			return statusStroge.Status, fmt.Errorf("[snapshot space job] set current status, db error: %v", err)
		}
	}

	return statusStroge.Status, nil
}

func (job *SnapshotSpaceJob) getSpaceTotalFromDB(ctx context.Context) (int32, error) {
	_, handlerSpan := otel.Tracer(traceSpaceJob).Start(ctx, "get_space_total_from_db")
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
	_, handlerSpan := otel.Tracer(traceSpaceJob).Start(ctx, "set_space_in_db")
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
			Network:  string(graphqlSpace.Network),
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
