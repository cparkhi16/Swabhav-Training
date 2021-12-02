package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	area() uint
}

type Square struct {
	side uint
}

type Rectangle struct {
	length  uint
	breadth uint
}

func (s *Square) area() uint {
	return s.side * s.side
}
func (r *Rectangle) area() uint {
	return r.length * r.breadth
}

/*func Measure(s interface{}) { // Violates OCP because if we want to add new Shape eg Triangle then we need to add it here again within switch case
	switch s.(type) {
	case Square:
		side := s.(Square).side
		fmt.Println("Area of square is ", side*side)
	case Rectangle:
		l := s.(Rectangle).length
		b := s.(Rectangle).breadth
		fmt.Println("Area of rectangle is", l*b)
	}
}*/
func Measure(s Shape) {
	fmt.Printf("Area of %v is %v \n", reflect.TypeOf(s), s.area())
}

func main() {
	s := &Square{side: 4}
	r := &Rectangle{length: 2, breadth: 3}
	Measure(s)
	Measure(r)
}
