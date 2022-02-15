package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var c int

func main() {
	fmt.Println("start of main")
	wg.Add(1)
	c = 0
	go ConcIteration1()
	go ConcIteration2()
	go ConcIteration3()
	wg.Wait()
	fmt.Println("end of main")
}

func ConcIteration1() {
	c++
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("ConcIteration1-", i)
	}
	wg.Done()
}

func ConcIteration2() {
	c--
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("ConcIteration2-", i)
	}
	wg.Done()
}

func ConcIteration3() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println("ConcIteration3-", i)
	}
	wg.Done()
}
