package gitcoin

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
)

func TestNew(t *testing.T) {
	type args struct {
		redisClient *redis.Client
	}
	var tests []struct {
		name string
		args args
		want worker.Worker
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.redisClient, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Name(t *testing.T) {
	type fields struct {
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	var tests []struct {
		name   string
		fields fields
		want   string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redisClient:            tt.fields.redisClient,
				gitcoinProjectCacheKey: tt.fields.gitcoinProjectCacheKey,
			}
			if got := s.Name(); got != tt.want {
				t.Errorf("service.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Networks(t *testing.T) {
	type fields struct {
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	var tests []struct {
		name   string
		fields fields
		want   []string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redisClient:            tt.fields.redisClient,
				gitcoinProjectCacheKey: tt.fields.gitcoinProjectCacheKey,
			}
			if got := s.Networks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Networks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Initialize(t *testing.T) {
	type fields struct {
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	type args struct {
		ctx context.Context
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redisClient:            tt.fields.redisClient,
				gitcoinProjectCacheKey: tt.fields.gitcoinProjectCacheKey,
			}
			if err := s.Initialize(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("service.Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Handle(t *testing.T) {
	type fields struct {
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	type args struct {
		ctx          context.Context
		message      *protocol.Message
		transactions []model.Transaction
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []model.Transaction
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redisClient:            tt.fields.redisClient,
				gitcoinProjectCacheKey: tt.fields.gitcoinProjectCacheKey,
			}
			got, err := s.Handle(tt.args.ctx, tt.args.message, tt.args.transactions)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Jobs(t *testing.T) {
	type fields struct {
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	var tests []struct {
		name   string
		fields fields
		want   []worker.Job
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redisClient:            tt.fields.redisClient,
				gitcoinProjectCacheKey: tt.fields.gitcoinProjectCacheKey,
			}
			if got := s.Jobs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Jobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
