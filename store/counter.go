package store

import (
	"context"
	"fmt"

	"github.com/MustansirZia/markdown-visitor-badge/configProvider"
	"github.com/redis/go-redis/v9"
)

// Counter - Interface for the counter type store.
type Counter interface {
	// IncrementAndGet - Increments the value of the counter with the given key and returns the new value.
	IncrementAndGet(ctx context.Context, key string) (uint64, error)
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

// NewCounter - Constructs and returns a new Counter.
func NewCounter(config configProvider.Config) Counter {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Username: config.RedisUsername,
		Password: config.RedisPassword,
		DB:       int(config.RedisDatabase),
		PoolSize: 1,
	})
	return &redisBasedCounter{redisClient: redisClient}
}
