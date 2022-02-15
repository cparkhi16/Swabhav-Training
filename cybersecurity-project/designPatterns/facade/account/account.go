package account

import "fmt"

type Account struct {
	accountNo int
}

func NewAccount(accountNo int) *Account {
	return &Account{
		accountNo: accountNo,
	}
}

func (a *Account) CheckAccount(accountNo int) {
	if a.accountNo == accountNo {
		fmt.Println("account number is equal")
	} else {
		fmt.Println("not equal")
	}
}

type SecurityCode struct {
	code int
}

func NewCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) CheckCode(securityCode int) {
	if s.code == securityCode {
		fmt.Println("code is equal")
	} else {
		fmt.Println("not equal")
	}
}
