package cache

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"strings"
	"time"
)

var (
	rdb *redis.Client
)

var CacheMissedError = redis.Nil

const CacheKeySeparator = ":"

func Dial(config configx.Redis) error {
	ctx := context.Background()
	var tlsConfig *tls.Config
	if config.TLSEnabled { // In AWS if a Redis has a password, it must use TLS
		tlsConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:      config.Addr,
		Password:  config.Password,
		DB:        config.DB,
		TLSConfig: tlsConfig,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return err
	}

	return nil
}

func ConstructKey(keyParts ...string) string {
	return strings.Join(keyParts, CacheKeySeparator)
}

func Get(ctx context.Context, key string, data interface{}) error {
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(value), &data)
}

func Set(ctx context.Context, key string, data interface{}, expire time.Duration) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, key, value, expire).Err()
}

func HSet(ctx context.Context, key string, field string, data interface{}) error {

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if _, err := rdb.HSet(ctx, key, field, value).Result(); err != nil {
		return err
	}

	return nil
}

func Exists(ctx context.Context, key string) (bool, error) {
	exist, err := rdb.Exists(ctx, key).Result()

	if err != nil {
		return false, err
	}

	return exist == 1, nil
}
