package job

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"gorm.io/gorm"
)

var (
	_ worker.Job = (*GitcoinProjectJob)(nil)
)

type GitcoinProjectJob struct {
	DatabaseClient *gorm.DB
}

func (job *GitcoinProjectJob) Name() string {
	return "gitcoin_project_job"
}

func (job *GitcoinProjectJob) Spec() string {
	return "* * * * *"
}

func (job *GitcoinProjectJob) Timeout() time.Duration {
	return time.Minute * 1
}

func (job *GitcoinProjectJob) Run() {
	// query lastest gitcoin grant id
	

	// requeset api

	// set db
}

func (job *GitcoinProjectJob) requestGitcoinGrantApi(id int64) (*model.GitcoinProject, error) {
	var (
		gitcoinProject = &model.GitcoinProject{}
		url            = fmt.Sprintf("https://gitcoin.co/grants/v1/api/grant/%v", id)
		client         = resty.New()
	)

	response, err := client.R().SetResult(&gitcoinProject).Get(url)
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, fmt.Errorf("response code: %v", response.StatusCode)
	}

	return gitcoinProject, nil
}
