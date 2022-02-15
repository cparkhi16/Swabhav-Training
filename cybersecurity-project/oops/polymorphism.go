package main

import "fmt"

type Duck interface {
	quack()
}

type rubberDuck struct {
	name string
}

func (r rubberDuck) quack() {
	fmt.Println("rubber!")
}

type mountainDuck struct {
	name string
}

func (m mountainDuck) quack() {
	fmt.Println("mountain!")
}

func makeNoise(duck Duck) {
	duck.quack()
}

func main() {
	anni := rubberDuck{
		name: "anni",
	}
	pan := mountainDuck{
		name: "pan",
	}
	makeNoise(anni) //makeNoise works for both the type of ducks
	makeNoise(pan)

}
