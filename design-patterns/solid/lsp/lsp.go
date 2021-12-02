package main

import "fmt"

type Human struct {
	name string
	age  uint
}

func (h Human) getName() string {
	return h.name
}

type Student struct {
	Human
	marks uint
}
type Teacher struct {
	Human
	course string
}

func (t Teacher) getName() string {
	return "Teacher: " + t.name
}

type Person interface {
	getName() string
}
type logger struct{}

func (l logger) Printer(p Person) {
	fmt.Println("Name of person is ", p.getName())
}

func main() {
	t := Teacher{course: "Maths"}
	t.name = "Raj"
	t.age = 35
	l := logger{}
	l.Printer(t)

	s := Student{marks: 55}
	s.name = "Manish"
	s.age = 22
	l.Printer(s)
}
