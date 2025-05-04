package main

import (
	"fmt"
	"sync"
)

func main() {
	var safeMap sync.Map
	var wg sync.WaitGroup

	// Writer Goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Store("key1", i)
			fmt.Println("Set key1 to", i)
		}(i)
	}

	// Reader Goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if val, exists := safeMap.Load("key1"); exists {
				fmt.Println("Read key1:", val)
			} else {
				fmt.Println("key1 not found")
			}
		}()
	}

	wg.Wait()
}
