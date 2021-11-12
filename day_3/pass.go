package main

import "fmt"

func main() {
	var num int = 80
	var num1 int = 90
	decrementByOne(&num)
	p := &num1
	ptr := &p
	fmt.Println("After decrementing by 1", num)
	fmt.Println("Address of num1 in main", &num1)
	decrementByTwo(*ptr)
	fmt.Println("After decrementing by 2", num1)
}
func decrementByOne(num *int) {
	*num = *num - 1
}
func decrementByTwo(num *int) {
	fmt.Println("Address of num1 in 2nd func", num)
	*num = *num - 2
}
