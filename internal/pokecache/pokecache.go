package pokecache

import "time"

type cacheEntry struct {
	value     []byte
	createdAt time.Time
}

type Cache struct {
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{make(map[string]cacheEntry)}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		value:     val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cache, ok := c.cache[key]

	return cache.value, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	expired := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(expired) {
			delete(c.cache, k)
		}
	}
}
