package main

import (
	math "example/math"
	"fmt"
	"reflect"

	greeting "github.com/headfirstgo/greeting"
)

func main() {
	num1 := 2
	num2 := 3
	fmt.Println(math.AddTwoNumbers(num1, num2))
	str1 := "abc"
	float1 := 2.5
	bool1 := true
	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(reflect.TypeOf(str1))
	fmt.Println(reflect.TypeOf(float1))
	fmt.Println(reflect.TypeOf(bool1))
	greeting.Hello()
	greeting.Hi()
}
