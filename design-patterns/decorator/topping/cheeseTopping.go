package topping

import p "decor/pizzaInterface"

type CheeseTopping struct {
	Pizza p.Pizza
}

func (c *CheeseTopping) GetPrice() uint32 {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 20
}
