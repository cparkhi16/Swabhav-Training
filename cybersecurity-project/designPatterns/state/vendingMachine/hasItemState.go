package vendingMachine

import (
	"fmt"
)

type HasItemState struct {
	vending *VendingMachine
}

func NewHasItemState(vending *VendingMachine) HasItemState {
	return HasItemState{
		vending: vending,
	}
}

func (h HasItemState) AddItem(count uint8) error {
	return fmt.Errorf("HasItemState- items are already added")
}

func (h HasItemState) RequestItem() error {
	fmt.Println("HasItemState- requesting item")
	h.vending.SetState(h.vending.itemRequestedState)
	return nil
}

func (h HasItemState) InsertMoney(amount uint16) error {
	return fmt.Errorf("HasItemState- first request the item")
}

func (h HasItemState) DispenseItem() error {
	return fmt.Errorf("HasItemState- first request the item")
}
