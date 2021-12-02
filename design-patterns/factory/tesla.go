package main

import "fmt"

type Tesla struct {
	modelName string
}

func (t *Tesla) Start() {
	fmt.Println("Starting  ...", t.modelName)
}

func (t *Tesla) Stop() {
	fmt.Println("Stopping ...", t.modelName)
}
