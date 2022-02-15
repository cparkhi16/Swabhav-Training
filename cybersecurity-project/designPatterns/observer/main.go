package main

import (
	"bank/account"
	"fmt"
)

func main() {
	shan := account.NewAccount(123, "shan", "shan@yahoo.com", 1000)
	shan.AddSubscriber(account.NewEmailSubscription())
	shan.AddSubscriber(account.NewSmsSubscription())
	shan.Deposit(2000)
	shan.Withdraw(1000)
	fmt.Println("///////")
	shan.RemoveSubscriber(account.NewEmailSubscription())
	shan.Deposit(2000)
}
