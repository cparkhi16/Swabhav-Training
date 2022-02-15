package main

import "fmt"

func main() {
	//Reslicing
	mySlice := make([]int, 5, 10)
	mySlice[0] = 10
	mySlice[1] = 20
	mySlice[2] = 30
	fmt.Println("mySlice-", mySlice)
	reslice := mySlice[:2]
	reslice[0] = 334
	fmt.Println("reslice-", reslice)
	fmt.Println("mySlice-", mySlice)

	newslice := mySlice
	newslice[0] = 223
	fmt.Println("newslice-", newslice)
	fmt.Println("myslice-", mySlice)

	copyslice := make([]int, 5, 10)
	num := copy(copyslice, mySlice)
	copyslice[0] = 989
	fmt.Println("number of elemets copied-", num)
	fmt.Println("copyslice-", copyslice)
	fmt.Println("mySlice-", mySlice)

	//maps
	myMap := make(map[string]int)
	myMap["mike"] = 23
	myMap["john"] = 45
	myMap["taylor"] = 89
	fmt.Println("map-", myMap)
	changeMap(myMap)
	fmt.Println("after changeMap fn call myMap-", myMap)

	//delete in map
	delete(myMap, "mike")
	delete(myMap, "wyn")
	fmt.Println("after deleting from map", myMap)

	map1 := make(map[string]int)
	var map2 map[string]int
	fmt.Println("map1", map1)
	fmt.Println("map2", map2)

	//in variaidic function pass slice
	test("hello", 20)
	test("hi", 30, 40)
	test("hii", mySlice...) //test("hi",10,20,30)

}

func changeMap(mymap map[string]int) {
	mymap["taylor"] = 13
}

func test(a string, b ...int) {
	fmt.Println("a", a)
	fmt.Println("b", b)
}
