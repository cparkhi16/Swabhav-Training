package main

import "fmt"

func main() {
	//when cap changes a new address is assigned to slice
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	slice[3] = 10
	fmt.Println("Slice before func call", slice)
	fmt.Println("Adress of slice[0] in main ", &slice[0])
	manipulateSlice(slice)
	fmt.Println("Slice after func call", slice)
	newSlice := make([]int, 5, 10)
	newSlice[0] = 700
	newSlice[1] = 800
	fmt.Println("New slice before append", newSlice)
	fmt.Println("Adress of newSlice[0] before append", &newSlice[0])
	newSlice = append(newSlice, 1000)
	fmt.Println("New slice after append", newSlice)
	fmt.Println("Adress of newSlice[0] after append", &newSlice[0])
	modifiedSlice := make([]int, 5, 5)
	modifiedSlice[1] = 8
	modifiedSlice[2] = 777
	fmt.Println("Before append address of modifiedSlice[0] when capacity is", cap(modifiedSlice), ": ", &modifiedSlice[0])
	manipulateSlice(modifiedSlice)
	//fmt.Println("After manipulating modified slice", modifiedSlice)
	modifiedSlice = append(modifiedSlice, 501, 122)
	fmt.Println("After append modifiedSlice", modifiedSlice)
	fmt.Println("After append address of modifiedSlice[0]", &modifiedSlice[0])
	fmt.Println("Modified slice cap", cap(modifiedSlice))
	fmt.Println()
	ptrOfSlices := make([]*string, 5, 5)
	var sOne string = "Chinmay"
	var sTwo string = "Parkhi"
	ptrOfSlices[0] = &sOne
	ptrOfSlices[1] = &sTwo
	fmt.Println("Pointers slice", ptrOfSlices)
	for i := 0; i < len(ptrOfSlices); i++ {
		if ptrOfSlices[i] != nil {
			fmt.Println("Element at", ptrOfSlices[i], "is ", *ptrOfSlices[i])
		} else {
			fmt.Println("Nil element found")
		}
	}
	fmt.Println("Slicing a pointers slice ", ptrOfSlices[0:2])
}
func manipulateSlice(s []int) {
	fmt.Println("Adress of slice[0] in func call ", &s[0])
	s[0] = 100
	s[1] = 200
	s[3] = 300
	fmt.Println("Slice in func call", s)
}
