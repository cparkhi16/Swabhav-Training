package main

// Observer pattern is used when one object wants to notify changes to other object (here account's object is modified and it wants to notify this change to the email and sms subscribers)
import (
	"fmt"
	a "observer/account"
)

func main() {

	acc := a.New("Chinmay", "cp@gmail.com", 231, 1000)

	email := a.NewEmailSubscription("email")
	acc.AddSubscription(email)
	acc.Deposit(200)
	acc.NotifyAll()
	sms := a.NewSMSSubscription("sms")
	acc.AddSubscription(sms)
	acc.WithDraw(100)
	acc.NotifyAll()
	acc.RemoveSubscription(sms)
	fmt.Println("After removing sms subscription")
	acc.NotifyAll()
}
