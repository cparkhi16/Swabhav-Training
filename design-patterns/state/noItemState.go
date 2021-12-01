package main

import "fmt"

type noItemState struct {
	vendingMachine *vendingMachine
}

func (n *noItemState) addItem(count uint) error {
	fmt.Printf("Added %v items to vending machine \n", count)
	n.vendingMachine.incrementItemCount(count)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}
func (n *noItemState) requestItem() error {
	return fmt.Errorf("sorry no item in vending machine")
}
func (n *noItemState) insertMoney(money uint) error {
	return fmt.Errorf("sorry no item in vending machine")
}
func (n *noItemState) dispenseItem() error {
	return fmt.Errorf("sorry no item in vending machine")
}
