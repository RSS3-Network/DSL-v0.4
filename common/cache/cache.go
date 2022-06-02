package cache

import (
	"context"
	"crypto/tls"

	"github.com/go-redis/redis/v8"
	configx "github.com/naturalselectionlabs/pregod/common/config"
)

func Dial(config *configx.Redis) (*redis.Client, error) {
	ctx := context.Background()
	var tlsConfig *tls.Config
	if config.TLSEnabled {
		// In AWS if a Redis has a password, it must use TLS
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:      config.Addr,
		Password:  config.Password,
		DB:        config.DB,
		TLSConfig: tlsConfig,
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
