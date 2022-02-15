package customer

import (
	"fmt"
	"shoppingApp/item"
	"shoppingApp/store"
	"sync"
)

type customer struct {
	id       int
	name     string
	itemList []*item.Item
	balance  uint
}

func New(id int, name string, balance uint) *customer {
	return &customer{
		id:      id,
		name:    name,
		balance: balance,
	}
}

func (c *customer) AddToCart(item *item.Item) {
	c.itemList = append(c.itemList, item)
}

func (c *customer) GetId() int {
	return c.id
}

func (c *customer) GetName() string {
	return c.name
}

func (c *customer) GetBalance() uint {
	return c.balance
}

func (c *customer) SetBalance(balance uint) {
	c.balance = balance
}

func (c *customer) DisplayItemList() {
	for i, v := range c.itemList {
		fmt.Println(i, " ", "itemId- ", v.GetId(), " itemCategory- ", v.GetCategory(), " itemPrice-", v.GetPrice(), " itemquantity-", v.GetQuantity())
	}
}

func (c *customer) PlaceOrder(s *store.Store, m *sync.Mutex, wg *sync.WaitGroup) {
	var totalPrice uint = 0
	fmt.Println("----------Placing order for ", c.GetName(), "-----------")
	m.Lock()
	for _, v := range c.itemList {
		fmt.Println("For item ", v.GetCategory())
		if v.GetQuantity() == 0 {
			fmt.Println("itemId- ", v.GetId(), " not available")
			continue
		} else if v.GetQuantity() > s.GetItemFromItemId(v.GetId()).GetQuantity() {
			fmt.Println("itemId- ", v.GetId(), " not available")
			continue
		}
		item := s.GetItemFromItemId(v.GetId())
		item.SetQuantity(item.GetQuantity() - v.GetQuantity())
		totalPrice = totalPrice + v.GetPrice()*v.GetQuantity()
	}
	m.Unlock()
	c.balance = c.balance - totalPrice
	wg.Done()
}
