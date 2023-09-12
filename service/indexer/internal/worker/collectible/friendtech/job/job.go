package job

import (
	"context"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/collectibe"
	friendtechClient "github.com/naturalselectionlabs/pregod/common/datasource/friendtech"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ worker.Job = (*Job)(nil)

type Job struct {
	databaseClient *gorm.DB
	client         *friendtechClient.Client
	crontab        *cron.Cron
}

func (job *Job) Name() string {
	return "friend.tech_metadata"
}

func (job *Job) Spec() string {
	return "* * * * *" // 1 min
}

func (job *Job) Timeout() time.Duration {
	return time.Minute * 2
}

func (job *Job) Run(renewal worker.RenewalFunc) error {
	var sleepTime time.Duration

	logrus.Info("Fetch friend.tech metadata run")

	for {
		status, err := job.fetchFriendTechUser(context.Background())
		if err != nil {
			logrus.Errorf("friend.tech_metadata run error: %v", err)
		}

		if status {
			sleepTime = time.Millisecond * 200
		} else {
			sleepTime = time.Minute * 5
		}

		if err = renewal(context.Background(), time.Minute*2); err != nil {
			return err
		}

		time.Sleep(sleepTime)
	}
}

func (job *Job) fetchFriendTechUser(ctx context.Context) (bool, error) {
	id, err := job.getLatestIDFromDB(context.Background())
	if err != nil {
		return false, err
	}

	var (
		user *friendtechClient.UserResponse
		flag = true
	)
	for ; flag; id++ {
		user, err = job.client.GetUserMetaByID(ctx, id+1)

		if err != nil || user.ID == 0 {
			return false, err
		}

		if user.ID != -1 {
			flag = false
		}
	}

	err = job.setItemToDB(ctx, *user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (job *Job) getLatestIDFromDB(_ context.Context) (int64, error) {
	var item collectibe.FriendTech

	if err := database.Global().
		Model(&collectibe.FriendTech{}).
		Order("ID DESC").
		Limit(1).
		Take(&item).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return 10, nil
		}
		logrus.Errorf("get latest item, db error: %v", err)
		return 0, err
	}

	return item.ID, nil
}

func (job *Job) setItemToDB(_ context.Context, user friendtechClient.UserResponse) (err error) {
	item := collectibe.FriendTech{
		ID:              int64(user.ID),
		Address:         strings.ToLower(user.Address),
		TwitterUsername: user.TwitterUsername,
		TwitterName:     user.TwitterName,
		TwitterPfpUrl:   user.TwitterPfpURL,
		TwitterUserId:   user.TwitterUserID,
	}

	if err := database.Global().Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"id"}),
		UpdateAll: true,
	}).Create(&item).Error; err != nil {
		return err
	}

	return nil
}

func New() worker.Job {
	return &Job{
		databaseClient: database.Global(),
		client:         friendtechClient.NewClient(),
		crontab:        cron.New(),
	}
}
