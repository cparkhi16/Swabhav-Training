package main

import "fmt"

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state
	currentState  state
	itemCount     uint
	itemPrice     uint
}

func newVendingMachine(itemCount, itemPrice uint) *vendingMachine {
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
func (vm *vendingMachine) requestItem() error {
	return vm.currentState.requestItem()
}

func (vm *vendingMachine) addItem(count uint) error {
	return vm.currentState.addItem(count)
}

func (vm *vendingMachine) insertMoney(money uint) error {
	return vm.currentState.insertMoney(money)
}

func (vm *vendingMachine) dispenseItem() error {
	return vm.currentState.dispenseItem()
}

func (vm *vendingMachine) setState(s state) {
	vm.currentState = s
}
func (vm *vendingMachine) getItemCount() uint {
	return vm.itemCount
}
