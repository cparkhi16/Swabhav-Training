package main

import "fmt"

// Button interface
type Button interface {
	Render()
}

// Checkbox interface
type Checkbox interface {
	Toggle()
}

// WindowsButton struct
type WindowsButton struct{}

func (w WindowsButton) Render() {
	fmt.Println("Rendering Windows button")
}

// MacOSButton struct
type MacOSButton struct{}

func (m MacOSButton) Render() {
	fmt.Println("Rendering MacOS button")
}

// WindowsCheckbox struct
type WindowsCheckbox struct{}

func (w WindowsCheckbox) Toggle() {
	fmt.Println("Toggling Windows checkbox")
}

// MacOSCheckbox struct
type MacOSCheckbox struct{}

func (m MacOSCheckbox) Toggle() {
	fmt.Println("Toggling MacOS checkbox")
}

// GUIFactory interface
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// WindowsFactory struct
type WindowsFactory struct{}

func (w WindowsFactory) CreateButton() Button {
	return WindowsButton{}
}

func (w WindowsFactory) CreateCheckbox() Checkbox {
	return WindowsCheckbox{}
}

// MacOSFactory struct
type MacOSFactory struct{}

func (m MacOSFactory) CreateButton() Button {
	return MacOSButton{}
}

func (m MacOSFactory) CreateCheckbox() Checkbox {
	return MacOSCheckbox{}
}

func main() {
	var factory GUIFactory

	// Simulating platform-based factory selection
	os := "windows" // Change to "macos" for MacOSFactory
	if os == "windows" {
		factory = WindowsFactory{}
	} else {
		factory = MacOSFactory{}
	}

	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	button.Render()
	checkbox.Toggle()
}
