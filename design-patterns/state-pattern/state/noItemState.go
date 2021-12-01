package state

import "fmt"

type noItemState struct {
	vendingMachine *vendingMachine
}

func (n *noItemState) AddItem(count uint) error {
	fmt.Printf("Added %v items to vending machine \n", count)
	n.vendingMachine.incrementItemCount(count)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
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
