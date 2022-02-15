package main

import (
	"machine/gate"
)

func main() {
	stationGate := gate.NewGate(20)
	stationGate.Pay(30)
	stationGate.Enter()
	stationGate.Pay(10)
	stationGate.Enter()
	stationGate.Enter()
	stationGate.PayOk()
	stationGate.PayFailed()
}
