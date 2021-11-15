package main

import (
	b "bank/accountHolder"
	"fmt"
)

func main() {
	accOne := b.NewAccountHolder("Chinmay", "Parkhi", "7876675432", 21, 231, 1000, b.Male)
	accTwo := b.NewAccountHolder("Rajesh", "Patil", "9879901234", 22, 123, 1200, b.Male)
	str, err := accOne.ShareMoney(accTwo, 500)
	if err != nil {
		fmt.Println(str, err)
	} else {
		fmt.Println(str)
		accOne.GetAccountDetails()
		accTwo.GetAccountDetails()
	}
}
