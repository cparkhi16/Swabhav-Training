package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var wg = &sync.WaitGroup{}

type logger struct {
	name string
}

var singleInstance *logger

func GetLoggerInstance() *logger {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	if singleInstance == nil {
		fmt.Println("Creating single instance now.")
		singleInstance = &logger{name: "Chinmay"}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go GetLoggerInstance()
	}
	wg.Wait()
}
