package main

import (
	"day_8/hotel"
	"fmt"
	"sync"
)

//var wg sync.WaitGroup
var users []*hotel.User

func main() {
	userOne := hotel.NewUser("Chinmay J Parkhi", "9823378652")
	userTwo := hotel.NewUser("Keyur Parkhi", "7689901234")
	userThree := hotel.NewUser("Rajesh M", "9989901234")
	users = append(users, userOne, userTwo, userThree)
	userOne.NewItemForUser("C", 312, 120, 111)
	//userTwo.NewItemForUser("D", 302, 180, 181)

	wg := &sync.WaitGroup{}
	for index, val := range users {
		wg.Add(1)
		go hotel.PlaceOrder(val, uint(index)+1, uint(index)+100, uint(index)+160, "B", wg)
	}
	wg.Wait()
	userOne.GetOrder()
	userTwo.GetOrder()
	userThree.GetOrder()
	generateBill := make(chan struct{})
	fmt.Println("--------------- Invoice for users ---------")

	go userOne.InvoiceForUser(generateBill)
	generateBill <- struct{}{}
	go userTwo.InvoiceForUser(generateBill)
	generateBill <- struct{}{}
	go userThree.InvoiceForUser(generateBill)
	generateBill <- struct{}{}
}
