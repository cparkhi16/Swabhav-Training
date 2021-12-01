package main

import (
	t "decor/topping"
	"fmt"
)

func main() {

	pizza := &extraVaganza{}
	fmt.Println("Price of extraVaganza pizza", pizza.GetPrice())
	pizzaWithCheese := &t.CheeseTopping{
		Pizza: pizza,
	}
	fmt.Println("Price of extraVaganza pizza with cheese topping", pizzaWithCheese.GetPrice())
	pizzaWithCheeseAndOnion := &t.OnionTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of pizza after adding cheese and onion Topping %d\n", pizzaWithCheeseAndOnion.GetPrice())
}
