package job

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/database/model/donation"
)

func TestGitcoinProjectJob_Name(t *testing.T) {
	var tests []struct {
		name string
		job  *GitcoinProjectJob
		want string
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
	var tests []struct {
		name string
		job  *GitcoinProjectJob
		want string
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
	var tests []struct {
		name string
		job  *GitcoinProjectJob
		want time.Duration
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
	var tests []struct {
		name string
		job  *GitcoinProjectJob
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := &GitcoinProjectJob{}
			if err := job.Run(func(ctx context.Context, duration time.Duration) error {
				return nil
			}); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestGitcoinProjectJob_requestGitcoinGrantApi(t *testing.T) {
	type args struct {
		id int
	}
	var tests []struct {
		name    string
		job     *GitcoinProjectJob
		args    args
		want    *donation.GitcoinProject
		wantErr bool
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
		RedisClient            *redis.Client
		GitcoinProjectCacheKey string
	}
	var tests []struct {
		name   string
		fields fields
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := &GitcoinProjectJob{
				RedisClient:            tt.fields.RedisClient,
				GitcoinProjectCacheKey: tt.fields.GitcoinProjectCacheKey,
			}
			job.SetCache()
		})
	}
}
