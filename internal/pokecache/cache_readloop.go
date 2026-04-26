package pokecache

import "time"

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		for k, v := range c.Entry {
			if time.Since(v.createdAt) > interval {
				delete(c.Entry, k)
			}
		}
	}
}
