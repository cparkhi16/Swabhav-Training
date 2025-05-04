package math

import "fmt"

func ArrayTest() {
	arr := [3]int{0, 1, 2}
	// arr := [...]int{1, 2, 3} this is also a go array
	fmt.Println(" before mani array ", arr)
	maniArray(&arr)
	fmt.Println(" after mani array ", arr)
}

func maniArray(ar *[3]int) {
	ar[1] = 100
	for _, v := range ar {
		fmt.Println(v)
	}
}

func SliceTest() {
	// slice can be initialized using make([]int,3,5) / []int{1,2,3} / or using : op on arrays
	arr := []int{1, 2, 3} // this is a slice in GO
	fmt.Printf("type of arr in slice test %T\n", arr)
	fmt.Println(" before Len and cap of slice is ", len(arr), cap(arr))
	arr = append(arr, 1, 44, 55, 9)
	fmt.Println("new slice is ", arr)
	fmt.Println(" after Len and cap of slice is ", len(arr), cap(arr))
	maniSlice(arr)
	fmt.Println("after manipulation slice is ", arr)
}

func maniSlice(arr []int) {
	arr[3] = 500
	arr = append(arr, 20, 30, 4, 69, 80)
	arr[0] = -9
	fmt.Println(" arr after append ", arr, len(arr), cap(arr)) //when cap changes a new address is assigned to slice
	//so above append changes won't be reflected in SliceTest()

}

func ConvertArrToSlice() {
	arr := [...]int{1, 2, 3}
	slice := arr[:]
	slice = append(slice, 4, 5)
	fmt.Println(" converted arr to slice ", slice)
	fmt.Println(" if a slice used as queue ,front elem ", slice[0])
	slice = slice[1:] // pop from front
	fmt.Println(" after pop slice ", slice)
	sl := slice[1:3] // slice[low:high:capacity]
	fmt.Println(" og slice and new sl after : operation ", slice, sl)
	sl[0] = 100
	fmt.Println(" og slice and new sl after : operation and manipulation ", slice, sl)

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	// Append slice2 to slice1
	slice1 = append(slice1, slice2...)

	// Print the resulting slice
	fmt.Println("Combined Slice:", slice1)

	new_slice := make([]int, 4, 8)
	// Initially, the slice will have 4 elements, all set to 0 (the zero value for int).
	new_slice = append(new_slice, 33)
	fmt.Println(" size and len of new slice ", new_slice, cap(new_slice), len(new_slice))
	ns := new_slice[1:3:8] // max should not extend og slice cap , ns slice cap will be 8-1
	fmt.Println(" new slice's details ", ns, len(ns), cap(ns))
}

func DeepCopySlice() {
	a := make([]int, 3)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	b := make([]int, 3)
	//b = append(b, 50)
	copy(b, a)
	//b = a
	b[1] = 100
	fmt.Println(" b is ", b, " a is ", a)

	c := append([]int{}, a...)
	c[0] = 900
	fmt.Println(" c is ", c, "  a is ", a)
}
