package main

import (
	"fmt"
	"pizzaShop/pizza"
)

func main() {
	pizza1 := pizza.NewExtraVaganza(100)
	fmt.Println("price of pizza1-", pizza1.GetPrice())
	pizza1WithCheeseTopping := pizza.NewCheeseTopping(pizza1, 20)
	fmt.Println("price of pizza1WithCheeseTopping-", pizza1WithCheeseTopping.GetPrice())
	pizza1WithCheeseAndOnionTopping := pizza.NewOnionTopping(pizza1WithCheeseTopping, 20)
	fmt.Println("price of pizza1WithCheeseAndOnionTopping-", pizza1WithCheeseAndOnionTopping.GetPrice())
}
