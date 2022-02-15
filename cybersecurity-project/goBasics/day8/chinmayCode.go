package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("In main")
	ch := make(chan int, 2)
	ch2 := make(chan string, 2)
	wg.Add(2)
	go ReadData(ch, ch2)
	go Write(ch)
	ch <- 90
	ch <- 100
	ch2 <- "hi"
	ch2 <- "Hello"
	wg.Wait()
	s := <-ch
	fmt.Println("reading from closed channel1 in main", s)
	k := <-ch2
	fmt.Println("reading from closed channel2 in main", k)
	fmt.Println("End of main")

}
func ReadData(ch chan int, ch2 chan string) {
	fmt.Println("Reading data from channel")
	//time.Sleep(1 * time.Second)
	select {
	case v := <-ch:
		fmt.Println("data from ch", v)
		for d := range ch {
			fmt.Println("Data from channel is ", d)
		}
	case <-ch2:
		for d := range ch2 {
			fmt.Println("Data from channel2 is ", d)
		}
	default:
		fmt.Println("No data from both channels")
	}
	wg.Done()
	fmt.Println("End of reading")
}
func Write(ch chan int) {
	fmt.Println("Writing data to channel")
	ch <- 900
	close(ch)
	wg.Done()
	fmt.Println("Writing done")
}
