package pizza

type cheeseTopping struct {
	pizza Pizza
	price uint32
}

func NewCheeseTopping(pizza Pizza, price uint32) *cheeseTopping {
	return &cheeseTopping{
		pizza: pizza,
		price: price,
	}
}

func (c cheeseTopping) GetPrice() uint32 {
	return c.pizza.GetPrice() + c.price
}
