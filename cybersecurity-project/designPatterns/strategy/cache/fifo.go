package cache

import "fmt"

type fifo struct {
}

func (f fifo) evict(c *Cache) {
	fmt.Println("fifo")
	firstInElementKey := c.orderOfArrival[0]
	delete(c.storage, firstInElementKey)
	delete(c.usageFrequency, firstInElementKey)
	c.orderOfArrival = c.orderOfArrival[1:]
}
