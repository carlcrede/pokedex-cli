package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	key       string
	Val       []byte
}

type Cache struct {
	mu       sync.Mutex
	interval time.Duration
	entries  map[string]CacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		entries:  make(map[string]CacheEntry),
	}
	cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) (CacheEntry, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.entries[key]
	return value, ok
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	ticker := time.NewTicker(c.interval)
	go func() {
		for range ticker.C {
			c.mu.Lock()
			for key, entry := range c.entries {
				if time.Since(entry.createdAt) > c.interval {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
