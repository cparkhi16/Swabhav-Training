package main

import (
	b "bank/accountHolder"
	l "bank/logger"
	"fmt"
)

func CheckError(err l.BankError) bool {
	return err.Error() == ""
}
func main() {
	userOne, err := b.NewAccountHolder("Chinmay", "Parkhi", "7876675432", 21, 231, 400, 1000, b.Male)

	e := CheckError(err)
	fmt.Println(err.Error())
	userTwo, err := b.NewAccountHolder("Rajesh", "Patil", "9879901234", 22, 123, 500, 1200, b.Male)

	er := CheckError(err)
	fmt.Println(err.Error())

	if er && e {
		userOne.AddAccountForUser(232, 5000)
		userTwo.AddAccountForUser(124, 100)
		str, err := userOne.ShareMoney(userTwo, 232, 123, 1000)
		if err.Error() != "" {
			fmt.Println(str, err)
		} else {
			fmt.Println(str)
			userOne.GetAccountDetails()
			userTwo.GetAccountDetails()
		}
	}

}
