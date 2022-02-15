package square

type square struct {
	side float64
}

func New(side float64) *square {
	return &square{
		side: side,
	}
}

func (s *square) GetSide() float64 {
	return s.side
}

func (s *square) SetSide(newside float64) {
	s.side = newside
}

func (s *square) Area() float64 {
	return s.side * s.side
}

func (s *square) Perimeter() float64 {
	return 4 * s.side
}
