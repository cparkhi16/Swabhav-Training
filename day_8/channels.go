package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/*	fmt.Println("In main")
		ch := make(chan int)
		go readData(ch)
		ch <- 90
		fmt.Println("End of main")*/
	fmt.Println("In main for communication between 2 go routines")
	wg.Add(1)
	ch := make(chan int)
	go read(ch)
	go write(ch)
	wg.Wait()

}
func read(ch chan int) {
	wg.Add(1)
	go readData(ch)
	wg.Wait()
	fmt.Println("Reading data from channel")
	d := <-ch
	fmt.Println("data is ", d)
	fmt.Println("End of read")
	wg.Done()
	fmt.Println("Done for read")
}
func readData(ch chan int) {
	fmt.Println("Reading data from channel call from read func")
	data := <-ch
	//time.Sleep(8 * time.Second)
	fmt.Println(data)
	fmt.Println("End of reading")
	wg.Done()
}
func write(ch chan int) {
	fmt.Println("Write data")
	ch <- 80
	fmt.Println("End of writing")
	wg.Done()
	fmt.Println("Done for write")
}
