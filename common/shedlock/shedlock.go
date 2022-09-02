package shedlock

import (
	"context"
	"sync"
	"time"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/robfig/cron/v3"
)

type Employer struct {
	crontab *cron.Cron
}

var (
	employer     *Employer
	globalLocker sync.RWMutex
)

func Global() *Employer {
	globalLocker.RLock()

	defer globalLocker.RUnlock()

	return employer
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
	return cache.Global().SetNX(context.Background(), name, 0, timeout).Val()
}

func (e *Employer) UnLock(name string) bool {
	if err := cache.Global().Del(context.Background(), name).Err(); err != nil {
		return false
	}

	return true
}

func (e *Employer) Start() {
	e.crontab.Start()
}

func (e *Employer) Stop() {
	e.crontab.Stop()
}

func (e *Employer) Renewal(ctx context.Context, name string, duration time.Duration) error {
	return cache.Global().SetEX(ctx, name, 0, duration).Err()
}

func New() {
	employer = &Employer{
		crontab: cron.New(),
	}
}
