package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start of main")
	ch := make(chan int, 3)
	wg.Add(2)
	go readFromChannel(ch)
	go writeToChannel(ch)
	ch <- 100
	ch <- 200
	wg.Wait()
	fmt.Println("end of main")
}

func readFromChannel(ch chan int) {
	fmt.Println("start of readFrom")

	for v := range ch {
		fmt.Println("Inside read-", v)
	}
	//data := <-ch
	//data2 := <-ch
	//data3 := <-ch
	//fmt.Println(data, data2, data3)
	wg.Done()
}

func writeToChannel(ch chan int) {
	fmt.Println("start of writeto")
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	//ch <- 70
	close(ch)
	wg.Done()
}
