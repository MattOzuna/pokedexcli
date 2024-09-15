package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu           *sync.Mutex
	cacheEntries map[string]CacheEntry
	duration     time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) Cache {
	ticker := time.NewTicker(duration)
	newCache := Cache{
		mu:           &sync.Mutex{},
		cacheEntries: map[string]CacheEntry{},
		duration:     duration,
	}
	go newCache.readLoop(ticker.C)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheEntries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return entry.val, ok
}

func (c *Cache) readLoop(interval <-chan time.Time) {
	defer c.mu.Unlock()
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
