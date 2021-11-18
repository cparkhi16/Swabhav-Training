package accountHolder

import (
	l "bank/logger"
	"fmt"
)

type account struct {
	accountNumber uint8
	balance       int
}
type person struct {
	id        int
	firstName string
	lastName  string
	age       uint8
	contact   string
	Gender    Gender
	accounts  []account
}
type Gender int

const (
	Male   Gender = 1
	Female Gender = 2
	Other  Gender = 3
)

func (p *person) AddAccountForUser(accountNumber uint8, balance int) {
	acc := account{
		balance:       balance,
		accountNumber: accountNumber,
	}
	p.accounts = append(p.accounts, acc)
	fmt.Println("Account added", *p)

}
func NewAccountHolder(fName, lName, contact string, age, accountNumber uint8, id, balance int, gender Gender) (p *person, e l.BankError) {
	if balance <= 0 {
		e = l.BankError{Err: "zero or negative balance not accepted"}
		return
	} else if gender.EnumIndex() > 3 {
		e = l.BankError{Err: "please provide appropriate gender"}
		return
	}
	acc := account{
		accountNumber: accountNumber,
		balance:       balance,
	}

	p = &person{firstName: fName, lastName: lName, age: age, contact: contact, Gender: gender, accounts: []account{acc}, id: id}
	return p, l.BankError{Err: ""}
}
func (g Gender) String() string {
	return [...]string{"Male", "Female", "Other"}[g-1]
}
func (g Gender) EnumIndex() int {
	return int(g)
}
func (a *person) GetAccountDetails() {
	fmt.Println("Account Number", a.accounts)
	fmt.Println("Account holder name ", a.firstName+" "+a.lastName)
	fmt.Println("Account balance", a.accounts)
	fmt.Println("Contact Number", a.contact)
	fmt.Println("Account holder age", a.age)
	fmt.Println("Gender :", a.Gender.String())
	fmt.Println("-----------------------------------------")
}
func (acc *person) isAccountPresent(accountNumber uint8) l.BankError {
	for _, val := range acc.accounts {
		if val.accountNumber == accountNumber {
			return l.BankError{Err: ""}

		}
	}
	return l.BankError{Err: "Account does not exist"}
}
func (acc *person) GetBalance(accNumber uint8) int {
	for _, val := range acc.accounts {
		if val.accountNumber == accNumber {
			return val.balance
		}
	}
	return 0
}
func (acc *person) UpdateAccount(accountNumber uint8, balance int) l.BankError {
	if balance < 0 {
		return l.BankError{Err: "Negative balance not accepted"}
	}
	for i, val := range acc.accounts {
		if val.accountNumber == accountNumber {
			acc.accounts[i].balance = balance
		}
	}
	return l.BankError{Err: ""}
}
func (acc *person) ShareMoney(accTwo *person, accountNumberOne, accountNumberTwo uint8, money int) (string, l.BankError) {
	if money < 0 {
		return "Transfer unsuccessful", l.BankError{Err: "negative money can't be transfered"}
	}
	err1 := acc.isAccountPresent(accountNumberOne)
	fmt.Println(err1.Error())
	err2 := accTwo.isAccountPresent(accountNumberTwo)
	fmt.Println(err2.Error())
	if err1.Error() != "" || err2.Error() != "" {
		return "Transfer unsuccessful", l.BankError{Err: "Please enter valid account details"}
	} else {
		acc.UpdateAccount(accountNumberOne, acc.GetBalance(accountNumberOne)-money)
		accTwo.UpdateAccount(accountNumberTwo, accTwo.GetBalance(accountNumberTwo)+money)
	}
	//fmt.Println(*acc, *accTwo)
	return "Transfer successful", l.BankError{Err: ""}
}
