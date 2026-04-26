package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		Entry: make(map[string]cacheEntry),
		mu:    sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return &cache
}
