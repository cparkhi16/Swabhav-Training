package main

import (
	"day_4/persons"
	"fmt"
)

func main() {
	personOne := NewPerson("Chinmay", "Parkhi", 23)
	fmt.Println(*personOne)
	//FuncPassStruct(person)
	var personsNew []person
	personsNew = append(personsNew, *personOne)
	new_person := NewPerson("Ram", "xyz", 21)
	personsNew = append(personsNew, *new_person)
	for _, val := range personsNew {
		fmt.Println("Struct value", val)
	}
	manipulateStruct(*personOne)
	fmt.Println("after manipulating personOne", *personOne)
	nPerson := *personOne
	myPtr := &nPerson
	manipulateStructWithPtr(myPtr)
	fmt.Println("after manipulating personOne with ptr", *personOne)

	//package persons
	p1 := persons.NewPerson("Manoj", "P", 21)
	persons.PrintPerson(*p1)
	e1 := persons.NewEmployee(123, *p1)
	persons.PrintEmployeeDetails(*e1)
	//Updating employee first Name
	e1.SetEmployeeFirstName("Raj")
	e1.SetEmployeeAddress(4, "Kalyan", "Maharashtra")
	persons.PrintEmployeeDetails(*e1)
	name := e1.GetEmployeeFirstName(123)
	fmt.Println(name)

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
