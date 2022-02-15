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

func main() {
	anni := rubberDuck{
		name: "anni",
	}
	pan := mountainDuck{
		name: "pan",
	}
	anni.quack() //we don't know the implementation of quack and we are hiding unnecessary details such as duck name
	pan.quack()

}
