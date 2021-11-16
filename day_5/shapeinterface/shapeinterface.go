package shapeinterface

import (
	"fmt"
	"reflect"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func Measure(s Shape) {
	fmt.Println("Type of shape :", reflect.TypeOf(s))
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}
