package state

import "fmt"

type noItemState struct {
	vendingMachine *vendingMachine
}

func (n *noItemState) AddItem(count uint) error {
	if count != 0 {
		fmt.Printf("Added %v items to vending machine \n", count)
		n.vendingMachine.incrementItemCount(count)
		n.vendingMachine.setState(n.vendingMachine.hasItem)
		return nil
	} else {
		return fmt.Errorf("zero items should not be added ")
	}
}
func (n *noItemState) RequestItem() error {
	return fmt.Errorf("sorry no item in vending machine")
}
func (n *noItemState) InsertMoney(money uint) error {
	return fmt.Errorf("sorry no item in vending machine")
}
func (n *noItemState) DispenseItem() error {
	return fmt.Errorf("sorry no item in vending machine")
}
