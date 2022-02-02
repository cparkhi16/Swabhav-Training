package model

type Item struct {
	Name  string
	Price string
}

func NewItem(name, price string) *Item {
	return &Item{Name: name, Price: price}
}
