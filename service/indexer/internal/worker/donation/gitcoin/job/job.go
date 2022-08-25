package job

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/donation"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

var _ worker.Job = (*GitcoinProjectJob)(nil)

type GitcoinProjectJob struct {
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

	// query latest gitcoin project id
	latestProject := &donation.GitcoinProject{}

	if err := database.Global().
		Model(&donation.GitcoinProject{}).
		Order("id DESC").
		First(&latestProject).Error; err != nil {
		logrus.Errorf("[gitcoin job] get latest grant, db error: %v", err)
		return err
	}

	for i := 1; i <= 100; i++ {
		// request api
		newProject, err := job.requestGitcoinGrantApi(latestProject.ID + 1)
		if err != nil || newProject == nil || newProject.ID == 0 {
			continue
		}

		// set db
		if err := database.Global().
			Model(&donation.GitcoinProject{}).
			Clauses(clause.OnConflict{
				DoNothing: true,
			}).
			Create(newProject).Error; err != nil {
			continue
		}

		latestProject.ID += 1
	}

	go job.SetCache()

	return nil
}

func (job *GitcoinProjectJob) SetCache() {
	var projectList []*donation.GitcoinProject
	if err := database.Global().
		Model(&donation.GitcoinProject{}).
		Order("id").
		Find(&projectList).Error; err != nil {
		logrus.Errorf("[gitcoin job] get gitcoin grants, db error: %v", err)
		return
	}

	for _, project := range projectList {
		projectByte, _ := json.Marshal(project)
		// set redis
		if err := cache.Global().HSet(
			context.Background(),
			job.GitcoinProjectCacheKey,
			project.AdminAddress,
			string(projectByte)).Err(); err != nil {
			logrus.Errorf("[gitcoin job] set redis error: %v", err)
			continue
		}
	}
}

func (job *GitcoinProjectJob) requestGitcoinGrantApi(id int) (*donation.GitcoinProject, error) {
	var (
		url      = fmt.Sprintf("https://gitcoin.co/grants/v1/api/grant/%v", id)
		client   = resty.New()
		response = struct {
			Status int                     `json:"status"`
			Grants donation.GitcoinProject `json:"grants"`
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
