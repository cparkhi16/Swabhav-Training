package algo

type Evicter interface {
	Evict(c *Cache)
}
