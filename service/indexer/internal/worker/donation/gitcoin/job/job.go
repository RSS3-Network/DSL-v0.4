package job

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
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

	for i := 1; i <= 50; i++ {
		time.Sleep(time.Second)
		latestProject.ID += 1

		// request api
		newProject, err := job.requestGitcoinGrantApi(latestProject.ID)
		if err != nil || newProject == nil || newProject.ID == 0 {
			continue
		}

		newProject.AdminAddress = strings.ToLower(newProject.AdminAddress)

		// set db
		if err := database.Global().
			Model(&donation.GitcoinProject{}).
			Clauses(clause.OnConflict{
				DoNothing: true,
			}).
			Create(newProject).Error; err != nil {
			continue
		}
	}

	return nil
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

	// log
	out, _ := json.Marshal(response)
	logrus.Info("[gitcoin job] requestGitcoinGrantApi url: %v, response: %v", url, string(out))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("grantApi response code: %v", resp.StatusCode())
	}

	return &response.Grants, nil
}
