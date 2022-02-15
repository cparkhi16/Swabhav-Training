package account

type subscriber interface {
	balanceModified(a account)
}
