package job

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/governance"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/snapshot"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/snapshot/graphql"
	"github.com/sirupsen/logrus"

	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"
)

type (
	PullInfoStatus int32
)

const (
	PullInfoStatusNotLatest PullInfoStatus = 1
	PullInfoStatusLatest    PullInfoStatus = 2
)

type StatusStorage struct {
	Pos    int32          `json:"pos"`
	Status PullInfoStatus `json:"status"`
}

type SnapshotJobBase struct {
	Name           string
	SnapshotClient *snapshot.Client

	Limit          int32
	HighUpdateTime time.Duration
	LowUpdateTime  time.Duration
}

func (job *SnapshotJobBase) Check() error {
	if job.Name == "" {
		return fmt.Errorf("job name is empty")
	}

	if job.SnapshotClient == nil {
		return fmt.Errorf("snapshot worker is nil")
	}

	return nil
}

func (job *SnapshotJobBase) setSpaceInDB(ctx context.Context, graphqlSpaces []graphqlx.Space) (err error) {
	tracer := otel.Tracer("snapshot_space_job")
	_, trace := tracer.Start(ctx, "snapshot_space_job:setSpaceInDB")

	defer func() { opentelemetry.Log(trace, graphqlSpaces, nil, err) }()

	var spaces []governance.SnapshotSpace

	for _, graphqlSpace := range graphqlSpaces {
		metadata, err := json.Marshal(graphqlSpace)
		if err != nil {
			logrus.Warnf("[snapshot space job] marshal space metadata, error: %v", err)
			continue
		}

		space := governance.SnapshotSpace{
			ID:       string(graphqlSpace.Id),
			Metadata: metadata,
			Network:  string(graphqlSpace.Network),
		}

		spaces = append(spaces, space)
	}

	if err := database.Global().Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		UpdateAll: true,
	}).Create(&spaces).Error; err != nil {
		return err
	}

	return nil
}

func (job *SnapshotJobBase) upsertSpacesInDB(ctx context.Context, spaceIds []graphql.String) error {
	// get space info from url
	variable := snapshot.GetMultipleSpacesVariable{
		First: graphql.Int(len(spaceIds)),
		Where: graphqlx.SpaceWhere{
			IdArray: spaceIds,
		},
	}

	spaces, err := job.SnapshotClient.GetMultipleSpaces(ctx, variable)
	if err != nil {
		return fmt.Errorf("[snapshot proposal job] get multiple spaces, graphql error: %v", err)
	}

	// set space info in db
	if len(spaces) > 0 {
		if err := job.setSpaceInDB(ctx, spaces); err != nil {
			return fmt.Errorf("[snapshot proposal job], set space in db, db error: %v", err)
		}
		logrus.Infof("[snapshot proposal job] set space [%d]", len(spaces))
	}

	return nil
}
