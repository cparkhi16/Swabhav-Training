package main

import "fmt"

func main() {
	personOne := NewPerson("Chinmay", "Parkhi", 23)
	fmt.Println(*personOne)
	//FuncPassStruct(person)
	var persons []person
	persons = append(persons, *personOne)
	new_person := NewPerson("Ram", "xyz", 21)
	persons = append(persons, *new_person)
	for _, val := range persons {
		fmt.Println("Struct value", val)
	}
	manipulateStruct(*personOne)
	fmt.Println("after manipulating personOne", *personOne)
	nPerson := *personOne
	myPtr := &nPerson
	manipulateStructWithPtr(myPtr)
	fmt.Println("after manipulating personOne with ptr", *personOne)
}
func manipulateStruct(p person) {
	p.firstName = "Raju"
	fmt.Println("in func call of manipulateStruct ", p)
}
func manipulateStructWithPtr(p *person) {
	p.firstName = "Raju"
	fmt.Println("With ptr", p)
}
func NewPerson(fName, lName string, age uint8) *person {
	p := person{
		firstName: fName,
		lastName:  lName,
		age:       age,
	}
	return &p
}

type person struct {
	firstName string
	lastName  string
	age       uint8
}
