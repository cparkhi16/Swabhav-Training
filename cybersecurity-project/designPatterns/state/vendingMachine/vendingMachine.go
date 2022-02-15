package vendingMachine

type State interface {
	AddItem(count uint8) error
	RequestItem() error
	InsertMoney(amount uint16) error
	DispenseItem() error
}

type VendingMachine struct {
	noItemState        State
	hasItemState       State
	itemRequestedState State
	hasMoneyState      State
	currentState       State
	itemCount          uint8
	itemPrice          uint16
}

func New(itemCount uint8, itemPrice uint16) *VendingMachine {
	v := VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	v.noItemState = NewNoItemState(&v)
	v.hasItemState = NewHasItemState(&v)
	v.itemRequestedState = NewItemRequestedState(&v)
	v.hasMoneyState = NewHasMoneyState(&v)
	v.SetState(v.noItemState)
	return &v
}

func (v *VendingMachine) AddItem(count uint8) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) InsertMoney(amount uint16) error {
	return v.currentState.InsertMoney(amount)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) SetState(newState State) {
	v.currentState = newState
}

func (v *VendingMachine) GetState() State {
	return v.currentState
}

func (v *VendingMachine) IncrementItemCount(count uint8) {
	v.itemCount = v.itemCount + count
}

func (v *VendingMachine) GetItemCount() uint8 {
	return v.itemCount
}

func (v *VendingMachine) SetItemCount(count uint8) {
	v.itemCount = count
}

func (v *VendingMachine) GetItemPrice() uint16 {
	return v.itemPrice
}
