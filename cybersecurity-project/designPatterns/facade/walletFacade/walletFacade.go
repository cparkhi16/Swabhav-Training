package walletfacade

import (
	"bankapp/account"
	"bankapp/ledger"
	"bankapp/notification"
	"bankapp/wallet"
	"fmt"
	"strconv"
)

type walletFacade struct {
	account      *account.Account
	securityCode *account.SecurityCode
	ledger       *ledger.Ledger
	notification *notification.Notification
	wallet       *wallet.Wallet
}

func New(accNo int, secCode int, balance int, entires []string, msg string) *walletFacade {
	return &walletFacade{
		account:      account.NewAccount(accNo),
		securityCode: account.NewCode(secCode),
		ledger:       ledger.New(entires),
		notification: notification.New(msg),
		wallet:       wallet.New(balance),
	}

}

func (w *walletFacade) CreditMoney(amount int) {
	w.account.CheckAccount(123)
	w.securityCode.CheckCode(333)
	w.wallet.Credit(1000)
	w.ledger.MakeEntry("money transfer to shanAcc-" + strconv.Itoa(1000))
	w.notification.SendNotification("money transfer to shanAcc" + strconv.Itoa(1000))
	fmt.Println(w.ledger.GetDetails())
}

func (w *walletFacade) DebitMoney(amount int) {
	w.account.CheckAccount(123)
	w.securityCode.CheckCode(333)
	w.wallet.Debit(1000)
	w.ledger.MakeEntry("money transfer from shanAcc-" + strconv.Itoa(1000))
	w.notification.SendNotification("money transfer from shanAcc" + strconv.Itoa(1000))
	fmt.Println(w.ledger.GetDetails())
}
