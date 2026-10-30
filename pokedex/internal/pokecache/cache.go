package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(wait time.Duration) *Cache {
	newMap := make(map[string]cacheEntry)
	c := Cache{
		cache:    newMap,
		interval: wait,
	}
	go c.reapLoop()
	return &c
}

func (c *Cache) Add(key string, info []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       info,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if info, ok := c.cache[key]; ok {
		return info.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop() {
	tick := time.NewTicker(c.interval)
	for range tick.C {
		c.mu.Lock()
		for key, value := range c.cache {
			if time.Since(value.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
