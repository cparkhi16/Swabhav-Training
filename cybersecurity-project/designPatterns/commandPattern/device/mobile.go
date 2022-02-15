package device

import "fmt"

type mobile struct {
	isRunning bool
}

func NewMobile() *mobile {
	return &mobile{
		isRunning: false,
	}
}

func (t *mobile) On() {
	if t.isRunning {
		fmt.Println("Mobile is already on")
	} else {
		fmt.Println("Mobile is turned on")
		t.isRunning = true
	}
}

func (t *mobile) Off() {
	if !t.isRunning {
		fmt.Println("Mobile is already off")
	} else {
		fmt.Println("Mobile is turned off")
		t.isRunning = false
	}
}
