package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i []interface{}
	i = append(i, "first")
	i = append(i, "second")
	i = append(i, 23)
	i = append(i, false)
	i = append(i, 34.56)
	i = append(i, 14)
	fmt.Println("slice before conversion-", i)
	convertAllIntegersToStrings(i)
	fmt.Println("slice after conversion-", i)
}

func convertAllIntegersToStrings(slice []interface{}) {
	for i := 0; i < len(slice); i++ {
		switch slice[i].(type) {
		case int:
			fmt.Println(i, " Type: int, Value:", slice[i].(int))
			slice[i] = strconv.Itoa(slice[i].(int)) + "number"
		case string:
			fmt.Println(i, " Type: string, Value: ", slice[i].(string))
		case float64:
			fmt.Println(i, " Type: float64, Value: ", slice[i].(float64))
		case bool:
			fmt.Println(i, " Type: bool, Value: ", slice[i].(bool))
		default:
			fmt.Println("Type not found")
		}
	}
}
