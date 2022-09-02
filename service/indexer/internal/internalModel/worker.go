package internalModel

import (
	"context"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/sirupsen/logrus"
)

type Worker interface {
	Name() string
	Networks() []string
	Initialize(ctx context.Context) error
	Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error)
	Jobs() []Job
}

type Job interface {
	Name() string
	Spec() string
	Timeout() time.Duration
	Run(renewal RenewalFunc) error
}

type RenewalFunc func(ctx context.Context, duration time.Duration) error

type CronJob struct {
	job Job
}

func (c *CronJob) Run() {
	if err := c.job.Run(c.renewal); err != nil {
		logrus.Errorf("job %s throws an error: %s", c.job.Name(), err)
	}
}

func (c *CronJob) renewal(ctx context.Context, duration time.Duration) error {
	return shedlock.Global().Renewal(ctx, c.job.Name(), duration)
}
