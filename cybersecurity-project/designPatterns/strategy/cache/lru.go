package cache

import "fmt"

type lru struct {
}

func (f lru) evict(c *Cache) {
	fmt.Println("lru")
	leastRecentlyUsedElement := ""
	//make a dummy map with all elements from storage
	elements := make(map[string]int, len(c.storage))
	for k, _ := range c.storage {
		elements[k] = 0
	}
	fmt.Println(c.recentlyUsed)
	//if recentlyUsed array is empty then we use fifo
	if len(c.recentlyUsed) == 0 {
		leastRecentlyUsedElement = c.orderOfArrival[0]
	} else {
		//delete recently used elements from dummy map
		for i := len(c.recentlyUsed) - 1; i > 0; i-- {
			if len(elements) > 1 {
				delete(elements, c.recentlyUsed[i])
			}
		}
		//Now only least recently used element will remain
		for k, _ := range elements {
			leastRecentlyUsedElement = k
		}
	}

	fmt.Println("leastRecentlyUsedElement-", leastRecentlyUsedElement)
	delete(c.storage, leastRecentlyUsedElement)
	delete(c.usageFrequency, leastRecentlyUsedElement)
	c.removeElementFromOrder(leastRecentlyUsedElement)
}
