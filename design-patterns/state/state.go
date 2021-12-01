package main

type state interface {
	addItem(uint) error
	requestItem() error
	insertMoney(money uint) error
	dispenseItem() error
}
