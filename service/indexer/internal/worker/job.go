package worker

import (
	"context"
	"time"

	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type RenewalFunc func(ctx context.Context, duration time.Duration) error

type Job interface {
	Name() string
	Spec() string
	Timeout() time.Duration
	Run(renewal RenewalFunc) error
}

// Need to make Job compatible with cron.Job
var _ cron.Job = (*cronJob)(nil)

type cronJob struct {
	employer *shedlock.Employer
	job      Job
}

func (c *cronJob) Run() {
	if err := c.job.Run(c.renewal); err != nil {
		logrus.Errorf("job %s throws an error: %s", c.job.Name(), err)
	}
}

func (c *cronJob) renewal(ctx context.Context, duration time.Duration) error {
	return c.employer.Renewal(ctx, c.job.Name(), duration)
}

func NewCronJob(employer *shedlock.Employer, job Job) cron.Job {
	return &cronJob{
		employer: employer,
		job:      job,
	}
}
