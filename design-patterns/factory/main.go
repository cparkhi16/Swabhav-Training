package main

func main() {
	autofactory := &AutomobileFactory{}
	model := autofactory.Make(bmw)
	model.Start()
	model.Stop()
	modelTwo := autofactory.Make(tesla)
	modelTwo.Start()
	modelTwo.Stop()
}
