package triangle

import "math"

type triangle struct {
	side1 float64
	side2 float64
	side3 float64
}

func New(side1 float64, side2 float64, side3 float64) *triangle {
	return &triangle{
		side1: side1,
		side2: side2,
		side3: side3,
	}
}

func (t *triangle) GetSides() (float64, float64, float64) {
	return t.side1, t.side2, t.side3
}

func (t *triangle) SetSides(side1 float64, side2 float64, side3 float64) {
	t.side1 = side1
	t.side2 = side2
	t.side3 = side3
}

func (t *triangle) Area() float64 {
	s := (t.side1 + t.side2 + t.side3) / 2
	return math.Sqrt(s * (s - t.side1) * (s - t.side2) * (s - t.side3))
}

func (t *triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.side3
}
