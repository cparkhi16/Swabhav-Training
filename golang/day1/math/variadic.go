package math

import "fmt"

func VariadicArg(value ...int) int {
	sum := 0
	for _, v := range value {
		sum += v
	}
	return sum
}

func Closure(i int) func(i int) {
	f := func(i int) {
		fmt.Println(" hii in closure", i)
	}
	return f
}

func Sum() func(int) int {
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}
}

func Switch(i int) bool {
	var b bool = false
	switch i {
	case 1:
		b = true
	case 2:
		b = false
	}
	return b
}

func ForLoop(i int) {
	var k int
	for k = 0; k < i; k++ {
		fmt.Println(" for loop  ", k)
	}
}

func WhileLoop(i int) {
	var k int
	for k < i {
		fmt.Println(" While lop emulation ", k)
		k++
	}
}

func DoMathOps(x int, y int, f func(int, int) int) int {
	return f(x, y)
}

func Addition(x int, y int) int {
	var res int = x + y
	return res
}

func Subs(x int, y int) int {
	var res int = x - y
	return res
}
