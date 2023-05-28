package maspool

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/mattn/go-mastodon"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"

	"go.uber.org/zap"
)

type Credential struct {
	Username string
	Password string
}

type Instance struct {
	Client    *mastodon.Client
	RateLimit int
	LastReset time.Time
	Available bool
	Lock      sync.Mutex
}

type InstancePool struct {
	instances []*Instance
	mu        sync.Mutex
}

func NewInstancePool(serversAndCredentials map[string]Credential, rateLimit int) (*InstancePool, error) {
	pool := &InstancePool{}

	for server, credential := range serversAndCredentials {
		app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
			Server:     server,
			ClientName: "RSS3",
			Scopes:     "read write follow",
		})
		if err != nil {
			loggerx.Global().Warn("error registering app for server", zap.String("server", server), zap.Error(err))
			continue
		}

		c := mastodon.NewClient(&mastodon.Config{
			Server:       server,
			ClientID:     app.ClientID,
			ClientSecret: app.ClientSecret,
		})

		err = c.Authenticate(context.Background(), credential.Username, credential.Password)
		if err != nil {
			loggerx.Global().Warn("error authenticating for server", zap.String("server", server), zap.Error(err))
			continue
		}

		instance := &Instance{
			Client:    c,
			RateLimit: rateLimit,
			LastReset: GetNext5MinuteTime(),
			Available: true,
		}

		loggerx.Global().Info("create mastodon instance success", zap.String("server", server))

		pool.instances = append(pool.instances, instance)
	}

	return pool, nil
}

func (p *InstancePool) GetAvailableInstance(count int) *Instance {
	p.mu.Lock()
	defer p.mu.Unlock()

	currentTime := time.Now()

	loggerx.Global().Info("mastodon instance number", zap.Int("servers", len(p.instances)))

	if len(p.instances) == 0 {
		loggerx.Global().Error("mastodon has no instance", zap.Int("servers", len(p.instances)))
	}
	// restore to initial state
	if currentTime.After(p.instances[0].LastReset) {
		for _, instance := range p.instances {
			instance.Lock.Lock()
			instance.Available = true
			instance.LastReset = GetNext5MinuteTime()
			instance.RateLimit = config.ConfigHub.Mastodon.RateLimit
			instance.Lock.Unlock()
		}
	}

	for _, instance := range p.instances {
		instance.Lock.Lock()

		switch {
		case !instance.Available:
			if currentTime.Before(instance.LastReset) {
				instance.Lock.Unlock()
				continue
			}
			instance.RateLimit = config.ConfigHub.Mastodon.RateLimit
			instance.LastReset = GetNext5MinuteTime()
			instance.Available = true

		case instance.Available && instance.RateLimit >= count && currentTime.Before(instance.LastReset):
			instance.RateLimit -= count
			if instance.RateLimit == 0 {
				instance.Available = false
			}
			instance.Lock.Unlock()
			sort.SliceStable(p.instances, func(i, j int) bool {
				return p.instances[i].RateLimit > p.instances[j].RateLimit
			})
			return instance
		}

		instance.Lock.Unlock()
	}

	return nil
}

func GetNext5MinuteTime() time.Time {
	now := time.Now()
	minutes := now.Minute()
	remainder := minutes % 5
	next5Minute := now.Add(time.Duration(5-remainder) * time.Minute)

	return time.Date(next5Minute.Year(), next5Minute.Month(), next5Minute.Day(), next5Minute.Hour(), next5Minute.Minute(), 0, 0, next5Minute.Location())
}
