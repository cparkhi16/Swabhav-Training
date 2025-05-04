package main

import "fmt"

type Button interface {
	Render()
}

type Checkbox interface {
	Toggle()
}

type WindowsButton struct {
}

func (w *WindowsButton) Render() {
	fmt.Println(" Rendering windows btn ")
}

type MacOSButton struct {
}

func (m *MacOSButton) Render() {
	fmt.Println(" Rendering mac btn ")
}

type WindowsCheckbox struct{}

func (wc *WindowsCheckbox) Toggle() {
	fmt.Println(" toggling windows checkbox ")
}

type MacOSCheckbox struct{}

func (mc *MacOSCheckbox) Toggle() {
	fmt.Println(" toggling macos checkbox ")
}

type GUIFactory interface {
	createButton() Button
	createCheckBox() Checkbox
}

type WindowsGUIFactory struct {
}

func (wgf *WindowsGUIFactory) createButton() Button {
	return &WindowsButton{}
}

func (wgf *WindowsGUIFactory) createCheckBox() Checkbox {
	return &WindowsCheckbox{}
}

type MacOSGUIFactory struct {
}

func (mgf *MacOSGUIFactory) createButton() Button {
	return &MacOSButton{}
}
func (mgf *MacOSGUIFactory) createCheckBox() Checkbox {
	return &MacOSCheckbox{}
}

func main() {
	var factory GUIFactory

	os := "windows"

	if os == "windows" {
		factory = &WindowsGUIFactory{}
	} else {
		factory = &MacOSGUIFactory{}
	}

	factory.createButton().Render()
	factory.createCheckBox().Toggle()
}
