package gate

import "fmt"

type OpenState struct {
	gate *Gate
}

func NewOpenState(gate *Gate) OpenState {
	return OpenState{
		gate: gate,
	}
}

func (g OpenState) Enter() {
	fmt.Println("OpenState-please enter, after sometime gate will be closed")
	g.gate.ChangeState(g.gate.closedState)
}

func (g OpenState) Pay(amount uint8) {
	fmt.Println("OpenState-Gate is already open")
}

func (g OpenState) PayOk() {
	fmt.Println("OpenState-Gate is already open")
}

func (g OpenState) PayFailed() {
	fmt.Println("OpenState-Gate is already open")
}
