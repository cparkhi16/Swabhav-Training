package main

// Singleton is used when we want to instantiate a object only one time in the program and that same object is to be used by other classes
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
