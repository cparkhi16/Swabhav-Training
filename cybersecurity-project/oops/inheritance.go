package main

import "fmt"

type user struct {
	name    string
	age     int
	address //composition
}

type address struct {
	roomNo int
	street string
}

func main() {
	var user1 user
	user1.name = "shan"
	user1.age = 34
	user1.roomNo = 23
	user1.street = "ss"
	fmt.Println(user1)

}
