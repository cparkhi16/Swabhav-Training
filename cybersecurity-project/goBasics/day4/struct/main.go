package main

import (
	"fmt"
	person "struct/person"
)

func main() {
	shan := person.New("shan", "mayer", 43, person.MALE, 4, "xyz", "lmn")
	fmt.Println(shan.GetFistName())
	shan.SetFirstName("john")
	shan.Display()

	fmt.Println("gender-", shan.GetGender())
	shan.SetGender(person.FEMALE)
	fmt.Println("gender-", shan.GetGender())

	shan.AddAddress(23, "adc", "def")
	shan.Display()

	shan.RemoveAddress(0)
	shan.Display()

	fmt.Println(shan.GetAddresses()[0].GetRoomNo())
	shan.GetAddresses()[0].SetRoomNo(334)
	shan.GetAddresses()[0].SetCity("ddd")
	shan.Display()

	//embedded struct example
	emp1 := person.NewEmployee(23, "red", shan)
	emp1.Display()
	fmt.Println(emp1.GetFirstName())
	emp1.SetFirstName("blue")
	emp1.Display()
	fmt.Println(emp1.GetPerson().GetFistName())
}
