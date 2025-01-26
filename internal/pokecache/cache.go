package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry    map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Create a new cache
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entry:    make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
		interval: interval,
	}

	go func() {
		for {
			time.Sleep(interval)
			cache.reapLoop()
		}
	}()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	output, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return output.val, true
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for key, val := range c.entry {
		if now.Sub(val.createdAt) > c.interval {
			delete(c.entry, key)
		}
	}
}
