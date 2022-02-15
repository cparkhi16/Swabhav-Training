package pizza

type onionTopping struct {
	pizza Pizza
	price uint32
}

func NewOnionTopping(pizza Pizza, price uint32) *onionTopping {
	return &onionTopping{
		pizza: pizza,
		price: price,
	}
}

func (o onionTopping) GetPrice() uint32 {
	return o.pizza.GetPrice() + o.price
}
