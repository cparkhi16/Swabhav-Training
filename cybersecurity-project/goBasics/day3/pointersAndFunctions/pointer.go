package main

import "fmt"

func main() {
	var a int = 13
	fmt.Println("value of a", a)
	fmt.Println("address of a", &a)

	var p *int = &a
	fmt.Println("value of p", p)
	fmt.Println("address of p", &p)
	fmt.Println("deref of p", *p)

	var pp **int = &p
	fmt.Println("value of pp", pp)
	fmt.Println("address of pp", &pp)
	fmt.Println("deref of pp", *pp)
	fmt.Println("double deref of pp", **pp)

}
