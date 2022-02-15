package person

import (
	"fmt"

	address "struct/address"
)

type person struct {
	firstname string
	lastname  string
	age       int
	gender    Gender
	addresses []*address.Address
}

type Gender string

const (
	MALE   Gender = "male"
	FEMALE Gender = "female"
)

//embedded struct example
type employee struct {
	empId     int
	firstname string
	person    *person
}

func NewEmployee(empId int, firstname string, person *person) *employee {
	return &employee{
		empId:     empId,
		firstname: firstname,
		person:    person,
	}
}

func (e *employee) Display() {
	fmt.Println("empId-", e.empId, "firstname-", e.firstname)
	e.person.Display()
}

func (e *employee) GetEmpId() int {
	return e.empId
}

func (e *employee) SetEmpId(newEmpId int) {
	e.empId = newEmpId
}

func (e *employee) GetFirstName() string {
	return e.firstname
}

func (e *employee) SetFirstName(newFirstName string) {
	e.firstname = newFirstName
}

func (e *employee) GetPerson() *person {
	return e.person
}

func (e *employee) SetPerson(newPerson *person) {
	e.person = newPerson
}

func New(firstname, lastname string, age int, gender Gender, roomNo int, city string, state string) *person {

	return &person{
		firstname: firstname,
		lastname:  lastname,
		age:       age,
		gender:    gender,
		addresses: []*address.Address{address.New(roomNo, city, state)},
	}
}

func (p *person) GetGender() Gender {
	return p.gender
}

func (p *person) SetGender(newgender Gender) {
	p.gender = newgender
}

func (p *person) GetFistName() string {
	return p.firstname
}

func (p *person) SetFirstName(newFirstName string) {
	p.firstname = newFirstName
}

func (p *person) GetLastName() string {
	return p.lastname
}

func (p *person) SetLastName(newLastName string) {
	p.lastname = newLastName
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetAge(newAge int) {
	p.age = newAge
}

func (p *person) GetAddresses() []*address.Address {
	return p.addresses
}

func (p *person) AddAddress(roomNo int, city string, state string) {
	p.addresses = append(p.addresses, address.New(roomNo, city, state))
}

func (p *person) ShowAddresses() {
	for i, v := range p.addresses {
		fmt.Println(i, *v)
	}
}

func (p *person) RemoveAddress(index int) {
	p.addresses = remove(p.addresses, index)
}

func remove(slice []*address.Address, s int) []*address.Address {
	return append(slice[:s], slice[s+1:]...)
}

func (p *person) Display() {
	fmt.Println("firstname-", p.firstname, " lastname-", p.lastname, " age-", p.age)
	for i, v := range p.addresses {
		fmt.Println("address", i, "-", *v)
	}
}
