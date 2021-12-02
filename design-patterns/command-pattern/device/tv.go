package device

import "fmt"

type TV struct {
	isRunning bool
}

func NewTV() *TV {
	return &TV{isRunning: false}
}
func (t *TV) On() {
	if !t.isRunning {
		fmt.Println("TV is ON")
		t.isRunning = true
	} else {
		fmt.Println("TV is already ON")
	}
}
func (t *TV) Off() {
	if t.isRunning {
		fmt.Println("TV is OFF ")
		t.isRunning = false
	} else {
		fmt.Println("TV is already OFF")
	}
}
