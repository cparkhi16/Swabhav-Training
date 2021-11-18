package main

import (
	"fmt"
)

type Population int
type Number int
type Liters float64
type Milliliters float64
type Gallons float64

func population() {
	var population Population
	population = Population(572)
	fmt.Println("Sleepy Creek Population:", population)
	fmt.Println("Congratulations.Kevin and Anna! It's a girl!")
	population += 1
	fmt.Println("Sleepy Creek Population:", population)
}
func (n Number) Add(otherNum int) {
	fmt.Println(n, "plus", otherNum, "is", int(n)+otherNum)
}
func (n Number) Sub(otherNum int) {
	fmt.Println(n, "minus", otherNum, "is", int(n)-otherNum)
}
func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(1 * 1000)
}
func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func main() {
	population()
	ten := Number(10)
	ten.Add(4)
	ten.Sub(5)
	four := Number(4)
	four.Add(3)
	four.Sub(2)
	l := Liters(3)
	fmt.Printf("%0.1f liters is %0.1f mililites\n", l, l.ToMilliliters())
	ml := Milliliters(500)
	fmt.Printf("%0.1f mililiters is %0.1f mililites\n", l, ml.ToLiters())

}
