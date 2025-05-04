package math

import "fmt"

func SelectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Hello world ch1"
	}()

	go func() {
		ch2 <- "Hello world ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case ms := <-ch2:
		fmt.Println(ms)
	}
}
