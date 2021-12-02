package main

// Command pattern is used when we want to execute commands(ON, OFF in this case) based on different objects (TV, mobile in this case)
import (
	c "appliance/command"
	d "appliance/device"
)

type Button struct {
	c c.Command
}

func NewButton(c c.Command) *Button {
	return &Button{c: c}
}
func (b *Button) Press() {
	b.c.Execute()
}

func main() {
	TV := d.NewTV()
	b := NewButton(c.NewOnCommand(TV))
	b.Press()
	Mobile := d.NewMobile()
	bt := NewButton(c.NewOnCommand(Mobile))
	bt.Press()
	bt.Press()
}
