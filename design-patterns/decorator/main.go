package main

// Decorator pattern is used for wrapping up the object based on its modification we need to update its state
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
