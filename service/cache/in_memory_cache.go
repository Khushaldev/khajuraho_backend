package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      string
	Expiration time.Time
}

type InMemoryCache struct {
	data  map[string]CacheItem
	mutex sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]CacheItem),
	}
}

func (c *InMemoryCache) Set(key, value string, expiration time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(expiration),
	}
	return nil
}

func (c *InMemoryCache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, exists := c.data[key]
	if !exists {
		return "", false
	}

	// Check if the item has expired
	if time.Now().After(item.Expiration) {
		delete(c.data, key)
		return "", false
	}

	return item.Value, true
}
