package account

type Subscriber interface {
	BalanceModified(a Account)
	GetID() string
}
