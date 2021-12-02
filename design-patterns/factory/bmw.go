package main

import "fmt"

type BMW struct {
	modelName string
}

func (b *BMW) Start() {
	fmt.Println("Starting  ...", b.modelName)
}

func (b *BMW) Stop() {
	fmt.Println("Stopping ...", b.modelName)
}
