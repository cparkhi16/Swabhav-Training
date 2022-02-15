package store

import (
	"fmt"
	"shoppingApp/item"
)

type Store struct {
	name     string
	itemList []*item.Item
}

func New(name string) *Store {
	return &Store{
		name: name,
	}
}

func (s *Store) AddItemToStore(item *item.Item) {
	s.itemList = append(s.itemList, item)
}

func (s *Store) GetItemFromItemId(id int) *item.Item {
	for _, v := range s.itemList {
		if id == v.GetId() {
			return v
		}
	}
	return nil
}

func (s *Store) DisplayItemList() {
	for i, v := range s.itemList {
		fmt.Println(i, " ", "itemId- ", v.GetId(), " itemCategory- ", v.GetCategory(), " itemPrice-", v.GetPrice(), " itemquantity-", v.GetQuantity())
	}
}
