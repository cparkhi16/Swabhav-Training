package state

import (
	s "design/stateinterface"
	"fmt"
)

type vendingMachine struct {
	hasItem       s.State
	itemRequested s.State
	hasMoney      s.State
	noItem        s.State
	currentState  s.State
	itemCount     uint
	itemPrice     uint
}

func NewVendingMachine(itemCount, itemPrice uint) *vendingMachine {
	vm := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: vm,
	}
	noItemState := &noItemState{
		vendingMachine: vm,
	}
	hasItemState := &hasItemState{
		vendingMachine: vm,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: vm,
	}

	vm.setState(hasItemState)
	vm.hasItem = hasItemState
	vm.itemRequested = itemRequestedState
	vm.hasMoney = hasMoneyState
	vm.noItem = noItemState
	return vm
}

func (vm *vendingMachine) incrementItemCount(count uint) {
	fmt.Printf("Adding %d items\n", count)
	vm.itemCount = vm.itemCount + count
}
func (vm *vendingMachine) RequestItem() error {
	return vm.currentState.RequestItem()
}

func (vm *vendingMachine) AddItem(count uint) error {
	return vm.currentState.AddItem(count)
}

func (vm *vendingMachine) InsertMoney(money uint) error {
	return vm.currentState.InsertMoney(money)
}

func (vm *vendingMachine) DispenseItem() error {
	return vm.currentState.DispenseItem()
}

func (vm *vendingMachine) setState(s s.State) {
	vm.currentState = s
}
func (vm *vendingMachine) GetItemCount() uint {
	return vm.itemCount
}
