package main

import c "strategy/algo"

func main() {
	//fifo := c.NewFIFO()
	//lfu := &c.Lfu{}
	cache := c.InitCache()
	//cache.SetEvictionAlgo(fifo)
	//lru := &c.Lru{}
	//cache.SetEvictionAlgo(lru)
	lfu := &c.Lfu{}
	cache.SetEvictionAlgo(lfu)
	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Get("a")
	cache.Get("b")
	cache.Get("a")
	cache.Get("b")
	cache.Add("c", "3")
	cache.Add("d", "5")
	cache.Add("e", "9")
}
