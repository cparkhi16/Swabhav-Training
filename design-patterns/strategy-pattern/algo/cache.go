package algo

import "fmt"

type Cache struct {
	storage      map[string]string
	evictionAlgo Evicter
	capacity     int
	maxCapacity  int
	algo         string
	Queue        []string
	freq         map[string]int
	recentlyUsed []string
}

func InitCache() *Cache {
	storage := make(map[string]string)
	freq := make(map[string]int)
	return &Cache{
		storage:      storage,
		evictionAlgo: nil,
		capacity:     0,
		maxCapacity:  2,
		algo:         "",
		freq:         freq,
	}
}

func (c *Cache) SetEvictionAlgo(e Evicter) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	//fmt.Println(c.capacity)
	//fmt.Println(c.maxCapacity)
	if c.capacity >= c.maxCapacity {
		//fmt.Println("calling")
		c.evict()
	}
	c.Queue = append(c.Queue, key)
	c.capacity++
	c.storage[key] = value

	fmt.Println(c.storage)
}

func (c *Cache) Remove(key string) {
	delete(c.storage, key)
}
func (c *Cache) Get(key string) string {
	c.freq[key] = c.freq[key] + 1
	//fmt.Println("Cache ", c.freq)
	c.recentlyUsed = append(c.recentlyUsed, key)
	return c.storage[key]
}
func (c *Cache) evict() {
	c.evictionAlgo.Evict(c)
	c.capacity--
}
