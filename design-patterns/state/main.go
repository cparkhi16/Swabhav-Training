package main

import (
	"fmt"
	"log"
)

func main() {
	vendingMachine := newVendingMachine(1, 100)
	fmt.Println("Number of items in vending machine --", vendingMachine.getItemCount())
	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.insertMoney(105)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Number of items in vending machine after dispensing one item--", vendingMachine.getItemCount())
	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Number of items in vending machine after adding --", vendingMachine.getItemCount())
	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.insertMoney(100)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Number of items in vending machine after dispensing one item--", vendingMachine.getItemCount())
}
