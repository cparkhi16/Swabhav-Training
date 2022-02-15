package main

import (
	customer "bank/customer"
	"fmt"
)

func main() {
	shan := customer.New(1, "shan", 123, 1000)
	john := customer.New(3, "john", 456, 3000)

	shan.AddAccount(333, 10000)
	shan.AddAccount(444, 6000)

	john.AddAccount(222, 3000)
	john.AddAccount(111, 2000)

	fmt.Println("Before Transaction...")
	shan.DisplayDetails()
	john.DisplayDetails()

	err := shan.TransferMoney(john, 333, 222, 1000)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("After Transaction...")
		shan.DisplayDetails()
		john.DisplayDetails()
	}

}
