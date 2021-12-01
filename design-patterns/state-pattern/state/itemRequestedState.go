package state

import "fmt"

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) InsertMoney(money uint) error {
	fmt.Println("Money ", money)
	fmt.Println("item price ", i.vendingMachine.itemPrice)
	if money < i.vendingMachine.itemPrice {
		e := fmt.Errorf("you have inserted less money, please insert %d", i.vendingMachine.itemPrice)
		return e
	}
	fmt.Println("Amount is accepted .. Please wait for further processing...")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) RequestItem() error {
	return fmt.Errorf("a item is being processed")
}

func (i *itemRequestedState) AddItem(count uint) error {
	return fmt.Errorf("a item is being processed")
}

func (i *itemRequestedState) DispenseItem() error {
	return fmt.Errorf("enter money and then dispense item")
}
