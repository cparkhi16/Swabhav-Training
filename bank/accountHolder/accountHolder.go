package accountHolder

import (
	"fmt"
)

type person struct {
	firstName     string
	lastName      string
	age           uint8
	balance       int
	contact       string
	accountNumber uint8
	Gender        Gender
}
type Gender int

const (
	Male   Gender = 1
	Female Gender = 2
	Other  Gender = 3
)

func NewAccountHolder(fName, lName, contact string, age, accountNumber uint8, balance int, gender Gender) (*person, error) {
	if balance <= 0 {
		return nil, fmt.Errorf("zero or negative balance not accepted")
	}
	return &person{firstName: fName, lastName: lName, age: age, contact: contact, balance: balance, accountNumber: accountNumber, Gender: gender}, nil
}
func (g Gender) String() string {
	return [...]string{"Male", "Female", "Other"}[g-1]
}
func (a person) GetAccountDetails() {
	fmt.Println("Account Number", a.accountNumber)
	fmt.Println("Account holder name ", a.firstName+" "+a.lastName)
	fmt.Println("Account balance", a.balance)
	fmt.Println("Contact Number", a.contact)
	fmt.Println("Account holder age", a.age)
	fmt.Println("Gender :", a.Gender.String())
	fmt.Println("-----------------------------------------")
}

func (acc *person) ShareMoney(accTwo *person, money int) (string, error) {
	if acc.balance < money {
		return "Transfer unsuccessful", fmt.Errorf("you have less balance")
	}
	acc.balance = acc.balance - money
	accTwo.balance = accTwo.balance + money
	return "Transfer successful", nil
}
