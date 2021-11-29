package main

import (
	"fmt"
)

type Sport interface{
	sportName() string
}
type Address struct{
	street string
	city string
	state string
}
type Human struct{
	name string //Encapsulation
	sport string
	address Address//Compostion Human HAS-A address
}
type Employee struct{
	name string //Encapsulation
	sport string
	address Address//Compostion Employee HAS-A address
}
func NewHuman(name,sport string)*Human{
	return &Human{name:name,sport:sport}
}
func (h Human) sportName() string{ //Defining methods for struct Human as we declare methods for classes
	return h.name + " plays " + h.sport + "."
}
func NewEmployee(name,sport string)*Employee{
	return &Employee{name:name,sport:sport}
}
func (e Employee) sportName() string{ //Defining methods for struct Employee as we declare methods for classes
	return e.name + " plays " + e.sport + "."
}

func PrintSportName(i Sport){// Polymorphism
	fmt.Println(i.sportName())
}
func main() {
	
	human1 := NewHuman("Chinmay","cricket")//Calling Constructors
	human1.address.city="Dombivli" //Composition
	fmt.Println(*human1)
	PrintSportName(human1)
	human2 :=  NewHuman("Rahul","football")
	PrintSportName(human2)
	e:=NewEmployee("Mahesh","chess")
	PrintSportName(e)
}
