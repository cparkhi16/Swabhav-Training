package main

import "fmt"

type Dog struct {
	Name  string
	Breed string
	Age   int
}

func main() {
	nums := make([]int, 0, 4)
	nums = append(nums, 1, 2)
	fmt.Println("nums-", nums)
	fmt.Printf("address of nums-%p\n", &nums[0])

	a := append(nums, 4) //1235
	fmt.Printf("address of a-%p\n", &a[0])
	fmt.Println("nums-", nums)
	fmt.Printf("address of nums-%p\n", &nums[0])
	fmt.Println("a is :", a)

	b := append(nums, 5) //1235
	fmt.Printf("address of b-%p\n", &b[0])
	fmt.Println("nums-", nums)
	fmt.Printf("address of nums-%p\n", &nums[0])

	fmt.Println("b is :", b)
	fmt.Println("a is :", a)

	mySlice := make([]int, 5, 5)
	myMap := make(map[string]int)
	mySlice[0] = 10
	mySlice[1] = 11
	mySlice[2] = 13
	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	myMap["john"] = 23
	myMap["mayer"] = 34
	myMap["linken"] = 76
	for i, v := range myMap {
		fmt.Println(i, v)
	}

	var dogHouse []Dog
	dogHouse = append(dogHouse, Dog{
		Name:  "moti",
		Breed: "persian",
		Age:   73,
	})
	dogHouse = append(dogHouse, Dog{
		Name:  "liu",
		Breed: "mexican",
		Age:   32,
	})
	takeStruct(dogHouse)
}

func takeStruct(dogHouse []Dog) {
	for i, v := range dogHouse {
		fmt.Println(i, v)
	}
	fmt.Println("dogHouse-", dogHouse)
}
