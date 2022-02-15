package main

import "fmt"

type square struct {
	side int
}

/*
VIOLATION
as per Single Responsibility principle, a function should only do one thing. area() method only should return
area and not display it. We can create a different method displayArea().
func (s square) area() {
	area := s.side * s.side
	fmt.Println("area-", area)
}*/

func (s square) area() int {
	return s.side * s.side
}

func (s square) displayArea() {
	fmt.Println("area-", s.side*s.side)
}

func main() {
	s := square{
		side: 3,
	}
	a := s.area()
	fmt.Println(a)
	s.displayArea()
}
