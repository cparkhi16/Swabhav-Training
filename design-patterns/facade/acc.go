package main

// Facade is used when we want to build a kind of interactive wall to communicate a big subsystem behind it
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
	wfl := wallet.NewWalletFacade(l, notify, w)
	wfl.AddMoney(100, 100)
	wfl.DeductMoney(100, 100)
	wfl.AddMoney(100, 300)
}
