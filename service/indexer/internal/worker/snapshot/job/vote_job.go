package job

import (
	"time"

	"gorm.io/gorm"
)

type SnapshotVoteJob struct {
	DatabaseClient *gorm.DB
}

func (job *SnapshotVoteJob) Name() string {
	return "snapshot_vote_job"
}

func (job *SnapshotVoteJob) Spec() string {
	return "*/10 * * * *" // 10 min
}

func (job *SnapshotVoteJob) Timeout() time.Duration {
	return time.Minute * 5
}

func (job *SnapshotVoteJob) Run() {
	//
}
