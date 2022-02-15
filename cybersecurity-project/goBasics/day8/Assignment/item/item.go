package item

type Item struct {
	id       int
	category string
	price    uint
	quantity uint
}

func New(id int, category string, price uint, quantity uint) *Item {
	return &Item{
		id:       id,
		category: category,
		price:    price,
		quantity: quantity,
	}
}

func (i *Item) GetId() int {
	return i.id
}

func (i *Item) GetCategory() string {
	return i.category
}

func (i *Item) GetPrice() uint {
	return i.price
}

func (i *Item) GetQuantity() uint {
	return i.quantity
}

func (i *Item) SetQuantity(quantity uint) {
	i.quantity = quantity
}
