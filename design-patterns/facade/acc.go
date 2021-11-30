package main

import (
	"design/account"
	"design/ledger"
	"design/notification"
	"design/wallet"
)

func main() {
	acc := account.NewAccount(100, 201)
	w := wallet.NewWallet(acc)
	l := ledger.NewLedger(*acc)
	notify := notification.NewNotifier(l)
	wfl := wallet.NewWalletFacade(acc, l, notify, w)
	wfl.AddMoney(100, 100)
	wfl.DeductMoney(100, 100)
}
