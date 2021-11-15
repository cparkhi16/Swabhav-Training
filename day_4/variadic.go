package main

import (
	"fmt"
	"reflect"
)

func echo(strings ...string) {
	fmt.Println(reflect.TypeOf(strings))
	fmt.Println("Hi I am variadic")
	for _, s := range strings {
		fmt.Println(s)
	}
}

func main() {
	//strings := []string{"a", "b", "c"}
	//echo(strings...)
	echo() //Can pass a single final param as optional
}
