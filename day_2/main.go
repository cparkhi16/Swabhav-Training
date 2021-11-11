package main

import (
	"fmt"
	"reflect"
	"time"
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

	t := time.Now()
	Current_Hour := t.Hour()
	if Current_Hour >= 8 && Current_Hour <= 11 {
		fmt.Println("Good Morning !")
	} else if Current_Hour >= 12 && Current_Hour <= 15 {
		fmt.Println("Good Afternoon ! ")
	} else if Current_Hour >= 16 && Current_Hour <= 20 {
		fmt.Println("Good Evening !")
	} else {
		fmt.Println("Good night !")
	}
	var i, j int
OUTERFOR:
	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {
			fmt.Println(j)
			if j == 5 {
				break OUTERFOR
			}
		}

	}
	fmt.Println("OUTSIDE OUTER FOR ")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Yaayy Today is weekend")
	default:
		fmt.Println("Today is working day")
	}
	//add := addition
	a := float32(10)
	b := float32(20)

	result := mathOperations(a, b, addition)
	fmt.Println("Addition is", result, "Type of result ", reflect.TypeOf(result))

	diff := mathOperations(a, b, subtraction)
	fmt.Println("Addition is", diff, "Type of result ", reflect.TypeOf(diff))

	mul := mathOperations(a, b, multiplication)
	fmt.Println("Addition is", mul, "Type of result ", reflect.TypeOf(mul))

	div := mathOperations(a, b, division)
	fmt.Println("Addition is", div, "Type of result ", reflect.TypeOf(div))

}
func mathOperations(f1, f2 float32, f func(num1, num2 float32) float32) float64 {
	return float64(f(f1, f2))
}
func addition(n1, n2 float32) float32 {
	ans := n1 + n2
	return ans
}
func subtraction(n1, n2 float32) float32 {
	ans := n1 - n2
	return ans
}
func multiplication(n1, n2 float32) float32 {
	ans := n1 * n2
	return ans
}
func division(n1, n2 float32) float32 {
	ans := n1 / n2
	return ans
}
