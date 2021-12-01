package stateinterface

type State interface {
	AddItem(uint) error
	RequestItem() error
	InsertMoney(money uint) error
	DispenseItem() error
}
