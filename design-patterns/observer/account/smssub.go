package account

import (
	"fmt"
)

type SMSSubscription struct {
	id string
}

func NewSMSSubscription(id string) *SMSSubscription {
	return &SMSSubscription{id: id}
}
func (s *SMSSubscription) BalanceModified(a Account) {
	fmt.Printf("Sending SMS to customer %s for balance modification : New balance is - %v\n ", a.UserName, a.Balance)
}
func (s *SMSSubscription) GetID() string {
	return s.id
}
