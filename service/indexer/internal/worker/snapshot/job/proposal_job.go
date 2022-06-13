package job

import (
	"time"

	"gorm.io/gorm"
)

type SnapshotProposalJob struct {
	DatabaseClient *gorm.DB
}

func (job *SnapshotProposalJob) Name() string {
	return "snapshot_proposal_job"
}

func (job *SnapshotProposalJob) Spec() string {
	return "*/10 * * * *" // 10 min
}

func (job *SnapshotProposalJob) Timeout() time.Duration {
	return time.Minute * 5
}

func (job *SnapshotProposalJob) Run() {}
