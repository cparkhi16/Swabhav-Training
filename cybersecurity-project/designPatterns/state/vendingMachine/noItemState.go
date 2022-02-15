package vendingMachine

import (
	"fmt"
)

type NoItemState struct {
	vending *VendingMachine
}

func NewNoItemState(vending *VendingMachine) NoItemState {
	return NoItemState{
		vending: vending,
	}
}

func (n NoItemState) AddItem(count uint8) error {
	fmt.Println("NoItemState- adding item")
	if count == 0 {
		return fmt.Errorf("item count cannot be zero")
	}
	n.vending.IncrementItemCount(count)
	n.vending.SetState(n.vending.hasItemState)
	return nil
}

func (n NoItemState) RequestItem() error {
	return fmt.Errorf("NoItemState- items are out of stock")
}

func (n NoItemState) InsertMoney(amount uint16) error {
	return fmt.Errorf("NoItemState- items are out of stock")
}

func (n NoItemState) DispenseItem() error {
	return fmt.Errorf("NoItemState- items are out of stock")
}
