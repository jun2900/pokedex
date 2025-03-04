package internal

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]cacheEntry)}

	cache.reapLoop(interval)
	return cache
}

type Cache struct {
	mu    sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheKey, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return cacheKey.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			c.mu.Lock()
			for k, v := range c.cache {
				if time.Since(v.createdAt) > interval {
					delete(c.cache, k)
				}
			}
			c.mu.Unlock()
		}
	}()
}
