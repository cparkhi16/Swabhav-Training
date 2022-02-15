package main

import "fmt"

type Number int

func (n Number) Add(otherNum int) {
	fmt.Println(n, "plus", otherNum, "is", int(n)+otherNum)
}
func (n Number) Subtract(otherNum int) {
	fmt.Println(n, "minus", otherNum, "is", int(n)-otherNum)
}

func main() {
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}
