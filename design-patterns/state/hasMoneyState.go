package main

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (h *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	h.vendingMachine.itemCount = h.vendingMachine.itemCount - 1
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.setState(h.vendingMachine.hasItem)
	}
	return nil
}

func (h *hasMoneyState) requestItem() error {
	return fmt.Errorf("a item is being processed")
}

func (h *hasMoneyState) addItem(count uint) error {
	return fmt.Errorf("a item is being processed")
}

func (h *hasMoneyState) insertMoney(money uint) error {
	return fmt.Errorf("a item is being processed")
}
