package shedlock

import (
	"context"
	"time"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/robfig/cron/v3"
)

type Employer struct {
	crontab *cron.Cron
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

func (e *Employer) WaitForLock(ctx context.Context, name string, timeout time.Duration) error {
	if e.DoLock(name, timeout) {
		return nil
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if e.DoLock(name, timeout) {
				return nil
			}
		}
	}
}

func New() *Employer {
	return &Employer{
		crontab: cron.New(),
	}
}
