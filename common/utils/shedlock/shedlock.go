package shedlock

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

type Employer struct {
	redisClient *redis.Client
	crontab     *cron.Cron
}

func (e *Employer) AddFunc(name, spec string, timeout time.Duration, cmd func()) error {
	_, err := e.crontab.AddFunc(spec, func() {
		if e.DoLock(name, timeout) {
			defer func() {
				_ = e.UnLock(name)

				cmd()
			}()
		}
	})

	return err
}

func (e *Employer) AddJob(name, spec string, timeout time.Duration, cmd cron.Job) error {
	_, err := e.crontab.AddFunc(spec, func() {
		if e.DoLock(name, timeout) {
			defer func() {
				cmd.Run()

				_ = e.UnLock(name)
			}()
		}
	})

	return err
}

func (e *Employer) DoLock(name string, timeout time.Duration) bool {
	return e.redisClient.SetNX(context.Background(), name, 0, timeout).Val()
}

func (e *Employer) UnLock(name string) bool {
	// TODO Need to consider the minimum preemption time of the lock

	return true
}

func (e *Employer) Start() {
	e.crontab.Start()
}

func (e *Employer) Stop() {
	e.crontab.Stop()
}

func (e *Employer) Renewal(ctx context.Context, name string, duration time.Duration) error {
	return e.redisClient.SetEX(ctx, name, 0, duration).Err()
}

func New(redisClient *redis.Client) *Employer {
	return &Employer{
		redisClient: redisClient,
		crontab:     cron.New(),
	}
}
