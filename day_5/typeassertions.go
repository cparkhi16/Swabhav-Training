package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{} = "Good morning"
	fmt.Println("type of i", reflect.TypeOf(i))

	i = 1299
	fmt.Println("Type of i", reflect.TypeOf(i))
	j := &i
	fmt.Println("Type of j", reflect.TypeOf(j))
	fmt.Println("Value of j ", *j)
	switch i.(type) {
	case string:
		fmt.Println("I am string")
	case int:
		fmt.Println("I am int")
	case bool:
		fmt.Println("I am bool")
	default:
		fmt.Println("None of the above")
	}
}
