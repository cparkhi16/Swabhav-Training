package main

import "fmt"

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (h *hasItemState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("no item present in vending machine")
	}
	fmt.Println("Item requestd")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *hasItemState) addItem(count uint) error {
	fmt.Printf("%d items added to vending machine \n", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h *hasItemState) insertMoney(money uint) error {
	return fmt.Errorf("a item is being processed")
}
func (h *hasItemState) dispenseItem() error {
	return fmt.Errorf("a item is being processed")
}
