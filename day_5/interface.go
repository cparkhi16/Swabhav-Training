package main

import (
	"day_5/shape/circle"
	"day_5/shape/square"
	i "day_5/shapeinterface"
)

func main() {

	s := square.New(3)
	c := circle.New(2)

	i.Measure(s)
	i.Measure(c)
}
