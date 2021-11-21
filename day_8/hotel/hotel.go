package hotel

import (
	"fmt"
	"sync"
)

//var wg sync.WaitGroup

type Order struct {
	orderID uint
	items   []Item
}
type Item struct {
	itemName  string
	itemID    uint
	itemPrice uint
}

func PlaceOrder(u *User, orderID, itemID, itemPrice uint, itemName string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Placing order for user", u.userName)
	i := Item{
		itemName:  itemName,
		itemID:    itemID,
		itemPrice: itemPrice,
	}
	order := Order{
		orderID: orderID,
		items:   []Item{i},
	}
	ch := make(chan string)
	deskNumber := make(chan int)
	go u.AddOrderForUser(order, ch)
	go u.AssignDeskForUser(deskNumber)
	ch <- "Order is added to your order list ..."
	deskNumber <- int(itemID) + 700

	fmt.Printf("Placed order for user %v  -============== \n", u.userName)
}

func (u *User) NewItemForUser(itemName string, itemID, itemPrice, orderID uint) {
	i := Item{
		itemName:  itemName,
		itemID:    itemID,
		itemPrice: itemPrice,
	}
	o := Order{
		orderID: orderID,
		items:   []Item{i},
	}
	ch := make(chan string)
	go u.AddOrderForUser(o, ch)
	ch <- "Order is added to your order list ..."
}
