// Copyright (c) 2012-2016 The Revel Framework Authors, All rights reserved.
// Revel Framework source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"time"

	"sync"

	"github.com/patrickmn/go-cache"
)

/***
Get(key string) interface{}
Set(key string, val interface{}, timeout time.Duration) error
IsExist(key string) bool
Delete(key string) error
*/
type InMemoryCache struct {
	cache cache.Cache  // Only expose the methods we want to make available
	mu    sync.RWMutex // For increment / decrement prevent reads and writes
}

func NewInMemoryCache(defaultExpiration time.Duration) InMemoryCache {
	return InMemoryCache{cache: *cache.New(defaultExpiration, time.Minute), mu: sync.RWMutex{}}
}

func (c InMemoryCache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.cache.Get(key)
	if !found {
		return nil
	}
	return value
}

func (c InMemoryCache) Set(key string, value interface{}, expires time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	// NOTE: go-cache understands the values of DefaultExpiryTime and ForEverNeverExpiry
	c.cache.Set(key, value, expires)
	return nil
}

// IsExist check value exists in memcache.
func (c InMemoryCache) IsExist(key string) bool {
	if _, ok := c.cache.Get(key); !ok {
		return false
	}
	return true
}

func (c InMemoryCache) Delete(key string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if _, found := c.cache.Get(key); !found {
		return ErrCacheMiss
	}
	c.cache.Delete(key)
	return nil
}
