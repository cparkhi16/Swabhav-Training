package main

import (
	"system/button"
	"system/command"
	"system/device"
)

func main() {
	sonyTv := device.NewTv()
	redButton := button.NewButton(command.NewOnCommand(sonyTv))
	redButton.Press()
	greenButton := button.NewButton(command.NewOffCommand(sonyTv))
	greenButton.Press()
	greenButton.Press()

	nokia := device.NewMobile()
	blueButton := button.NewButton(command.NewOnCommand(nokia))
	blueButton.Press()
	yellowButton := button.NewButton(command.NewOffCommand(nokia))
	yellowButton.Press()

}
