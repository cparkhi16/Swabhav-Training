package algo

import (
	"fmt"
)

type Lru struct {
}

func (l *Lru) Evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
	lruMap := make(map[string]string)
	for k, val := range c.storage {
		lruMap[k] = val
	}
	//fmt.Println("--", lruMap)
	lruElement := ""
	fmt.Println("Queue", c.queue)
	if len(c.recentlyUsed) != 0 {
		for i := len(c.recentlyUsed) - 1; i >= 0; i-- {
			if len(lruMap) == 1 {
				break
			}
			delete(lruMap, c.recentlyUsed[i])
		}
		for k, _ := range lruMap {
			lruElement = k
		}
	} else {
		lruElement = c.queue[0]
		c.queue = c.queue[1:]
	}
	//fmt.Println(lruElement)
	//fmt.Println(c.recentlyUsed)
	c.Remove(lruElement)
}
