package job

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"strconv"
)

type (
	PullInfoStatus int
)

const (
	PullInfoStatusNotLatest PullInfoStatus = 1
	PullInfoStatusLatest    PullInfoStatus = 2
)

type StatusStroge struct {
	pos    int32
	status PullInfoStatus
}

type SnapshotJobBase struct {
	Name           string
	DatabaseClient *gorm.DB
	RedisClient    *redis.Client
	SnapshotClient *snapshot.Client
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
		pos:    0,
		status: PullInfoStatusNotLatest,
	}

	// If you can pull data from redis normally,
	// you don't need to enter the database to get it.

	_, handlerSpan := tracer.Start(ctx, "get_by_cache")
	defer handlerSpan.End()

	resultMap, err := job.RedisClient.HGetAll(ctx, statusKey).Result()
	if err != nil {
		logrus.Errorf("get %s from cache error:%+v", statusKey, err)
	}

	if len(resultMap) != 0 {
		pos, ok := resultMap["pos"]
		if ok {
			currPos, err := strconv.ParseInt(pos, 10, 0)
			if err != nil {
				logrus.Warnf("parse [%s] pos [%s] error:%+v", statusKey, pos, err)
			}
			statusStroge.pos = int32(currPos)
		}

		status, ok := resultMap["status"]
		if ok {
			currStatus, err := strconv.Atoi(status)
			if err != nil {
				logrus.Warnf("parse [%s] status [%s] error:%+v", statusKey, pos, err)
			} else {
				statusStroge.status = PullInfoStatus(currStatus)
			}
		}
		return statusStroge, nil
	}

	return StatusStroge{}, nil
}

func (job *SnapshotJobBase) SetCurrentStatus(ctx context.Context, filter StatusStroge) error {
	if job.Name == "" {
		return fmt.Errorf("job name is empty")
	}

	if job.RedisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	// Hset
	if filter.pos <= 0 {
		return fmt.Errorf("pos is less than 0")
	}

	job.RedisClient.HSet(ctx, job.Name+"_status", "pos", strconv.Itoa(int(filter.pos)))
	return nil
}
