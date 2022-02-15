//IIFE-->Immediately Invoked Function Expressions
package main

import "fmt"

func main() {
	var num1 float32 = 4.5
	var num2 float32 = 3.2
	fmt.Println(returnMathOperations()(num1, num2, add))
	fmt.Println(mathOperations(num1, num2, add))
	fmt.Println(mathOperations(num1, num2, subtract))
	fmt.Println(mathOperations(num1, num2, multiply))
	fmt.Println(mathOperations(num1, num2, divide))

}

func returnMathOperations() func(f1, f2 float32, f func(num1, num2 float32) float64) float64 {
	return func(f1, f2 float32, f func(num1, num2 float32) float64) float64 {
		return f(f1, f2)
	}
}

func mathOperations(f1, f2 float32, f func(num1, num2 float32) float64) float64 {
	return f(f1, f2)
}

func add(num1, num2 float32) float64 {
	return float64(num1 + num2)
}

func subtract(num1, num2 float32) float64 {
	return float64(num1 - num2)
}

func multiply(num1, num2 float32) float64 {
	return float64(num1 * num2)
}

func divide(num1, num2 float32) float64 {
	return float64(num1 / num2)
}
