package cache

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-redis/redis/v8"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/vmihailenco/msgpack"
)

var redisClient *redis.Client

func Dial(config *configx.Redis) (*redis.Client, error) {
	ctx := context.Background()
	var tlsConfig *tls.Config
	if config.TLSEnabled {
		// In AWS if a Redis has a password, it must use TLS
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	redisClient = redis.NewClient(&redis.Options{
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

// return exists, error
func GetMsgPack(ctx context.Context, key string, dest interface{}) (bool, error) {
	data, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, msgpack.Unmarshal([]byte(data), dest)
}

func SetMsgPack(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}
	return redisClient.Set(ctx, key, data, ttl).Err()
}

/* only for test below */

func Clear(ctx context.Context) error {
	return redisClient.FlushAll(ctx).Err()
}

func Close() {
	redisClient.Close()
	redisClient = nil
}
