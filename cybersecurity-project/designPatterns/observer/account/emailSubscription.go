package account

import "fmt"

type emailSubscription struct {
}

func NewEmailSubscription() emailSubscription {
	return emailSubscription{}
}

func (e emailSubscription) balanceModified(a account) {
	fmt.Println("New Email for- ", a.username, ", -Balance modified. current balance is-", a.balance)
}
