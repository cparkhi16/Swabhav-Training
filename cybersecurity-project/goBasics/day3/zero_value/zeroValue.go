package main

import "fmt"

func main() {
	var a int
	var b string
	var c bool
	var d float64
	var p *int
	fmt.Println("value of int a", a)
	fmt.Println("value of int b", b)
	fmt.Println("value of int c", c)
	fmt.Println("value of int d", d)
	fmt.Println("value of *int p", p)
	fmt.Println("deref of *int p", *p)
}
