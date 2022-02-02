package model

type Employee struct {
	Name   string
	City   string
	Skills string
}

func NewEmployee(name string, city string, skills string) *Employee {
	return &Employee{Name: name, City: city, Skills: skills}
}
