package command

import d "appliance/device"

type OffCommand struct {
	device d.Device
}

func NewOffCommand(device d.Device) *OffCommand {
	return &OffCommand{device: device}
}
func (o *OffCommand) Execute() {
	o.device.Off()
}
