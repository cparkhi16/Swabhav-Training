package main

import "fmt"

func main() {
	var i interface{} = "hello"
	checkType(i)
	i = 3
	checkType(i)

}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: int, Value:", i.(int))
	case string:
		fmt.Println("Type: string, Value: ", i.(string))
	case float64:
		fmt.Println("Type: float64, Value: ", i.(float64))
	default:
		fmt.Println("Type not found")
	}
}
