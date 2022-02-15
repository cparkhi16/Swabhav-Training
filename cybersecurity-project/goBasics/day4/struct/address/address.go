package address

type Address struct {
	roomNo int
	city   string
	state  string
}

func New(roomNo int, city string, state string) *Address {
	return &Address{
		roomNo: roomNo,
		city:   city,
		state:  state,
	}
}

func (a *Address) GetRoomNo() int {
	return a.roomNo
}

func (a *Address) SetRoomNo(newRoomNo int) {
	a.roomNo = newRoomNo
}

func (a *Address) GetCity() string {
	return a.city
}

func (a *Address) SetCity(newCity string) {
	a.city = newCity
}

func (a *Address) GetState() string {
	return a.state
}

func (a *Address) SetState(newState string) {
	a.state = newState
}
