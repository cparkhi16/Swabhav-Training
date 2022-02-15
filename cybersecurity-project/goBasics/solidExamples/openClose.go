package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}
type square struct {
	side float32
}

func (s square) area() float32 {
	return s.side
}

type rectangle struct {
	length float32
	width  float32
}

func (r rectangle) area() float32 {
	return r.length * r.width
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

/*VIOLATION
As per Open-close principle an entity should be open for extension and closed for modification. With this method
if we add one more struct circle then it needed slight modification in switch case hence this type of design
is not acceptable so we created new interface shape and implemented area() method on all the structs.
func areaSum(shapes ...interface{}) float32 {
	var sum float32 = 0.0
	for _, shape := range shapes {
		switch shape.(type) {
		case square:
			sum = sum + shape.(square).side*shape.(square).side
		case rectangle:
			sum = sum + shape.(rectangle).length*shape.(rectangle).width
		}

	}
	return sum
}*/

func areaSum(shapes ...shape) float32 {
	var sum float32 = 0.0
	for _, shape := range shapes {
		sum = sum + shape.area()
	}
	return sum
}

func main() {
	s := square{side: 3}
	r := rectangle{
		length: 2,
		width:  2,
	}
	c := circle{radius: 4}
	fmt.Println(areaSum(s, r, c))

}
