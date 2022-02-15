package gate

import "fmt"

type ProcessingState struct {
	gate *Gate
}

func NewProcessingState(gate *Gate) ProcessingState {
	return ProcessingState{
		gate: gate,
	}
}

func (g ProcessingState) Enter() {
	fmt.Println("ProcessingState-wait processing")
}

func (g ProcessingState) Pay(amount uint8) {
	fmt.Println("ProcessingState-paid already")
}

func (g ProcessingState) PayOk() {
	fmt.Println("ProcessingState-payok opening the gate")
	g.gate.ChangeState(g.gate.openState)
}

func (g ProcessingState) PayFailed() {
	fmt.Println("ProcessingState-payFailed closing the gate")
	g.gate.ChangeState(g.gate.closedState)
}
