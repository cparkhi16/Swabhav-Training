package gate

import "fmt"

type ClosedState struct {
	gate *Gate
}

func NewClosedState(gate *Gate) ClosedState {
	return ClosedState{
		gate: gate,
	}
}

func (g ClosedState) Enter() {
	fmt.Println("ClosedState-Gate is closed pay first")
	g.gate.ChangeState(g.gate.closedState)
}

func (g ClosedState) Pay(amount uint8) {
	fmt.Println("ClosedState-changing state to processing")
	g.gate.ChangeState(g.gate.processingState)
}

func (g ClosedState) PayOk() {
	fmt.Println("ClosedState-Gate is closed pay first")
}

func (g ClosedState) PayFailed() {
	fmt.Println("ClosedState-Gate is closed pay first")
}
