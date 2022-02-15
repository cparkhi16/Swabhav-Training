package main

import "fmt"

func main() {
	var num int = 10
	fmt.Println("address of num from main", &num)
	increment(num)
	fmt.Println("value of num from main after increment fn call", num)
	incrementWithPointer(&num)
	fmt.Println("value of num from main after incrementWithPointer fn call", num)
}

func increment(num int) {
	fmt.Println("address of num from increment fn", &num)
	num = num + 20
	fmt.Println("value of num from increment fn", num)
}

func incrementWithPointer(num *int) {
	fmt.Println("address of num from incrementWithPointer fn", &num)
	fmt.Println("value of num from incrementWithPointer fn", num)
	*num = *num + 20
}
