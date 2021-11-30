package wallet

import (
	a "design/account"
	l "design/ledger"
	n "design/notification"
	"fmt"
	"log"
)

type Wallet struct {
	acc *a.Account
}
type WalletFacade struct {
	w      *Wallet
	ledger *l.Ledger
	notify *n.Notification
}

func NewWalletFacade(l *l.Ledger, n *n.Notification, w *Wallet) *WalletFacade {
	return &WalletFacade{ledger: l, notify: n, w: w}
}
func (wf *WalletFacade) AddMoney(accNo, money uint) {
	isAccount := wf.w.acc.CheckAccount(accNo)
	if !isAccount {
		log.Fatal("No account found")
	}
	secCode := a.NewSecurityCode(wf.w.acc)
	isSecCode := wf.w.acc.CheckSecurityCode(secCode)
	if !isSecCode {
		log.Fatal("Security code didn't matched")
	}
	wf.w.Credit(money, wf.ledger)
	wf.w.acc.PrintDetails()
	wf.ledger.PrintLedger(wf.w.acc.AccountNumber)
	wf.notify.SendNotification()
}
func (wf *WalletFacade) DeductMoney(accNo, money uint) {
	isAccount := wf.w.acc.CheckAccount(accNo)
	if !isAccount {
		log.Fatal("No account found")
	}
	secCode := a.NewSecurityCode(wf.w.acc)
	isSecCode := wf.w.acc.CheckSecurityCode(secCode)
	if !isSecCode {
		log.Fatal("Security code didn't matched")
	}
	wf.w.Debit(money, wf.ledger)
	wf.w.acc.PrintDetails()
	wf.ledger.PrintLedger(wf.w.acc.AccountNumber)
	wf.notify.SendNotification()
}
func NewWallet(acc *a.Account) *Wallet {
	return &Wallet{acc: acc}
}
func (w *Wallet) Credit(money uint, l *l.Ledger) {
	w.acc.Balance = w.acc.Balance + money
	s := fmt.Sprintf("%v added to your account ", money)
	l.MakeEntry(s)
}
func (w *Wallet) Debit(money uint, l *l.Ledger) {
	w.acc.Balance = w.acc.Balance - money
	s := fmt.Sprintf("%v debited from your account ", money)
	l.MakeEntry(s)
}
