package circle

import "math"

type Circle struct {
	radius float64
}

func New(radius float64) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(newRadius float64) {
	c.radius = newRadius
}

func (c *Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
