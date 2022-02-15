package command

import (
	"system/device"
)

type onCommand struct {
	device device.Device
}

func NewOnCommand(device device.Device) onCommand {
	return onCommand{
		device: device,
	}
}

func (o onCommand) Execute() {
	o.device.On()
}
