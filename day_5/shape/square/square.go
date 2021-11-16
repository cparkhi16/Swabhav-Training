package square

type Square struct {
	side uint
}

func New(s uint) *Square {
	return &Square{side: s}
}

/*func (s *Square) SetSquareDimensions(side uint) {
	s.side = side
}*/
func (s Square) Area() float64 {
	return float64(s.side * s.side)
}

func (s Square) Perimeter() float64 {
	return float64(4 * s.side)
}
