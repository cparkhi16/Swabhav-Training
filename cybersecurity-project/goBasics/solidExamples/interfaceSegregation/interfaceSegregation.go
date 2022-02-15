package main

import "fmt"

type twoDimensionalShape interface {
	getArea() float32
}

type threeDimentionalShape interface {
	twoDimensionalShape
	getVolume() float32
}

type square struct {
	side float32
}

func (s square) getArea() float32 {
	return s.side * s.side
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

func areaSum(shapes ...twoDimensionalShape) float32 {
	var sum float32 = 0.0
	for _, v := range shapes {
		sum = sum + v.getArea()
	}
	return sum
}

func areaVolumeSum(shapes ...threeDimentionalShape) float32 {
	var sum float32 = 0.0
	for _, v := range shapes {
		sum = sum + v.getArea() + v.getVolume()
	}
	return sum
}

func main() {
	s := square{side: 3}
	c := cube{side: 3}
	c2 := cube{side: 4}
	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(c, c2))
}
