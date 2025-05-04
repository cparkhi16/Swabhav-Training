package main

import (
	"fmt"
	"sync"
)

type MySingleton struct {
}

var singleton *MySingleton
var m sync.Mutex

func getSingleton() *MySingleton {
	if singleton == nil {
		m.Lock()
		if singleton == nil {
			singleton = &MySingleton{}
			fmt.Println(" singleton instance created ")
		}
		m.Unlock()
	}
	return singleton
}

func (m *MySingleton) log() {
	fmt.Println(" log from singleton")
}
func main() {
	// var a = getSingleton()
	// a.log()
	// var b = getSingleton()
	// b.log()
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			var a = getSingleton()
			a.log()
			wg.Done()
		}()
	}
	wg.Wait()
}
