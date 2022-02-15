//calc package
package calc
func Add(first float64, second float64)float64{
	return first+second
}
func Subtract(first float64,second float64)float64{
	return first-second
}

//main package
package main
import (
	"calc"
	"fmt"
)
func main(){
	fmt.Println(calc.Add(1,2))
	fmt.Println(calc.Subtract(7,3))
}