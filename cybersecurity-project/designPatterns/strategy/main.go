package main

import (
	"disk/cache"
	"fmt"
)

func main() {
	mobileCache := cache.NewCache(3)
	mobileCache.Add("a", "1")
	mobileCache.Add("b", "2")
	mobileCache.Add("c", "3")
	v1, _ := mobileCache.Get("a")
	fmt.Println(v1)
	v2, _ := mobileCache.Get("b")
	fmt.Println(v2)
	v3, _ := mobileCache.Get("a")
	fmt.Println(v3)
	v4, _ := mobileCache.Get("c")
	fmt.Println(v4)
	mobileCache.DisplayStorage()
	mobileCache.Add("d", "4")
	fmt.Println("after adding d")
	mobileCache.DisplayStorage()
	//mobileCache.SetEvictionAlgo(cache.Fifo)
	mobileCache.Add("e", "5")
	fmt.Println("after adding e")
	mobileCache.DisplayStorage()
	v5, _ := mobileCache.Get("e")
	fmt.Println(v5)
	//mobileCache.SetEvictionAlgo(cache.Lfu)
	mobileCache.DisplayStorage()
	mobileCache.Add("f", "6")
	fmt.Println("after adding f")
	mobileCache.DisplayStorage()

}
