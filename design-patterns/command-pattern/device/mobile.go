package device

import "fmt"

type Mobile struct {
	isRunning bool
}

func NewMobile() *Mobile {
	return &Mobile{isRunning: false}
}
func (t *Mobile) On() {
	if !t.isRunning {
		fmt.Println("Mobile is ON")
		t.isRunning = true
	} else {
		fmt.Println("Mobile is already ON")
	}
}
func (t *Mobile) Off() {
	if t.isRunning {
		fmt.Println("Mobile is OFF ")
		t.isRunning = false
	} else {
		fmt.Println("Mobile is already OFF")
	}
}
