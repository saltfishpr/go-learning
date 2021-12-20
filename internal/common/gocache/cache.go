// @description:
// @file: cache.go
// @date: 2021/12/8

package gocache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	GeneralExpiration = 15 * time.Minute
	NoExpiration      = cache.NoExpiration
)

var c = cache.New(5*time.Minute, 10*time.Minute)

func Set(key string, val interface{}, duration time.Duration) {
	c.Set(key, val, duration)
}

func Get(key string) (interface{}, bool) {
	return c.Get(key)
}

func Del(key string) {
	c.Delete(key)
}
