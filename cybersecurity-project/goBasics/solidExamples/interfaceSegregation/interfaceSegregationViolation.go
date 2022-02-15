package main

import "fmt"

//VIOLATION
/*Here square struct implements shape interface and in method getVolume it returns zero. square is a 2D shape
so it does not have volume so it it should not be forced to implement shape interface. That's why we segregate the interfaces
for square and cube, as twoDimensional and threedimensonal.
*/
type shape interface {
	getArea() float32
	getVolume() float32
}

type square struct {
	side float32
}

func (s square) getArea() float32 {
	return s.side * s.side
}

func (s square) getVolume() float32 {
	return 0.0
}

type cube struct {
	side float32
}

func (c cube) getArea() float32 {
	return 6 * c.side * c.side
}

func (c cube) getVolume() float32 {
	return c.side * c.side * c.side
}

func areaSum(shapes ...shape) float32 {
	var sum float32 = 0.0
	for _, v := range shapes {
		sum = sum + v.getArea()
	}
	return sum
}

func areaVolumeSum(shapes ...shape) float32 {
	var sum float32 = 0.0
	for _, v := range shapes {
		sum = sum + v.getArea() + v.getVolume()
	}
	return sum
}

func main() {
	s := square{side: 3}
	c := cube{side: 3}
	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(s, c))
}
