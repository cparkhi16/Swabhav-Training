package device

import "fmt"

type tv struct {
	isRunning bool
}

func NewTv() *tv {
	return &tv{
		isRunning: false,
	}
}

func (t *tv) On() {
	if t.isRunning {
		fmt.Println("TV is already on")
	} else {
		fmt.Println("TV is turned on")
		t.isRunning = true
	}
}

func (t *tv) Off() {
	if !t.isRunning {
		fmt.Println("TV is already off")
	} else {
		fmt.Println("TV is turned off")
		t.isRunning = false
	}
}
