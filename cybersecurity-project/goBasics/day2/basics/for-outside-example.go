package main

import (
	"fmt"
)

func main() {
	num1 := 10

	for {
		for {
			fmt.Println("go outside")
			goto outsideFor
		}
		fmt.Println("inside")
	}
outsideFor:

	fmt.Println(num1)
}
