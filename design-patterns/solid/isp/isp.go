package main

import "fmt"

type Shape interface {
	//Volume() uint	 // Violates ISP as every shape will not have volume
	Area() uint
}

type ThreeDimensionalShape interface {
	Shape
	Volume() uint
}
type Square struct {
	side uint
}
type Cube struct {
	edge uint
}

/*func (s *Square) Volume() uint {
	return 0
}*/
func (c *Cube) Area() uint {
	return 6 * c.edge * c.edge
}
func (c *Cube) Volume() uint {
	return c.edge * c.edge * c.edge
}
func (s *Square) Area() uint {
	return s.side * s.side
}
func main() {
	s := &Square{side: 3}
	fmt.Println("Area of square :", s.Area())
	//fmt.Println(s.Volume())
	c := &Cube{edge: 2}
	fmt.Println("Area of cube :", c.Area())
	fmt.Println("Volume of cube :", c.Volume())
}
