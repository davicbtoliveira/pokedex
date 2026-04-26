package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.Entry[key]
	if !ok {
		return nil, false
	}

	return value.val, true
}
