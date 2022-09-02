package worker

import (
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/internalModel"
	"github.com/robfig/cron/v3"
)

// Need to make Job compatible with cron.Job
var _ cron.Job = (*internalModel.cronJob)(nil)

func NewCronJob(job internalModel.Job) cron.Job {
	return &internalMdoel.CronJob{
		job: job,
	}
}
