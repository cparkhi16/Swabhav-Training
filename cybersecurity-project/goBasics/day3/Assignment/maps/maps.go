package main

import "fmt"

func main() {
	//to create a map
	myMap := make(map[string]int)
	myMap["mike"] = 23
	myMap["john"] = 45
	myMap["taylor"] = 89
	fmt.Println("map-", myMap)

	//to delete a element from map
	delete(myMap, "mike")
	fmt.Println("After deleting a element map-", myMap)

	value, ok := myMap["taylor"]
	fmt.Println("The value from myMap for key 'taylor':", value, "Present?", ok)

}
