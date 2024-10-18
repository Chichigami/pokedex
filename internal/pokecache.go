package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	pokemon map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		pokemon: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(k string, v []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.pokemon[k] = cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.pokemon[k]
	if !ok {
		return nil, false
	}
	return v.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.pokemon {
			if time.Since(v.createdAt) > interval {
				delete(c.pokemon, k)
			}
		}
		c.mu.Unlock()
	}
}
