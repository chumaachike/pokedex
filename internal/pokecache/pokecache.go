package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	values   map[string]CacheEntry
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		values:   map[string]CacheEntry{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cachEntry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.values[key] = cachEntry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.values[key]
	if !ok {
		return nil, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.values {
			if time.Since(v.createdAt) > c.interval {
				delete(c.values, k)
			}
		}
		c.mu.Unlock()
	}

}
