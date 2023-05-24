package store

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/MustansirZia/markdown-visitor-badge/configProvider"
	"github.com/redis/go-redis/v9"
)

// Counter - Interface for the counter type store.
type Counter interface {
	// IncrementAndGet - Increments the value of the counter with the given key and returns the new value.
	IncrementAndGet(ctx context.Context, key string) (uint64, error)
}

// NewCounter - Constructs and returns a new Counter.
func NewCounter(config configProvider.Config) Counter {
	var tlsConfig *tls.Config
	if config.RedisUseTLS {
		tlsConfig = &tls.Config{}
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:      fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Username:  config.RedisUsername,
		Password:  config.RedisPassword,
		TLSConfig: tlsConfig,
		DB:        0,
		PoolSize:  1,
	})
	return &redisBasedCounter{redisClient: redisClient}
}

type redisBasedCounter struct {
	redisClient *redis.Client
}

func (c *redisBasedCounter) IncrementAndGet(ctx context.Context, key string) (uint64, error) {
	if value, err := c.redisClient.Incr(ctx, key).Result(); err != nil {
		return 0, err
	} else {
		return uint64(value), nil
	}
}
