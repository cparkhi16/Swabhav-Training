package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	s.m[key] = value
	s.mu.Unlock()
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, exists := s.m[key]
	return val, exists
}

func main() {
	safeMap := SafeMap{m: make(map[string]int)}
	var wg sync.WaitGroup

	// Writer Goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Set("key1", i)
			fmt.Println("Set key1 to", i)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish

	// Read the final value
	if val, exists := safeMap.Get("key1"); exists {
		fmt.Println("Final value of key1:", val)
	} else {
		fmt.Println("key1 not found")
	}
}
