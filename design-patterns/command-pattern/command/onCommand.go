package command

import d "appliance/device"

type OnCommand struct {
	device d.Device
}

func NewOnCommand(device d.Device) *OnCommand {
	return &OnCommand{device: device}
}
func (o *OnCommand) Execute() {
	o.device.On()
}
