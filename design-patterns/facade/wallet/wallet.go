package wallet

import (
	a "design/account"
	l "design/ledger"
	"fmt"
)

type Wallet struct {
	acc *a.Account
}

func NewWallet(acc *a.Account) *Wallet {
	return &Wallet{acc: acc}
}
func (w *Wallet) Credit(money uint, l *l.Ledger) {
	w.acc.Balance = w.acc.Balance + money
	s := fmt.Sprintf("%v added to your account ", money)
	//l.transactionDetails = append(l.transactionDetails, s)
	l.MakeEntry(s)
}
func (w *Wallet) Debit(money uint, l *l.Ledger) {
	w.acc.Balance = w.acc.Balance - money
	s := fmt.Sprintf("%v debited from your account ", money)
	//l.transactionDetails = append(l.transactionDetails, s)
	l.MakeEntry(s)
}
