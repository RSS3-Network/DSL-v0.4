package job

import (
	"context"
	"encoding/json"
	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	graphqlx "github.com/naturalselectionlabs/pregod/common/snapshot/graphql"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"
	"time"

	"github.com/sirupsen/logrus"
)

type SnapshotSpaceJob struct {
	SnapshotJobBase
}

var TraceSpaceJob = "space_job"

const LimitOnce = 10

func (job *SnapshotSpaceJob) Name() string {
	return "snapshot_space_job"
}

func (job *SnapshotSpaceJob) Spec() string {
	return "*/10 * * * *" // 10 min
}

func (job *SnapshotSpaceJob) Timeout() time.Duration {
	return time.Minute * 5
}

func (job *SnapshotSpaceJob) Run() {
	err := job.Check()
	if err != nil {
		logrus.Errorf("[snapshot space job] check error: %v", err)
		return
	}

	ctx, runSnap := otel.Tracer(TraceSpaceJob).Start(context.Background(), "run")
	defer runSnap.End()

	logrus.Info("[snapshot space job] run")

	// get latest space id
	statusStroge, err := job.GetLastStatusFromCache(ctx)
	if err != nil {
		logrus.Errorf("[snapshot space job] get last status, db error: %v", err)
		statusStroge.pos = 0
		statusStroge.status = PullInfoStatusNotLatest
	}

	// get space info from url
	skip := statusStroge.pos + LimitOnce
	variable := snapshot.GetMultipleSpacesVariable{
		First:          LimitOnce,
		Skip:           graphql.Int(skip),
		OrderBy:        "created",
		OrderDirection: snapshot.OrderDirectionAsc,
	}

	spaces, err := job.SnapshotClient.GetMultipleSpaces(ctx, variable)
	if err != nil {
		logrus.Errorf("[snapshot space job] get multiple spaces, graphql error: %v", err)
		return
	}

	setInDb := false
	if len(spaces) == 0 {
		statusStroge.status = PullInfoStatusLatest
	} else if len(spaces) < LimitOnce {
		setInDb = true
		statusStroge.pos = statusStroge.pos + int32(len(spaces))
		statusStroge.status = PullInfoStatusLatest
	} else {
		setInDb = true
		statusStroge.pos = skip
		statusStroge.status = PullInfoStatusNotLatest
	}

	// set space info in db
	if setInDb {
		if err := job.setSpaceInDB(ctx, spaces); err != nil {
			logrus.Errorf("[snapshot space job] set space in db, db error: %v, set pos[%d]", err, statusStroge.pos)
			return
		}
	}

	// set space status in cache and db
	err = job.SetCurrentStatus(ctx, statusStroge)
	if err != nil {
		logrus.Errorf("[snapshot space job] set current status, db error: %v", err)
		return
	}
}

func (job *SnapshotSpaceJob) SpaceRun() {

}

func (job *SnapshotSpaceJob) getSpaceTotalFromDB(ctx context.Context) (int, error) {
	_, handlerSpan := otel.Tracer(TraceSpaceJob).Start(ctx, "get_space_total_from_db")
	defer handlerSpan.End()

	var count int64

	if err := job.DatabaseClient.
		Select("id").
		Count(&count).Error; err != nil {
		logrus.Errorf("get space total, db error: %v", err)
		return 0, err
	}

	return int(count), nil
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
