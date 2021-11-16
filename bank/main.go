package main

import (
	b "bank/accountHolder"
	"fmt"
)

func PrintError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func main() {
	accOne, err := b.NewAccountHolder("Chinmay", "Parkhi", "7876675432", 21, 231, 1000, b.Male)
	e := PrintError(err)
	accTwo, err := b.NewAccountHolder("Rajesh", "Patil", "9879901234", 22, 123, 1200, b.Male)
	er := PrintError(err)
	if er && e {
		str, err := accOne.ShareMoney(accTwo, -1000)
		if err != nil {
			fmt.Println(str, err)
		} else {
			fmt.Println(str)
			accOne.GetAccountDetails()
			accTwo.GetAccountDetails()
		}
	}

}
