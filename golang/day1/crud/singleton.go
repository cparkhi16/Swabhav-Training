package main

import (
	"fmt"
	"sync"
)

type MySingleton struct {
}

var singleton *MySingleton
var once sync.Once

func getSingleton() *MySingleton {
	once.Do(func() {
		if singleton == nil {
			singleton = &MySingleton{}
			fmt.Println("Singleton instance created")
		}
	})
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
