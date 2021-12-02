package main

// State pattern is used when we want to change the state of objects and based on that change the functionality is to be provided to the user
import (
	s "design/state"
	"fmt"
	"log"
)

func main() {
	vendingMachine := s.NewVendingMachine(1, 100)
	fmt.Println("Number of items in vending machine --", vendingMachine.GetItemCount())
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(105)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Number of items in vending machine after dispensing one item--", vendingMachine.GetItemCount())
	err = vendingMachine.AddItem(0)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Number of items in vending machine after adding --", vendingMachine.GetItemCount())
	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(100)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Number of items in vending machine after dispensing one item--", vendingMachine.GetItemCount())
}
