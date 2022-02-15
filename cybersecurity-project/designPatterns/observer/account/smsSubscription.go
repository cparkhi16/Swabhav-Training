package account

import "fmt"

type smsSubscription struct {
}

func NewSmsSubscription() smsSubscription {
	return smsSubscription{}
}

func (e smsSubscription) balanceModified(a account) {
	fmt.Println("New Email for- ", a.username, ", -Balance modified. current balance is-", a.balance)
}
