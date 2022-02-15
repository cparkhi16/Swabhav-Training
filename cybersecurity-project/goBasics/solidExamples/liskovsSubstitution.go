package main

import "fmt"

type animal interface {
	getName() string
}

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	degree string
}

type student struct {
	human
	marks int
}

func (h student) getName() string {
	return "student"
}

func printDetails(a animal) {
	fmt.Println(a.getName())
}

//Liskovs Substitution principle-->
/*child can be substituted at a place of parent. so child should implement all the methods
that are defined on parent struct. In golang we don't have inheritance, we use composition instead so violating
this principle of liskovs substitution is not possible. All the methods of parent struct can be called from its child struct
*/
func main() {
	//shan := human{name: "shan"}
	/*
		john := teacher{
			human:  human{name: "john"},
			degree: "n",
		}*/
	kyle := student{
		human: human{name: "kyle"},
		marks: 68,
	}
	fmt.Println(kyle.getName())
	fmt.Println(kyle.human.getName())
	//printDetails(shan)
	//printDetails(john)
	//printDetails(kyle)
}
