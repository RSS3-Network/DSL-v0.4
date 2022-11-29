package polygonstaking

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

var once sync.Once

func initialize(t *testing.T) {
	once.Do(func() {
		config.Initialize()

		logger, err := zap.NewDevelopment()
		assert.NoError(t, err, "initialize logger")

		zap.ReplaceGlobals(logger)

		cacheClient, err := cache.Dial(config.ConfigIndexer.Redis)
		assert.NoError(t, err)

		cache.ReplaceGlobal(cacheClient)
	})
}

func TestNew(t *testing.T) {
	initialize(t)

	testcases := []struct {
		name string
		want assert.ValueAssertionFunc
	}{
		{
			name: "default",
			want: assert.NotNil,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			job := New(cache.Global())

			testcase.want(t, job, fmt.Sprintf("nil job"))

			assert.NotEmpty(t, job.Name(), "empty name")
			assert.NotEmpty(t, job.Spec(), "empty spec")
			assert.NotZero(t, job.Timeout(), "zero timeout")
		})
	}
}

func TestJob_Run(t *testing.T) {
	initialize(t)

	type fields struct {
		redisClient *redis.Client
	}

	testcases := []struct {
		name   string
		fields fields
		want   assert.ErrorAssertionFunc
	}{
		{
			name: "default",
			fields: fields{
				redisClient: cache.Global(),
			},
			want: assert.NoError,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			job := &Job{
				redisClient: testcase.fields.redisClient,
			}

			renewalFunc := func(ctx context.Context, duration time.Duration) error {
				return nil
			}

			testcase.want(t, job.Run(renewalFunc))
		})
	}
}
