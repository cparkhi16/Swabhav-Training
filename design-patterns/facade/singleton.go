package main

import "fmt"

type log struct {
	name string
}

var instance *log

func GetInstance() *log {

	if instance == nil {
		instance = &log{name: "Chinmay"}
	}

	return instance
}
func (l log) GetName() string {
	return l.name
}
func main() {
	i := GetInstance()
	fmt.Println(i.GetName())
}
