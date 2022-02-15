package command

import (
	"system/device"
)

type offCommand struct {
	device device.Device
}

func NewOffCommand(device device.Device) offCommand {
	return offCommand{
		device: device,
	}
}

func (o offCommand) Execute() {
	o.device.Off()
}
