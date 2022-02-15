package button

import (
	"system/command"
)

type button struct {
	command command.Command
}

func NewButton(command command.Command) *button {
	return &button{
		command: command,
	}
}

func (b *button) Press() {
	b.command.Execute()
}
