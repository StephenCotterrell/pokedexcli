// Package pokecache provides a caching layer for the external pokeapi API
// to reduce latency and to avoid redundant network requests
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]CacheEntry
	mu           sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntries: make(map[string]CacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	copied := make([]byte, len(val))
	copy(copied, val)
	newCacheEntry := CacheEntry{createdAt: time.Now(), val: copied}
	c.cacheEntries[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntry, exists := c.cacheEntries[key]
	if exists {
		return cacheEntry.val, true
	} else {
		return []byte{}, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
