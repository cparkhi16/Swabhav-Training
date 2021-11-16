package circle

import "math"

type Circle struct {
	radius float64
}

func New(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
