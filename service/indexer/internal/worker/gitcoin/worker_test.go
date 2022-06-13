package gitcoin

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	type args struct {
		databaseClient *gorm.DB
		redisClient    *redis.Client
	}
	tests := []struct {
		name string
		args args
		want worker.Worker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.databaseClient, tt.args.redisClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Name(t *testing.T) {
	type fields struct {
		databaseClient         *gorm.DB
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				databaseClient:         tt.fields.databaseClient,
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
		databaseClient         *gorm.DB
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				databaseClient:         tt.fields.databaseClient,
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
		databaseClient         *gorm.DB
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				databaseClient:         tt.fields.databaseClient,
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
		databaseClient         *gorm.DB
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	type args struct {
		ctx       context.Context
		message   *protocol.Message
		transfers []model.Transfer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Transfer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				databaseClient:         tt.fields.databaseClient,
				redisClient:            tt.fields.redisClient,
				gitcoinProjectCacheKey: tt.fields.gitcoinProjectCacheKey,
			}
			got, err := s.Handle(tt.args.ctx, tt.args.message, tt.args.transfers)
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
		databaseClient         *gorm.DB
		redisClient            *redis.Client
		gitcoinProjectCacheKey string
	}
	tests := []struct {
		name   string
		fields fields
		want   []worker.Job
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				databaseClient:         tt.fields.databaseClient,
				redisClient:            tt.fields.redisClient,
				gitcoinProjectCacheKey: tt.fields.gitcoinProjectCacheKey,
			}
			if got := s.Jobs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Jobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
