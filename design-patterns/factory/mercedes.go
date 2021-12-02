package main

import "fmt"

type Mercedes struct {
	modelName string
}

func NewMercedes(modelName string) *Mercedes {
	return &Mercedes{modelName: modelName}
}
func (m *Mercedes) Start() {
	fmt.Println("Starting ...", m.modelName)
}

func (m *Mercedes) Stop() {
	fmt.Println("Stopping ...", m.modelName)
}
