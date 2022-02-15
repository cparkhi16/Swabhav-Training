package main

import (
	"fmt"
	"reflect"
	Ishape "shape/Ishape"
	circle "shape/circle"
	rectangle "shape/rectangle"
	square "shape/square"
	triangle "shape/triangle"
)

func main() {
	c := circle.New(3.2)
	getDetails(c)

	r := rectangle.New(2, 3)
	getDetails(r)

	t := triangle.New(41, 28, 15)
	getDetails(t)

	s := square.New(3)
	getDetails(s)
}

func getDetails(shape Ishape.Ishape) {
	fmt.Println(reflect.TypeOf(shape).String(), "-", shape)
	if reflect.TypeOf(shape).String() == "*circle.Circle" {
		fmt.Println("area of a circle-", shape.Area())
		fmt.Println("circumference of a circle-", shape.Perimeter())
	} else {
		fmt.Println("area of a shape-", shape.Area())
		fmt.Println("perimeter of a shape-", shape.Perimeter())
	}
}
