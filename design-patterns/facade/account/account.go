package account

import "fmt"

type Account struct {
	AccountNumber uint
	securityCode  uint
	Balance       uint
}

func NewAccount(accountNumber uint, balance uint) *Account {
	return &Account{AccountNumber: accountNumber, Balance: balance}
}

func (a *Account) CheckAccount(accNumber uint) bool {
	return a.AccountNumber == accNumber
}

func NewSecurityCode(acc *Account) uint {
	acc.securityCode = acc.AccountNumber + 500
	return acc.securityCode
}

func (a *Account) CheckSecurityCode(code uint) bool {
	return a.securityCode == code
}

func (a *Account) PrintDetails() {
	fmt.Println("==== Account details ====")
	fmt.Println("Your account number ", a.AccountNumber)
	fmt.Println("Your account balance ", a.Balance)
}
