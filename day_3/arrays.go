package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 10}
	manipulateArray(arr)
	fmt.Println("Main array is ", arr)

}
func manipulateArray(arr [3]int) {
	arr[2] = 1
	fmt.Println("Array in func call ", arr)
}
