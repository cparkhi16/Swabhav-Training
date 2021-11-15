package main

import "fmt"

func main() {
	//when cap changes a new address is assigned to slice
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	slice[3] = 10
	aar := []int{9, 0}
	aar = append(aar, 57, 78)
	fmt.Println(aar)
	fmt.Println("Slice before func call", slice)
	fmt.Println("Adress of slice[0] in main ", &slice[0])
	manipulateSlice(slice)
	fmt.Println("Slice after func call", slice)
	newSlice := make([]int, 5, 10)
	newSlice[0] = 700
	newSlice[1] = 800
	fmt.Println("New slice before append", newSlice)
	fmt.Println("Adress of newSlice[0] before append", &newSlice[0])
	fmt.Printf("Address using printf %p", &newSlice[0])
	newSlice = append(newSlice, 1000)
	fmt.Println("\nNew slice after append", newSlice)
	fmt.Println("Adress of newSlice[0] after append", &newSlice[0])
	fmt.Printf("Address using printf %p after appending ------\n", &newSlice[0])
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

	var abc []int
	abc = append(abc, 1)
	fmt.Println(abc)
}
func manipulateSlice(s []int) {
	fmt.Println("Adress of slice[0] in func call ", &s[0])
	s[0] = 100
	s[1] = 200
	s[3] = 300
	s = append(s, 7000) //MAIN DIFFERENCE TO NOTICE APPEND CHANGES NOT REFLECTS IN MAIN SINCE ADDRESS OF SLICE IS CHANGING
	fmt.Println("Address after append in func call", &s[0])
	fmt.Println("Slice in func call", s)
	fmt.Println()
	newSlice := []int{1, 2, 3, 4, 5}
	reSlice := newSlice[1:4]
	fmt.Println("Re slicing a slice ", reSlice, len(reSlice), cap(reSlice))
	nSlice := reSlice[2:]
	fmt.Println(" a reSlice ", reSlice, len(reSlice), cap(reSlice))
	fmt.Println("nSlice ", nSlice, len(nSlice), cap(nSlice))
	fmt.Println()
	a := []int{1, 2, 3, 4, 5, 6}
	for _, val := range a {
		fmt.Println("Range loop for slice", val)
	}
	b := a[1:3:5]
	c := a[1:3]
	fmt.Println("b ", b, len(b), cap(b))
	fmt.Println("c ", c, len(c), cap(c))
}
