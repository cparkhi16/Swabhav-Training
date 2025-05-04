package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

// Writer function
func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock() // Lock for writing
	s.m[key] = value
	s.mu.Unlock()
}

// Reader function
func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.RLock() // Read lock (allows multiple readers)
	defer s.mu.RUnlock()
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

	// Reader Goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if val, exists := safeMap.Get("key1"); exists {
				fmt.Println("Read key1:", val)
			} else {
				fmt.Println("key1 not found")
			}
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish
}
