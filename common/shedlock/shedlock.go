package shedlock

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

type Employer struct {
	redisClient *redis.Client
	crontab     *cron.Cron
}

func (e *Employer) AddFunc(name, spec string, cmd func()) error {
	_, err := e.crontab.AddFunc(spec, func() {
		if e.DoLock(name) {
			defer func() {
				_ = e.UnLock(name)

				cmd()
			}()
		}
	})

	return err
}

func (e *Employer) AddJob(name, spec string, cmd cron.Job) error {
	_, err := e.crontab.AddFunc(spec, func() {
		if e.DoLock(name) {
			defer func() {
				_ = e.UnLock(name)

				cmd.Run()
			}()
		}
	})

	return err
}

func (e *Employer) DoLock(name string) bool {
	return e.redisClient.SetNX(context.Background(), name, 0, 0).Val()
}

func (e *Employer) UnLock(name string) bool {
	return e.redisClient.Del(context.Background(), name).Err() == nil
}

func (e *Employer) Start() {
	e.crontab.Start()
}

func (e *Employer) Stop() {
	e.crontab.Stop()
}

func New(redisClient *redis.Client) *Employer {
	return &Employer{
		redisClient: redisClient,
		crontab:     cron.New(),
	}
}
