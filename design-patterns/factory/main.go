package main

// Factory pattern is used when we want to create any object avaialable to be created which implements a interface.
func main() {
	autofactory := &AutomobileFactory{}
	model := autofactory.Make(bmw)
	if model != nil {
		model.Start()
		model.Stop()
	}
	modelTwo := autofactory.Make(tesla)
	if modelTwo != nil {
		modelTwo.Start()
		modelTwo.Stop()
	}
}
