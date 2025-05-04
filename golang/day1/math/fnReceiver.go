package main

import "fmt"

type Square struct {
	side float64
}

func (s *Square) SetSide(newSide float64) {
	s.side = newSide
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func main() {
	sq := Square{side: 5}
	sq.SetSide(10)         // Modifies the original struct
	fmt.Println(sq.Area()) // Output: 100
}
