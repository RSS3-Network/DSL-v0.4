package job

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ worker.Job = (*GitcoinProjectJob)(nil)

type GitcoinProjectJob struct {
	DatabaseClient         *gorm.DB
	RedisClient            *redis.Client
	GitcoinProjectCacheKey string
}

func (job *GitcoinProjectJob) Name() string {
	return "gitcoin_project_job"
}

func (job *GitcoinProjectJob) Spec() string {
	return "*/10 * * * *" // 10 min
}

func (job *GitcoinProjectJob) Timeout() time.Duration {
	return time.Minute * 5
}

func (job *GitcoinProjectJob) Run(renewal worker.RenewalFunc) error {
	logrus.Info("[gitcoin job] run")

	// query lastest gitcoin project id
	lastestProject := &model.GitcoinProject{}

	if err := job.DatabaseClient.
		Model(&model.GitcoinProject{}).
		Order("id DESC").
		First(&lastestProject).Error; err != nil {
		logrus.Errorf("[gitcoin job] get lastest grant, db error: %v", err)
		return err
	}

	for i := 1; i <= 10; i++ {
		// requeset api
		newProject, err := job.requestGitcoinGrantApi(lastestProject.ID + 1)
		if err != nil || newProject == nil || newProject.ID == 0 {
			continue
		}

		// set db
		if err := job.DatabaseClient.
			Model(&model.GitcoinProject{}).
			Clauses(clause.OnConflict{
				DoNothing: true,
			}).
			Create(newProject).Error; err != nil {
			logrus.Errorf("[gitcoin job] create lastest grant, db error: %v", err)
			continue
		}
	}

	go job.SetCache()

	return nil
}

func (job *GitcoinProjectJob) SetCache() {
	projectList := []*model.GitcoinProject{}
	if err := job.DatabaseClient.
		Model(&model.GitcoinProject{}).
		Order("id").
		Find(&projectList).Error; err != nil {
		logrus.Errorf("[gitcoin job] get gitcoin grants, db error: %v", err)
		return
	}

	for _, project := range projectList {
		if !project.Active {
			continue
		}

		projectByte, _ := json.Marshal(project)
		// set redis
		if err := job.RedisClient.HSet(
			context.Background(),
			job.GitcoinProjectCacheKey,
			project.AdminAddress,
			string(projectByte)).Err(); err != nil {
			logrus.Errorf("[gitcoin job] set redis error: %v", err)
			continue
		}
	}
}

func (job *GitcoinProjectJob) requestGitcoinGrantApi(id int) (*model.GitcoinProject, error) {
	var (
		url      = fmt.Sprintf("https://gitcoin.co/grants/v1/api/grant/%v", id)
		client   = resty.New()
		response = struct {
			Status int                  `json:"status"`
			Grants model.GitcoinProject `json:"grants"`
		}{}
	)

	resp, err := client.R().SetResult(&response).Get(url)
	if err != nil {
		logrus.Errorf("[gitcoin job] requestGitcoinGrantApi http get error: %v", err)
		return nil, err
	}

	if resp.IsError() {
		logrus.Errorf("[gitcoin job] requestGitcoinGrantApi http response error, code: %v", resp.StatusCode())
		return nil, fmt.Errorf("grantApi response code: %v", resp.StatusCode())
	}

	return &response.Grants, nil
}