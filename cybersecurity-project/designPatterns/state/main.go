package main

import (
	"fmt"
	"vendingApp/vendingMachine"
)

func main() {
	cokeMachine := vendingMachine.New(0, 13)

	err := cokeMachine.AddItem(100)
	if err != nil {
		fmt.Println(err)
	}
	err = cokeMachine.RequestItem()
	if err != nil {
		fmt.Println(err)
	}
	err = cokeMachine.InsertMoney(1)
	if err != nil {
		fmt.Println(err)
	}
	err = cokeMachine.DispenseItem()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("itemCount in cokeMachine-", cokeMachine.GetItemCount())

}
