package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i []interface{}
	i = append(i, "Hi")
	i = append(i, "Good Night")
	i = append(i, true)
	i = append(i, false)
	i = append(i, 89)
	i = append(i, 10.78)
	i = append(i, 1)
	segregateData(i)
}

func segregateData(si []interface{}) {
	var integer []int
	var boolean []bool
	var float []float64
	var s []string
	for i := 0; i < len(si); i++ {
		switch si[i].(type) {
		case int:
			fmt.Println(i, " Type:", reflect.TypeOf(si[i]))
			integer = append(integer, si[i].(int))
		case string:
			fmt.Println(i, " Type:", reflect.TypeOf(si[i]))
			s = append(s, si[i].(string))
		case float64:
			fmt.Println(i, " Type:", reflect.TypeOf(si[i]))
			float = append(float, si[i].(float64))
		case bool:
			fmt.Println(i, " Type:", reflect.TypeOf(si[i]))
			boolean = append(boolean, si[i].(bool))
		default:
			fmt.Println("None of the above")
		}
	}
	fmt.Println("Bool slice ", boolean)
	fmt.Println("String slice", s)
	fmt.Println("Float slice", float)
	fmt.Println("Integer slice", integer)
}
