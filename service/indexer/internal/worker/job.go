package worker

import "github.com/robfig/cron/v3"

type Job interface {
	Name() string
	Spec() string

	cron.Job
}
