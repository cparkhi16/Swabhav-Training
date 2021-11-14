package main

import "fmt"

func main() {
	var mapExample map[int]int // Zero value of map -> Nil
	fmt.Printf("Address of map without make %p \n", &mapExample)
	map_2 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	map_2[96] = "Elephant"
	fmt.Println("Map-2: ", map_2)
	fmt.Println("Trying to print a non-existent key's value", map_2[95])
	var myNewMap = make(map[int]string)
	myNewMap[1] = "Chinmay"
	myNewMap[2] = "J"
	myNewMap[3] = "Parkhi"
	fmt.Println(myNewMap)
	fmt.Println("Trying to print a non-existent key's value with make", myNewMap[4])
	var NewMap = make(map[int]string)
	fmt.Printf("Address of map using make %p \n", &NewMap)
	if NewMap == nil {
		fmt.Println("nil NewMap")
	} else {
		fmt.Println("Not a nil NewMap")
	}

	value, ok := myNewMap[4]
	if ok == true {
		fmt.Println("Value for myNewMap[4]", value)
	} else {
		fmt.Println("Value not found")
	}

	//mapExample[0] = 1 //Error in assigning a nil map
	//fmt.Println(mapExample)
	/*mapExample={
		0:1,
		1:2
	}
	fmt.Println(mapExample)*/
	fmt.Println("myNewMap before copying it to copyMap", myNewMap)
	copyMap := myNewMap
	copyMap[0] = "Abc"
	copyMap[2] = "Pqr"
	fmt.Println("myNewMap after modification in copyMap ", myNewMap)
	fmt.Println("copy map using range for loop")
	for key, val := range copyMap {
		fmt.Println(key, val)
	}

}
