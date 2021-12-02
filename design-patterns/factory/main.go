package main

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
