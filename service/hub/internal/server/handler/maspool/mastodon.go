package maspool

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mattn/go-mastodon"
)

type Credentials struct {
	Username  string
	Password  string
	RateLimit int
}

type InstanceEntry struct {
	Client       *mastodon.Client
	Requests     int
	ResetTime    time.Time
	RateLimit    int
	Available    bool
	Availability chan bool
}

type MastodonPool struct {
	mu             sync.Mutex
	instances      map[string]*InstanceEntry
	requests       chan *InstanceEntry
	resetTimers    map[string]*time.Timer // 保存每个实例的定时器
	resetTimersMux sync.Mutex             // 定时器的互斥锁
}

func NewMastodonPool(serversAndCredentials map[string]Credentials) (*MastodonPool, error) {
	pool := &MastodonPool{
		instances:   make(map[string]*InstanceEntry),
		requests:    make(chan *InstanceEntry),
		resetTimers: make(map[string]*time.Timer),
	}

	for server, credentials := range serversAndCredentials {
		app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
			Server:     server,
			ClientName: "RSS3",
			Scopes:     "read write follow",
		})
		if err != nil {
			return nil, fmt.Errorf("error registering app for server %s: %w", server, err)
		}

		c := mastodon.NewClient(&mastodon.Config{
			Server:       server,
			ClientID:     app.ClientID,
			ClientSecret: app.ClientSecret,
		})

		err = c.Authenticate(context.Background(), credentials.Username, credentials.Password)
		if err != nil {
			return nil, fmt.Errorf("error authenticating for server %s: %w", server, err)
		}

		entry := &InstanceEntry{
			Client:       c,
			Requests:     0,
			ResetTime:    getNext5MinuteTime(),
			RateLimit:    credentials.RateLimit,
			Available:    true,
			Availability: make(chan bool, 1),
		}
		entry.Availability <- true

		pool.instances[server] = entry
	}

	go pool.processRequests()

	return pool, nil
}

func (p *MastodonPool) processRequests() {
	for instance := range p.requests {
		<-instance.Availability
		instance.Requests++
		// fmt.Printf("Sending request to %s, requests: %d, available: %v, resetTime: %s\n", instance.Client.Config.Server, instance.Requests, instance.Available, instance.ResetTime.Format(time.RFC3339))

		if instance.Requests > instance.RateLimit && time.Now().Before(instance.ResetTime) {
			instance.Available = false
			duration := time.Until(instance.ResetTime)

			p.resetTimersMux.Lock()
			timer, ok := p.resetTimers[instance.Client.Config.Server]
			if ok {
				timer.Stop() // Stop the previous timer
			}
			timer = time.AfterFunc(duration, func() {
				p.mu.Lock()
				instance.Available = true
				instance.Requests = 0
				instance.ResetTime = getNext5MinuteTime()
				instance.Availability <- true
				p.mu.Unlock()
			})
			p.resetTimers[instance.Client.Config.Server] = timer
			p.resetTimersMux.Unlock()
		} else {
			if instance.ResetTime.Before(time.Now()) {
				instance.Requests = 0
				instance.ResetTime = getNext5MinuteTime()
			}
			instance.Availability <- true
		}
	}
}

func (p *MastodonPool) GetClient() (*mastodon.Client, error) {
	var availableInstances []*InstanceEntry

	p.mu.Lock()
	for _, instance := range p.instances {
		if instance.Available {
			availableInstances = append(availableInstances, instance)
		}
	}
	p.mu.Unlock()

	if len(availableInstances) == 0 {
		return nil, fmt.Errorf("mastodon api rate limit exceeded, try again after %s", getNext5MinuteTime().Format(time.RFC3339))
	}

	// Randomly select an available instance
	instance := availableInstances[time.Now().UnixNano()%int64(len(availableInstances))]

	return instance.Client, nil
}

func (p *MastodonPool) DoRequest(ctx context.Context, f func(client *mastodon.Client) error) error {
	client, err := p.GetClient()
	if err != nil {
		return err
	}

	err = f(client)
	if err != nil {
		p.mu.Lock()
		instance := p.instances[client.Config.Server]
		select {
		case instance.Availability <- true:
		default:
		}
		p.mu.Unlock()

		return err
	}

	p.requests <- p.instances[client.Config.Server]

	return nil
}

func getNext5MinuteTime() time.Time {
	now := time.Now()
	minutes := now.Minute()
	remainder := minutes % 5
	next5Minute := now.Add(time.Duration(5-remainder) * time.Minute)

	return time.Date(next5Minute.Year(), next5Minute.Month(), next5Minute.Day(), next5Minute.Hour(), next5Minute.Minute(), 0, 0, next5Minute.Location())
}
