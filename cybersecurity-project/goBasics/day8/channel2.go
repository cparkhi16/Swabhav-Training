package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start of main")
	ch := make(chan int)
	wg.Add(4)
	go readFromChannel(ch)
	go writeToChannel(ch)
	go readAndWriteFromChannel(ch)
	wg.Wait()
	fmt.Println("end of main")
}

func readFromChannel(ch chan int) {
	fmt.Println("start of readFrom")
	data := <-ch
	fmt.Println(data)
	wg.Done()
}

func writeToChannel(ch chan int) {
	fmt.Println("start of writeto")
	ch <- 10
	wg.Done()
}

func readAndWriteFromChannel(ch chan int) {
	fmt.Println("read and write")
	go readFromChannel(ch)
	ch <- 20
	wg.Done()
}
