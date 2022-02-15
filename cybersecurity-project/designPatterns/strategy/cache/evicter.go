package cache

type evicter interface {
	evict(c *Cache)
}
