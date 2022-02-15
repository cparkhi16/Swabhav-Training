package main

import "fmt"

func main() {
	numbers := []int{3, 16, -2, 10, 23, 12}
	for i, number := range numbers {
		if number >= 10 && number <= 20 {
			fmt.Println(i, number)
		}
	}
}
