package main

import (
	"fmt"
	"sync"
	"time"
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
	//ch2 <- "bye"
	//close(ch2)
	//ch <- 1000
	wg.Wait()
	//Writing to a closed channel
	//ch <- 87 //Panic send on closed channel
	//closing a channel once again
	//close(ch) // Panic close of closed channel
	//Reading from closed channel in main
	s := <-ch
	fmt.Println("reading from closed channel1 in main", s)
	k := <-ch2
	fmt.Println("reading from closed channel2 in main", k)
	fmt.Println("End of main")

}
func ReadData(ch chan int, ch2 chan string) {
	fmt.Println("Reading data from channel")
	/*data := <-ch
	d := <-ch
	dNew := <-ch*/
	/*for v := range ch {
		fmt.Println("Data from channel is ", v)
	}
	for v := range ch2 {
		fmt.Println("Data from channel2 is ", v)
	}*/
	//Reading from closed channel
	//f := <-ch
	//fmt.Println("reading from closed channel", f)
	//time.Sleep(8 * time.Second)
	/*fmt.Println("data 1 from ch", data)
	fmt.Println("data2 from ch ", d)
	fmt.Println("data 3 from ch", dNew)*/

	// select statement
	select {
	case d := <-ch:
		fmt.Println("data from ch", d)
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
	time.Sleep(1 * time.Second)
	fmt.Println("Writing data to channel")
	ch <- 900
	//ch <- 1000
	//ch <- 10000
	wg.Done()
	close(ch)
	fmt.Println("Writing done")
}
