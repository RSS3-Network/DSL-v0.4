package worker

import (
	"time"

	"github.com/robfig/cron/v3"
)

type Job interface {
	Name() string
	Spec() string
	Timeout() time.Duration

	cron.Job
}
