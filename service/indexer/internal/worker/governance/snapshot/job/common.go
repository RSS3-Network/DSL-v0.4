package job

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/snapshot"
	"go.opentelemetry.io/otel"
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

func (job *SnapshotJobBase) GetLastStatusFromCache(ctx context.Context) (statusStroge StatusStorage, err error) {
	tracer := otel.Tracer("snapshot_job")
	_, trace := tracer.Start(ctx, "snapshot_job:GetLastStatusFromCache")

	defer func() { opentelemetry.Log(trace, nil, statusStroge, err) }()

	if job.Name == "" {
		return StatusStorage{}, fmt.Errorf("job name is empty")
	}

	if cache.Global() == nil {
		return StatusStorage{}, fmt.Errorf("redis worker is nil")
	}

	statusKey := job.Name + "_status"
	statusStroge = StatusStorage{
		Pos:    0,
		Status: PullInfoStatusNotLatest,
	}

	data, err := cache.Global().Get(ctx, statusKey).Result()
	if err != nil {
		return StatusStorage{}, err
	}

	if err = json.Unmarshal([]byte(data), &statusStroge); err != nil {
		return StatusStorage{}, fmt.Errorf("unmarshal %s from cache error:%+v", statusKey, err)
	}

	return statusStroge, nil
}

func (job *SnapshotJobBase) SetCurrentStatus(ctx context.Context, stroge StatusStorage) error {
	if job.Name == "" {
		return fmt.Errorf("job name is empty")
	}

	if cache.Global() == nil {
		return fmt.Errorf("redis worker is nil")
	}

	if stroge.Pos <= 0 {
		return fmt.Errorf("pos is less than 0")
	}

	data, err := json.Marshal(stroge)
	if err != nil {
		return fmt.Errorf("marshal %+v to json error:%+v", stroge, err)
	}

	cache.Global().Set(ctx, job.Name+"_status", data, 0)

	return nil
}
