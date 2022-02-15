package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start of main")
	ch := make(chan int, 2)
	wg.Add(2)
	go readFromChannel(ch)
	go writeToChannel(ch)
	wg.Wait()
	fmt.Println("end of main")
}

func readFromChannel(ch chan int) {
	fmt.Println("start of readFrom")
	data := <-ch
	data2 := <-ch
	data3 := <-ch
	data4 := <-ch
	fmt.Println(data, data2, data3, data4)
	wg.Done()
}

func writeToChannel(ch chan<- int) {
	fmt.Println("start of writeto")
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	ch <- 50
	ch <- 60
	//ch <- 70
	wg.Done()
}
