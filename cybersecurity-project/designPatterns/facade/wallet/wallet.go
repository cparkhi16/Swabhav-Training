package wallet

type Wallet struct {
	balance int
}

func New(balance int) *Wallet {
	return &Wallet{
		balance: balance,
	}
}

func (w *Wallet) Credit(amount int) {
	w.balance = w.balance + amount
}

func (w *Wallet) Debit(amount int) {
	w.balance = w.balance - amount
}
