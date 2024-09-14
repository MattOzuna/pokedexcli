package pokecache

import (
	"sync"
	"time"
)

type cache struct {
	mu           *sync.Mutex
	cacheEntries map[string]cacheEntry
	duration     time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) cache {
	ticker := time.NewTicker(duration)
	newCache := cache{
		mu:           &sync.Mutex{},
		cacheEntries: map[string]cacheEntry{},
		duration:     duration,
	}
	go newCache.readLoop(ticker.C)
	return newCache
}

func (c *cache) Add(key string, val []byte) {
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheEntries[key] = newEntry
}

func (c *cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return entry.val, ok
}

func (c *cache) readLoop(interval <-chan time.Time) {
	for range interval {
		c.mu.Lock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > c.duration {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
