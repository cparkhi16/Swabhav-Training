package main

import "fmt"

func main() {
	mySlice := make([]int, 8, 10)
	fmt.Printf("address of mySlice:%p\n", &mySlice)
	fmt.Println("address of first element of mySlice-", &mySlice[0])
	fmt.Println("lenght of mySlice", len(mySlice))
	fmt.Println("capacity of mySlice", cap(mySlice))
	mySlice[0] = 10
	mySlice[1] = 20
	mySlice[2] = 30
	var copyOfMySlice []int = mySlice
	copyOfMySlice[0] = 5
	fmt.Println("copyOfMySlice-", copyOfMySlice)
	fmt.Println("mySlice-", mySlice)

	takeSliceAndChange(mySlice)
	fmt.Println("mySlice after takeSliceAndChange fn call-", mySlice)

	mySlice = append(mySlice, 40)
	mySlice = append(mySlice, 40)
	mySlice = append(mySlice, 40)
	mySlice = append(mySlice, 40)
	mySlice = append(mySlice, 40)
	mySlice = append(mySlice, 40)
	fmt.Println("mySlice after exceeding capacity", mySlice)
	fmt.Printf("address of mySlice after exceeding capacity:%p\n", &mySlice)
	fmt.Println("address of first element of mySlice after exceeding capacity-", &mySlice[0])
	fmt.Println("lenght of mySlice after exceeding capacity", len(mySlice))
	fmt.Println("capacity of mySlice after exceeding capacity", cap(mySlice))

	//slice of pointers to string
	s := make([]string, 5, 5)
	s[0] = "abc"
	s[1] = "pqr"
	s[2] = "lmn"
	fmt.Println("slice of strings", s)
	p := make([]*string, 5, 5)
	p[0] = &s[0]
	p[1] = &s[1]
	fmt.Println("slice of pointers", p) //zero value of pointer is nil
	fmt.Println("deref 1st element of slice", *p[0])
	fmt.Println("deref 2nd element of slice", *p[1])
	fmt.Println("deref 3rd element of slice", *p[2]) //panic

	//reslicing
	fmt.Println("reslicing and showing first 3 elements-", mySlice[:4])

}

func takeSliceAndChange(slice []int) {
	//slice[0] = 45
	slice = append(slice, 45)
	fmt.Println("slice from takeSliceAndChange fn-", slice)
}
