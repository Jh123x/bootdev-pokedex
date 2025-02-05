package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	cache map[string]cacheEntry
	mutex sync.RWMutex
}

func NewCache() *Cache {
	cache := &Cache{
		cache: make(map[string]cacheEntry, 100),
		mutex: sync.RWMutex{},
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) reapLoop() {
	// TODO
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	v, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	return v.val, true
}
