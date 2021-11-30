package main

import (
	"design/account"
	"design/ledger"
	"design/notification"
	"design/wallet"
	"log"
)

/*type Account struct {
	accountNumber uint
	securityCode  uint
	balance       uint
}

func NewAccount(accountNumber uint, balance uint) *Account {
	return &Account{accountNumber: accountNumber, balance: balance}
}

func (a *Account) CheckAccount(accNumber uint) bool {
	return a.accountNumber == accNumber
}

func NewSecurityCode(acc *Account) uint {
	acc.securityCode = acc.accountNumber + 500
	return acc.securityCode
}

func (a *Account) CheckSecurityCode(code uint) bool {
	return a.securityCode == code
}

type Wallet struct {
	acc *Account
}

func NewWallet(acc *Account) *Wallet {
	return &Wallet{acc: acc}
}
func (w *Wallet) Credit(money uint, l *Ledger) {
	w.acc.balance = w.acc.balance + money
	s := fmt.Sprintf("%v added to your account ", money)
	//l.transactionDetails = append(l.transactionDetails, s)
	l.MakeEntry(s)
}
func (w *Wallet) Debit(money uint, l *Ledger) {
	w.acc.balance = w.acc.balance - money
	s := fmt.Sprintf("%v debited from your account ", money)
	//l.transactionDetails = append(l.transactionDetails, s)
	l.MakeEntry(s)
}

type Ledger struct {
	acc                Account
	transactionDetails []string
}

func NewLedger(acc Account) *Ledger {
	return &Ledger{acc: acc}
}
func (l *Ledger) PrintLedger(accNumber uint) {
	if l.acc.accountNumber == accNumber {
		fmt.Printf("Transaction details for this account number %v \n", l.acc.accountNumber)
		fmt.Println(l.transactionDetails)
	} else {
		fmt.Println("No account found")
	}
}
func (l *Ledger) MakeEntry(s string) {
	l.transactionDetails = append(l.transactionDetails, s)
}
func (a *Account) PrintDetails() {
	fmt.Println("==== Account details ====")
	fmt.Println("Your account number ", a.accountNumber)
	fmt.Println("Your account balance ", a.balance)
}

type Notification struct {
	l Ledger
}

func (n *Notification) sendNotification() {
	if n.l.transactionDetails[len(n.l.transactionDetails)-1] != "" {
		fmt.Println(n.l.transactionDetails[len(n.l.transactionDetails)-1])
	} else {
		fmt.Println()
	}
}*/
func main() {
	acc := account.NewAccount(100, 201)
	isAccount := acc.CheckAccount(100)
	if !isAccount {
		log.Fatal("No account found")
	}
	secCode := account.NewSecurityCode(acc)
	isSecCode := acc.CheckSecurityCode(secCode)
	if !isSecCode {
		log.Fatal("Security code didn't matched")
	}
	acc.PrintDetails()
	w := wallet.NewWallet(acc)
	l := ledger.NewLedger(*acc)
	w.Credit(100, l)
	acc.PrintDetails()
	l.PrintLedger(100)
	w.Debit(100, l)
	l.PrintLedger(100)
	acc.PrintDetails()
	notify := notification.Notification{L: *l}
	notify.SendNotification()
}
