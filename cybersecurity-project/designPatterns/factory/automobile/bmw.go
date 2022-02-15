package automobile

import "fmt"

type bmw struct {
	modelName string
}

func NewBmw(modelName string) bmw {
	return bmw{
		modelName: modelName,
	}
}

func (b bmw) Start() {
	fmt.Println("BMW start")
}

func (b bmw) Stop() {
	fmt.Println("BMW stop")
}
