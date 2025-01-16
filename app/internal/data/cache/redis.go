package cache

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheClient interface {
	CheckCache() bool
	Set(key string, item interface{}, expiration time.Duration) error
	Get(key string) ([]byte, error)
	Delete(key string) (*int64, error)
}

type cacheClient struct {
	cache   *redis.Client
	context context.Context
}

func New() *cacheClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_ADDRESS"),
		Password: os.Getenv("CACHE_PASSWORD"),
		DB:       0,
	})

	return &cacheClient{
		context: context.Background(),
		cache:   rdb,
	}
}

func (c *cacheClient) CheckCache() bool {
	defer c.cache.Close()

	_, err := c.cache.Ping(c.context).Result()

	if err != nil {
		return false
	}
	return true
}

func (c *cacheClient) Set(key string, item interface{}, expiration time.Duration) error {
	return c.cache.Set(c.context, key, item, expiration).Err()
}

func (c *cacheClient) Get(key string) ([]byte, error) {
	return c.cache.Get(c.context, key).Bytes()
}

func (c *cacheClient) Delete(key string) (*int64, error) {
	result, err := c.cache.Del(c.context, key).Result()

	if err != nil {
		return nil, err
	}

	return &result, nil
}
