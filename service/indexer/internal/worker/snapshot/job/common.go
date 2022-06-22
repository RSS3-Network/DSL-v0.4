package job

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type (
	PullInfoStatus int32
)

const (
	PullInfoStatusNotLatest PullInfoStatus = 1
	PullInfoStatusLatest    PullInfoStatus = 2
)

type StatusStroge struct {
	Pos    int32          `json:"pos"`
	Status PullInfoStatus `json:"status"`
}

type SnapshotJobBase struct {
	Name           string
	DatabaseClient *gorm.DB
	RedisClient    *redis.Client
	SnapshotClient *snapshot.Client

	Limit          int32
	HighUpdateTime time.Duration
	LowUpdateTime  time.Duration
}

func (job *SnapshotJobBase) Check() error {
	if job.Name == "" {
		return fmt.Errorf("job name is empty")
	}

	if job.DatabaseClient == nil {
		return fmt.Errorf("database client is nil")
	}

	if job.SnapshotClient == nil {
		return fmt.Errorf("snapshot client is nil")
	}

	return nil
}

func (job *SnapshotJobBase) GetLastStatusFromCache(ctx context.Context) (StatusStroge, error) {
	if job.Name == "" {
		return StatusStroge{}, fmt.Errorf("job name is empty")
	}

	if job.RedisClient == nil {
		return StatusStroge{}, fmt.Errorf("redis client is nil")
	}

	statusKey := job.Name + "_status"
	tracerName := "get_" + statusKey
	tracer := otel.Tracer(tracerName)
	statusStroge := StatusStroge{
		Pos:    0,
		Status: PullInfoStatusNotLatest,
	}

	_, handlerSpan := tracer.Start(ctx, "get_by_cache")
	defer handlerSpan.End()

	data, err := job.RedisClient.Get(ctx, statusKey).Result()
	if err != nil {
		return StatusStroge{}, err
	}

	if err = json.Unmarshal([]byte(data), &statusStroge); err != nil {
		return StatusStroge{}, fmt.Errorf("unmarshal %s from cache error:%+v", statusKey, err)
	}

	return statusStroge, nil
}

func (job *SnapshotJobBase) SetCurrentStatus(ctx context.Context, stroge StatusStroge) error {
	if job.Name == "" {
		return fmt.Errorf("job name is empty")
	}

	if job.RedisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	if stroge.Pos <= 0 {
		return fmt.Errorf("pos is less than 0")
	}

	data, err := json.Marshal(stroge)
	if err != nil {
		return fmt.Errorf("marshal %+v to json error:%+v", stroge, err)
	}

	job.RedisClient.Set(ctx, job.Name+"_status", data, 0)

	return nil
}
