package gate

type GateState interface {
	Enter()
	Pay(amount uint8)
	PayOk()
	PayFailed()
}

type Gate struct {
	entryFee        uint8
	currentState    GateState
	openState       OpenState
	closedState     ClosedState
	processingState ProcessingState
}

func NewGate(entryFee uint8) *Gate {
	g := Gate{
		entryFee: entryFee,
	}
	g.openState = NewOpenState(&g)
	g.closedState = NewClosedState(&g)
	g.processingState = NewProcessingState(&g)
	g.ChangeState(g.closedState)
	return &g
}

func (g *Gate) GetEntryFee() uint8 {
	return g.entryFee
}

func (g *Gate) ChangeState(newState GateState) {
	g.currentState = newState
}

func (g *Gate) Enter() {
	g.currentState.Enter()
}

func (g *Gate) Pay(amount uint8) {
	g.ChangeState(g.processingState)
	if amount >= g.entryFee {
		g.currentState.PayOk()
	} else {
		g.currentState.PayFailed()
	}
}

func (g *Gate) PayOk() {
	g.currentState.PayOk()
}

func (g *Gate) PayFailed() {
	g.currentState.PayFailed()
}
