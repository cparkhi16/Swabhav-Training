package main

import "fmt"

func main() {
	myArray := [4]int{1, 2, 3, 4}
	var copyOfMyarray [4]int = myArray
	copyOfMyarray[3] = 0
	fmt.Println("myArray", myArray)
	fmt.Println("copyOfMyarray", copyOfMyarray)
	takeArrayAndChange(myArray)
	fmt.Println("myArray", myArray)
	takArrayAsPointer(&myArray)
	fmt.Println("myArray", myArray)
}

func takeArrayAndChange(array [4]int) {
	array[0] = 34
	fmt.Println("array", array)
}

func takArrayAsPointer(array *[4]int) {
	array[0] = 35
	fmt.Println(array[0])
}
