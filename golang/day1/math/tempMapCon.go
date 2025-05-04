package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	m  map[string]int
	mu sync.Mutex
}

func (mp *MyMap) set(s string, i int) {
	mp.mu.Lock()
	mp.m[s] = i
	mp.mu.Unlock()
}

func (mp *MyMap) get(s string) (bool, int) {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	val, ok := mp.m[s]
	return ok, val
}

func main() {
	safeMap := MyMap{m: make(map[string]int)}

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.set("k1", i)
			fmt.Println("set k1 to ", i)
		}(i)
	}
	wg.Wait()
	ok, val := safeMap.get("k1")
	if ok {
		fmt.Println("value exists in map ", val)
	} else {
		fmt.Println(" val doesn't exist")
	}

}
