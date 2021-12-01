package state

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (h *hasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	h.vendingMachine.itemCount = h.vendingMachine.itemCount - 1
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.setState(h.vendingMachine.hasItem)
	}
	return nil
}

func (h *hasMoneyState) RequestItem() error {
	return fmt.Errorf("a item is being processed")
}

func (h *hasMoneyState) AddItem(count uint) error {
	return fmt.Errorf("a item is being processed")
}

func (h *hasMoneyState) InsertMoney(money uint) error {
	return fmt.Errorf("a item is being processed")
}
