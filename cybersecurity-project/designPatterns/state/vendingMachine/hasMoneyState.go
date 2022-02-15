package vendingMachine

import "fmt"

type HasMoneyState struct {
	vending *VendingMachine
}

func NewHasMoneyState(vending *VendingMachine) HasMoneyState {
	return HasMoneyState{
		vending: vending,
	}
}

func (h HasMoneyState) AddItem(count uint8) error {
	return fmt.Errorf("HasMoneyState- items are already added")
}

func (h HasMoneyState) RequestItem() error {
	return fmt.Errorf("HasMoneyState- items are already requested")
}

func (h HasMoneyState) InsertMoney(amount uint16) error {
	return fmt.Errorf("HasMoneyState- money is already inserted")
}

func (h HasMoneyState) DispenseItem() error {
	fmt.Println("HasMoneyState- dispencing item")
	if h.vending.GetItemCount() == 0 {
		h.vending.SetState(h.vending.noItemState)
		return fmt.Errorf("No items are available for dispensing")
	} else {
		h.vending.SetItemCount(h.vending.GetItemCount() - 1)
		h.vending.SetState(h.vending.hasItemState)
		fmt.Println("Item dispense done")
		return nil
	}
}
