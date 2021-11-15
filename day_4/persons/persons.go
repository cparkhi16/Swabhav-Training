package persons

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       uint8
	address   []address
}
type address struct {
	roomNo uint8
	city   string
	state  string
}
type employee struct {
	empId uint8
	person
}

func New(fName, lname, city, state string, age, rNo uint8) *person {
	return &person{firstName: fName, lastName: lname, age: age, address: []address{{
		roomNo: rNo,
		city:   city,
		state:  state,
	},
	}}
}
func NewEmployee(empId uint8, p person) *employee {
	return &employee{empId: empId, person: p}
}
func (e *employee) UpdateEmployeeFirstName(uFirstName string) {
	e.person.firstName = uFirstName
}
func PrintPerson(p person) {
	fmt.Println("Person name ", p.firstName+" "+p.lastName)
	fmt.Println("Person age ", p.age)
	fmt.Println("Address of person :", p.address)
}
func PrintEmployeeDetails(e employee) {
	fmt.Println("Employee ID ---------------", e.empId)
	fmt.Println("Employee name ", e.firstName+" "+e.lastName)
	fmt.Println("Employee age ", e.age)
	fmt.Println("Address of employee :", e.address)
}
