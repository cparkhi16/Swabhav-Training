package main

import (
	"fmt"
)

//var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	go readChannel(ch)

	ch <- 10
}

func readChannel(ch chan int) {
	fmt.Println("yes")
	data := <-ch
	fmt.Println(data)
}
