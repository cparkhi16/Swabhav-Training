package cache

import "fmt"

type Cache struct {
	storage        map[string]string
	evictionAlgo   evicter
	capacity       uint8          //remaining capacity of storage
	orderOfArrival []string       //fifo
	usageFrequency map[string]int //lfu
	recentlyUsed   []string       //lru
}

type EvictionAlog string

const (
	Fifo EvictionAlog = "fifo"
	Lru  EvictionAlog = "lru"
	Lfu  EvictionAlog = "lfu"
)

func NewCache(maxCapacity uint8) *Cache {
	return &Cache{
		storage:        make(map[string]string, maxCapacity),
		capacity:       maxCapacity,
		evictionAlgo:   lfu{},
		usageFrequency: make(map[string]int, maxCapacity),
	}

}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
}

func (c *Cache) Add(k, v string) {
	c.orderOfArrival = append(c.orderOfArrival, k)
	if c.capacity != 0 {
		c.storage[k] = v
		c.capacity = c.capacity - 1
	} else {
		c.evict()
		c.storage[k] = v
	}
	c.usageFrequency[k] = 0
}

func (c *Cache) DisplayStorage() {
	for k, v := range c.storage {
		fmt.Println(k, v, "usage Frequency-", c.usageFrequency[k])
	}
}

func (c *Cache) SetEvictionAlgo(newEvictionAlgo EvictionAlog) {
	switch newEvictionAlgo {
	case Fifo:
		c.evictionAlgo = fifo{}
	case Lru:
		c.evictionAlgo = lru{}
	case Lfu:
		c.evictionAlgo = lfu{}
	}
}

func (c *Cache) Get(k string) (string, error) {
	v, found := c.storage[k]
	if found {
		c.usageFrequency[k] = c.usageFrequency[k] + 1
		c.recentlyUsed = append(c.recentlyUsed, k)
		return v, nil
	}
	err := fmt.Errorf("element not found")
	return "", err
}

func (c *Cache) removeElementFromOrder(element string) {
	index := 0
	for i, v := range c.orderOfArrival {
		if v == element {
			index = i
			break
		}
	}
	c.orderOfArrival = append(c.orderOfArrival[:index], c.orderOfArrival[index+1:]...)
}
