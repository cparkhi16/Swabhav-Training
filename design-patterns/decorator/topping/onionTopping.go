package topping

import p "decor/pizzaInterface"

type OnionTopping struct {
	Pizza p.Pizza
}

func (o *OnionTopping) GetPrice() uint32 {
	pizzaPrice := o.Pizza.GetPrice()
	return pizzaPrice + 20
}
