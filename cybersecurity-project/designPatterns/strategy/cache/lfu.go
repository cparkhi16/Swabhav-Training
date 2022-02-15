package cache

import "fmt"

type lfu struct {
}

func (f lfu) evict(c *Cache) {
	fmt.Println("lfu")
	//to find minimum frequency of usage
	min := 10000
	leastFrequentlyUsedElement := ""

	for k, v := range c.usageFrequency {
		if v <= min {
			leastFrequentlyUsedElement = k
			min = v
		}
	}
	//if there is no such element or usageFrequency map is empty then we use fifo
	if leastFrequentlyUsedElement == "" {
		leastFrequentlyUsedElement = c.orderOfArrival[0]
	}
	fmt.Println("leastFrequentlyUsedElement- ", leastFrequentlyUsedElement)
	delete(c.storage, leastFrequentlyUsedElement)
	c.removeElementFromOrder(leastFrequentlyUsedElement)
	delete(c.usageFrequency, leastFrequentlyUsedElement)
}
