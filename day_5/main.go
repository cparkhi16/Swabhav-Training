package main

import (
	"day_5/shape"
	"fmt"
)

func main() {
	circle := shape.Circle{}
	circle.SetCircelRadius(3.5)
	aCircle := circle.Area()
	fmt.Println("Area of circle ", aCircle)
	pCircle := circle.Perimeter()
	fmt.Println("Perimeter of circle ", pCircle)
	fmt.Println()
	sq := shape.Square{}
	sq.SetSquareDimensions(3)
	aSquare := sq.Area()
	fmt.Println("Area of square ", aSquare)
	pSquare := sq.Perimeter()
	fmt.Println("Perimeter of square ", pSquare)
	fmt.Println()
	re := shape.Rectangle{}
	re.SetRectangleDimensions(3, 4)
	aRect := re.Area()
	fmt.Println("Area of rectangle ", aRect)
	pRect := re.Perimeter()
	fmt.Println("Perimeter of rectangle ", pRect)
	fmt.Println()
	t := shape.Triangle{}
	t.SetTriangleDimensions(3, 4, 1, 2)
	aTri := t.Area()
	fmt.Println("Area of triangle ", aTri)
	pTri := t.Perimeter()
	fmt.Println("Perimeter of triangle ", pTri)
}
