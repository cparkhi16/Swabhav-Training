package vendingMachine

import "fmt"

type ItemRequestedState struct {
	vending *VendingMachine
}

func NewItemRequestedState(vending *VendingMachine) ItemRequestedState {
	return ItemRequestedState{
		vending: vending,
	}
}

func (h ItemRequestedState) AddItem(count uint8) error {
	return fmt.Errorf("ItemRequestedState- items are already added")
}

func (h ItemRequestedState) RequestItem() error {
	return fmt.Errorf("ItemRequestedState- items are already requested")
}

func (h ItemRequestedState) InsertMoney(amount uint16) error {
	fmt.Println("ItemRequestedState- inserting money")
	if amount >= h.vending.GetItemPrice() {
		h.vending.SetState(h.vending.hasMoneyState)
		return nil
	} else {
		return fmt.Errorf("ItemRequestedState- inserted money is not sufficient to buy item")
	}

}

func (h ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("ItemRequestedState- first insert the money")
}
