package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start of main")
	ch := make(chan int, 3)
	ch2 := make(chan int)
	wg.Add(2)
	go readFromChannel(ch, ch2)
	go writeToChannel(ch, ch2)
	//ch <- 1close(ch)
	wg.Wait()
	fmt.Println("end of main")
}

func readFromChannel(ch chan int, ch2 chan int) {
	fmt.Println("start of readFrom")
	for v := range ch {
		select {
		case data := <-ch:
			fmt.Println("data from ch-", data, v)
		case data := <-ch2:
			fmt.Println("data from ch2-", data)
			wg.Done()
		default:
			fmt.Println("no data")
		}
	}
	//data := <-ch
	//data2 := <-ch
	//data3 := <-ch
	fmt.Println("end of read")
}

func writeToChannel(ch chan int, ch2 chan int) {
	fmt.Println("start of writeto")
	ch <- 20
	ch <- 10
	ch2 <- 64
	//ch <- 70
	close(ch)
	wg.Done()
}
