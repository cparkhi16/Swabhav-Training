package main

import "fmt"

func main() {
	if true {
		fmt.Println("true")
	}
	if false {
		fmt.Println("false")
	}
	if !false {
		fmt.Println("!false")
	}
	if true {
		fmt.Println("if true")
	} else {
		fmt.Println("else")
	}
	if false {
		fmt.Println("if false")
	} else if true {
		fmt.Println("else if true")
	}
	if 12 == 12 {
		fmt.Println("12==12")
	}
	if 12 != 12 {
		fmt.Println("12!=12")
	}
	if 12 > 12 {
		fmt.Println("12>12")
	}
	if 12 >= 12 {
		fmt.Println("12>=12")
	}
	if 12 == 12 && 5.9 == 5.9 {
		fmt.Println("12==12 && 5.9==5.9")
	}
	if 12 == 12 || 5.9 == 6.4 {
		fmt.Println("12==12 || 5.9==6.4")
	}

}
