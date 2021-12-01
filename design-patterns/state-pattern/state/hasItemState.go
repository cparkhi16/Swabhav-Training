package state

import "fmt"

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (h *hasItemState) RequestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("no item present in vending machine")
	}
	fmt.Println("Item requestd")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *hasItemState) AddItem(count uint) error {
	fmt.Printf("%d items added to vending machine \n", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h *hasItemState) InsertMoney(money uint) error {
	return fmt.Errorf("a item is being processed")
}
func (h *hasItemState) DispenseItem() error {
	return fmt.Errorf("a item is being processed")
}
