package ledger

import (
	a "design/account"
	"fmt"
)

type Ledger struct {
	acc                a.Account
	TransactionDetails []string
}

func NewLedger(acc a.Account) *Ledger {
	return &Ledger{acc: acc}
}
func (l *Ledger) PrintLedger(accNumber uint) {
	if l.acc.AccountNumber == accNumber {
		fmt.Printf("Transaction details for this account number %v \n", l.acc.AccountNumber)
		fmt.Println(l.TransactionDetails)
	} else {
		fmt.Println("No account found")
	}
}
func (l *Ledger) MakeEntry(s string) {
	l.TransactionDetails = append(l.TransactionDetails, s)
}
