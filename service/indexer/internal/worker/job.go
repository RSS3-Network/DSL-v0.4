package worker

import (
	"context"
	"time"

	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"

	"go.uber.org/zap"
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
	zap.L().Info("worker job start", zap.String("name", c.job.Name()), zap.String("spec", c.job.Spec()), zap.Duration("timeout", c.job.Timeout()))

	if err := c.job.Run(c.renewal); err != nil {
		logrus.Errorf("job %s throws an error: %s", c.job.Name(), err)
	}

	zap.L().Info("worker job end", zap.String("name", c.job.Name()), zap.String("spec", c.job.Spec()), zap.Duration("timeout", c.job.Timeout()))
}

func (c *cronJob) renewal(ctx context.Context, duration time.Duration) error {
	zap.L().Info("worker job renewal", zap.String("name", c.job.Name()), zap.String("spec", c.job.Spec()), zap.Duration("timeout", c.job.Timeout()), zap.Duration("duration", duration))

	return c.employer.Renewal(ctx, c.job.Name(), duration)
}

func NewCronJob(employer *shedlock.Employer, job Job) cron.Job {
	return &cronJob{
		employer: employer,
		job:      job,
	}
}
