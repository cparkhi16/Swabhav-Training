package main

import "fmt"

func main() {
	var i int = 10
	defer test(&i)
	i = 20
	fmt.Println("i in main", i)
	var j int = 11
	defer testing(j)
	j = 12
	fmt.Println("j in main", j)
	defer func() {
		fmt.Println("Hi in IIFE")
		val := recover()
		fmt.Println("inside IIFE ", val)
		defer t()
		defer t2()
	}()
	if j == 12 {
		panic("\nError occured ")
	}
	/*defer func() {
		fmt.Println("Hi in IIFE")
		//val := recover()
		//fmt.Println("inside IIFE ", val)
	}()*/ //panic checks defer first after panic so it will not Run
}
func test(i *int) {
	fmt.Println("Val of i in defer func", *i)
}
func testing(j int) {
	fmt.Println("Val of j in defer ", j)
}
func t() {
	fmt.Println("I will be called after t2 ")
}
func t2() {
	fmt.Println("First I will be called then t will be called")
}
