package main

import (
	"fmt"
	"reflect"
	"training/test/math"

	test "github.com/headfirstgo/greeting"
)

func main() {
	fmt.Println("hello world")
	ans := math.AddTwoNumbers(1, 2)
	fmt.Println("Addition of (1+2) numbers is ", ans)
	fmt.Println("Type of ans is ", reflect.TypeOf(ans))
	MyName := "chinmay"
	fmt.Println("Type of myname is ", reflect.TypeOf(MyName))
	MyName = MyName + "-123"
	fmt.Println("Type of updates myname ", reflect.TypeOf(MyName))
	IsHealthy := true
	fmt.Printf("Type of isHealthy %t", IsHealthy)
	TestNum := 189.278008
	fmt.Printf("\nType of test_num is %t", TestNum)
	fmt.Println("\nType of test_num using reflect ", reflect.TypeOf(TestNum))
	test.Hello()

}
