package main

import (
	"bankapp/walletfacade"
)

func main() {
	var str = make([]string, 10)
	wallet1 := walletfacade.New(123, 333, 2000, str, "")
	wallet1.CreditMoney(1000)
	wallet1.DebitMoney(1000)
}
