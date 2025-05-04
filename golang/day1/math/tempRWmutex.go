package main

import (
	"fmt"
	"sync"
)

type RWMap struct {
	m  map[string]int
	mu sync.RWMutex
}

func (rm *RWMap) set(s string, i int) {
	rm.mu.Lock()
	rm.m[s] = i
	rm.mu.Unlock()
}

func (rm *RWMap) get(s string) (bool, int) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	val, ok := rm.m[s]
	return ok, val
}

func main() {
	safeMap := RWMap{m: make(map[string]int)}

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.set("k1", i)
			fmt.Println(" set k1 to ", i)
		}(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ok, v := safeMap.get("k1")
			if ok {
				fmt.Println(" value found", v)
			} else {
				fmt.Println(" value not found")
			}
		}()
	}
	wg.Wait()
}
