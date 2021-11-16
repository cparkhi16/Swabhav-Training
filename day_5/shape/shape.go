package shape

import (
	"math"
)

type Circle struct {
	radius float64
}

func (c *Circle) SetCircelRadius(radius float64) {
	c.radius = radius
}
func (r *Circle) Area() float64 {
	return math.Pi * r.radius * r.radius
}
func (r *Circle) Perimeter() float64 {
	return 2 * math.Pi * r.radius
}

type Square struct {
	side uint
}

func (s *Square) SetSquareDimensions(side uint) {
	s.side = side
}
func (s *Square) Area() uint {
	return s.side * s.side
}
func (s *Square) Perimeter() uint {
	return 4 * s.side
}

type Rectangle struct {
	length  uint
	breadth uint
}

func (r *Rectangle) SetRectangleDimensions(l, b uint) {
	r.breadth = b
	r.length = l
}
func (r *Rectangle) Area() uint {
	return r.length * r.breadth
}
func (r *Rectangle) Perimeter() uint {
	return 2 * (r.length + r.breadth)
}

type Triangle struct {
	base    uint
	height  uint
	sideOne uint
	sideTwo uint
}

func (t *Triangle) SetTriangleDimensions(b, h, s1, s2 uint) {
	t.base = b
	t.height = h
	t.sideOne = s1
	t.sideTwo = s2
}
func (t *Triangle) Area() float64 {
	return 0.5 * float64(t.base) * float64(t.height)
}
func (t *Triangle) Perimeter() uint {
	return t.base + t.sideOne + t.sideTwo
}
