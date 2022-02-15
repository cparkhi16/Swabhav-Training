package automobile

import "fmt"

type tesla struct {
	modelName string
}

func NewTesla(modelName string) tesla {
	return tesla{
		modelName: modelName,
	}
}

func (t tesla) Start() {
	fmt.Println("tesla start")
}

func (t tesla) Stop() {
	fmt.Println("tesla stop")
}
