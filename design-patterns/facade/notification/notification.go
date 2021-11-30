package notification

import (
	l "design/ledger"
	"fmt"
)

type Notification struct {
	L *l.Ledger
}

func NewNotifier(l *l.Ledger) *Notification {
	return &Notification{L: l}
}
func (n *Notification) SendNotification() {
	if n.L.TransactionDetails[len(n.L.TransactionDetails)-1] != "" {
		fmt.Println(n.L.TransactionDetails[len(n.L.TransactionDetails)-1])
	} else {
		fmt.Println()
	}
}
