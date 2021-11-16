package main

import "fmt"

func main() {
	numbers := []int{3, 16, -2, 10, 23, 12}
	for i, number := range numbers {
		if number >= 10 && number <= 20 {
			fmt.Println(i, number)
		}
	}
	var number [3]int
	number[0] = 42
	number[2] = 108
	var letters = [3]string{"a", "b", "c"}
	fmt.Println(number[0])
	fmt.Println(number[1])
	fmt.Println(number[2])
	fmt.Println(letters[2])
	fmt.Println(letters[0])
	fmt.Println(letters[1])
}
