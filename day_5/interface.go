package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Square struct {
	side uint
}

func (s *Square) SetSquareDimensions(side uint) {
	s.side = side
}

type Circle struct {
	radius float64
}

func (c *Circle) SetCircelRadius(radius float64) {
	c.radius = radius
}

type Rectangle struct {
	length  uint
	breadth uint
}

func (r *Rectangle) SetRectangleDimensions(l, b uint) {
	r.breadth = b
	r.length = l
}
func (s Square) area() float64 {
	return float64(s.side * s.side)
}

func (s Square) perimeter() float64 {
	return float64(4 * s.side)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r Rectangle) area() float64 {
	return float64(r.length * r.breadth)
}
func (r Rectangle) perimeter() float64 {
	return float64(2 * (r.length + r.breadth))
}
func measure(s Shape) {
	fmt.Println("Type of shape :", reflect.TypeOf(s))
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	s := Square{}
	s.SetSquareDimensions(3)
	c := Circle{}
	c.SetCircelRadius(2)
	r := Rectangle{}
	r.SetRectangleDimensions(5, 6)

	measure(s)
	measure(c)
	measure(r)
}
