package main

import (
	"fmt"
	"singleton/log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		log1 := log.GetInstance("here is log1")
		fmt.Println(log1.GetMsg())
		wg.Done()
	}()
	go func() {
		log2 := log.GetInstance("here is log2")
		fmt.Println(log2.GetMsg())
		wg.Done()
	}()
	wg.Wait()

}
