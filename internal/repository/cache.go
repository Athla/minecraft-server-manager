package repository

import (
	"mine-server-manager/internal/internalErrors"
	"sync"
	"time"
)

type CacheRepository interface {
	Add(key string, value any, ttl time.Duration) error
	Get(key string) (any, error)
}

type cacheItem struct {
	value      any
	expiration int64
}

type InMemoryCache struct {
	items map[string]cacheItem
	mu    sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		items: make(map[string]cacheItem),
	}
}

func (c *InMemoryCache) Add(key string, value any, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	exp := time.Now().Add(ttl).Unix()
	c.items[key] = cacheItem{
		value:      value,
		expiration: exp,
	}

	return nil
}

func (c *InMemoryCache) Get(key string) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, internalErrors.ErrCacheItemNotFound
	}

	if item.expiration < time.Now().Unix() {
		delete(c.items, key)
		return nil, internalErrors.ErrCacheItemExpired
	}

	return item.value, nil
}
