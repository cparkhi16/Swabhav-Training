package main

import "fmt"

func main() {
	slice_1 := []string{"Chinmay", "J", "Parkhi"}
	slice_2 := make([]string, 3)
	newSlice := copy(slice_2, slice_1)
	fmt.Println("Copying using copy func ", newSlice)
	fmt.Println("After copying slice_1", slice_1)
	fmt.Println("slice_2 ", slice_2)
	slice_2[0] = "GOLANG"
	fmt.Println("After changing slice_2", slice_2)
	fmt.Println("slice_1", slice_1)
}
