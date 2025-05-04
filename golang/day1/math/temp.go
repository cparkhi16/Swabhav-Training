package math

import "fmt"

func f1(ch chan int, i int) {
	ch <- i
}
func f2(ch chan int, i int) {
	ch <- i
}
func Temp() {
	ch := make(chan int, 2)
	go f1(ch, 1)
	go f2(ch, 2)
	//var i int
	i := <-ch
	//var j int
	j := <-ch
	fmt.Println(" sum is ", i+j)
}
