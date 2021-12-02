package account

import (
	"fmt"
)

type EmailSubscription struct {
	id string
}

func NewEmailSubscription(id string) *EmailSubscription {
	return &EmailSubscription{id: id}
}
func (e *EmailSubscription) BalanceModified(a Account) {
	fmt.Printf("Sending email to customer %s for balance modification : New balance is - %v \n ", a.UserName, a.Balance)
}
func (e *EmailSubscription) GetID() string {
	return e.id
}
