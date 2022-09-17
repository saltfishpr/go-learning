// @file: cache.go
// @date: 2021/12/22

// Package rediscache 提供redis缓存.
package rediscache

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var rc struct {
	instance *cache.Cache
	once     sync.Once
}

func New() *cache.Cache {
	rc.once.Do(func() {
		ring := redis.NewRing(&redis.RingOptions{
			Addrs: map[string]string{
				"server1": ":6379",
			},
		})
		rc.instance = cache.New(&cache.Options{
			Redis: ring,
		})
	})
	return rc.instance
}

func Set(key string, val interface{}, duration time.Duration) error {
	return New().Set(&cache.Item{Key: key, Value: val, TTL: duration})
}

func Get(key string, val interface{}) error {
	return New().Get(context.Background(), key, val)
}

func Del(key string) error {
	return New().Delete(context.Background(), key)
}
