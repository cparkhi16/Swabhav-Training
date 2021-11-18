package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("start of main")
	wg.Add(5)
	go funcOne()
	go funcTwo(false)
	wg.Wait()
}
func funcOne() {
	for i := 0; i < 4; i++ {
		fmt.Println("calling funcTwo for time:", i+1)
		go funcTwo(true)
	}
	wg.Done()
}
func funcTwo(isMainCall bool) {
	for i := 0; i < 4; i++ {
		if isMainCall {
			fmt.Println("Func2 i from main call :", i)
		} else {
			fmt.Println("Func2 i from funcOne call :", i)
		}
	}
	wg.Done()
}
