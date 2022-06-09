package job

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"gorm.io/gorm"
)

func TestGitcoinProjectJob_Name(t *testing.T) {
	tests := []struct {
		name string
		job  *GitcoinProjectJob
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := &GitcoinProjectJob{}
			if got := job.Name(); got != tt.want {
				t.Errorf("GitcoinProjectJob.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitcoinProjectJob_Spec(t *testing.T) {
	tests := []struct {
		name string
		job  *GitcoinProjectJob
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := &GitcoinProjectJob{}
			if got := job.Spec(); got != tt.want {
				t.Errorf("GitcoinProjectJob.Spec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitcoinProjectJob_Timeout(t *testing.T) {
	tests := []struct {
		name string
		job  *GitcoinProjectJob
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := &GitcoinProjectJob{}
			if got := job.Timeout(); got != tt.want {
				t.Errorf("GitcoinProjectJob.Timeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitcoinProjectJob_Run(t *testing.T) {
	tests := []struct {
		name string
		job  *GitcoinProjectJob
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := &GitcoinProjectJob{}
			job.Run()
		})
	}
}

func TestGitcoinProjectJob_requestGitcoinGrantApi(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		job     *GitcoinProjectJob
		args    args
		want    *model.GitcoinProject
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := &GitcoinProjectJob{}
			got, err := job.requestGitcoinGrantApi(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GitcoinProjectJob.requestGitcoinGrantApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitcoinProjectJob.requestGitcoinGrantApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitcoinProjectJob_SetCache(t *testing.T) {
	type fields struct {
		DatabaseClient         *gorm.DB
		RedisClient            *redis.Client
		GitcoinProjectCacheKey string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := &GitcoinProjectJob{
				DatabaseClient:         tt.fields.DatabaseClient,
				RedisClient:            tt.fields.RedisClient,
				GitcoinProjectCacheKey: tt.fields.GitcoinProjectCacheKey,
			}
			job.SetCache()
		})
	}
}
