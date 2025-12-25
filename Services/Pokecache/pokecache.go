package Pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	mu       sync.RWMutex
	store    map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		interval: interval,
		store:    make(map[string]cacheEntry),
	}

	go cache.reapLoop()

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = cacheEntry{createdAt: time.Now(), val: val}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry := c.store[key]
	if entry.val == nil {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.store {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.store, key)
			}
		}
		c.mu.Unlock()
	}
}
