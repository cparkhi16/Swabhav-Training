package algo

import (
	"fmt"
)

type Lru struct {
}

func (l *Lru) Evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
	lruMap := make(map[string]int)
	visited := make(map[string]bool)
	for _, val := range c.recentlyUsed {
		if !visited[val] {
			lruMap[val] = 1
			visited[val] = true
		} else {
			lruMap[val] = lruMap[val] + 1
		}
	}
	//fmt.Println(lruElem)
	//fmt.Println(lruMap)
	min := 500
	lruElement := ""
	for k, v := range lruMap {
		if v < min {
			lruElement = k
			min = v
		}
	}
	//fmt.Println(lruElement)
	c.Remove(lruElement)
}
