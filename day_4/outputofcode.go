package main

import "fmt"

func main() {
	nums := make([]int, 0, 4)
	nums = append(nums, 1, 2, 3)

	a := append(nums, 4)
	fmt.Println(nums)
	fmt.Println("Address of a[3]: before b ", &a[3])
	fmt.Println("a is : before b", a)
	b := append(nums, 5)
	fmt.Println("Address of b[3]: after b ", &b[3])
	fmt.Println("b is :", b)
	fmt.Println("a is :", a)
}
