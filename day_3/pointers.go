package main

import "fmt"

func addition(n1, n2 float32) float32 {
	ans := n1 + n2
	return ans
}
func main() {
	f := a()
	ans := f(float32(10), float32(30), addition)
	fmt.Println(ans)
	var a string
	var b int
	var p *int
	var c bool
	ptr := &b
	fmt.Println("Zero value of a ", a)
	fmt.Println("Zero value of c", c)
	fmt.Println("Zero value of b ", b)
	fmt.Println("Zero value of p ", p)
	//fmt.Println("Zero value of *p ", *p)
	fmt.Println("Value of ptr ", ptr)
	fmt.Println("Value of a", a)
	fmt.Println("Value of *ptr", *ptr)
	ptr2 := &ptr
	fmt.Println("Value of ptr2 ", ptr2)
	fmt.Println("Value of *ptr", *ptr2)
	fmt.Println("Value of **ptr2 ", **ptr2)
}
func a() func(f1, f2 float32, f func(num1, num2 float32) float32) float64 {
	return func(f1, f2 float32, f func(num1, num2 float32) float32) float64 {
		return float64(f(f1, f2))
	}
}
