package job

import (
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SnapshotSpaceJob struct {
	DatabaseClient *gorm.DB
}

func (job *SnapshotSpaceJob) Name() string {
	return "snapshot_space_job"
}

func (job *SnapshotSpaceJob) Spec() string {
	return "*/10 * * * *" // 10 min
}

func (job *SnapshotSpaceJob) Timeout() time.Duration {
	return time.Minute * 5
}

func (job *SnapshotSpaceJob) Run() {
	logrus.Info("[snapshot space job] run")
	// get latest space id
	lastestSpace := &model.Space{}

	if err := job.DatabaseClient.
		Model(&model.Space{}).
		Order("id DESC").
		First(&lastestSpace).Error; err != nil {
		logrus.Errorf("[gitcoin job] get lastest grant, db error: %v", err)
		return
	}

	// if id > current db id, then update cache

}
