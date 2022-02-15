package account

type account struct {
	number        uint16
	username      string
	email         string
	balance       uint16
	subscriptions []subscriber
}

func NewAccount(number uint16, username string, email string, balance uint16) *account {
	return &account{
		number:   number,
		username: username,
		email:    email,
		balance:  balance,
	}
}

func (a *account) AddSubscriber(s subscriber) {
	a.subscription = append(a.subscriptions, s)
}

func (a *account) RemoveSubscriber(s subscriber) {
	indexToRemove := 0
	for i, v := range a.subscriptions {
		if v == s {
			indexToRemove = i
			break
		}
	}
	a.subscriptions[indexToRemove] = a.subscriptions[len(a.subscriptions)-1]
	a.subscriptions = a.subscriptions[:len(a.subscriptions)-1]
}

func (a *account) Deposit(amount uint16) {
	a.balance = a.balance + amount
	for _, v := range a.subscriptions {
		v.balanceModified(*a)
	}
}

func (a *account) Withdraw(amount uint16) {
	a.balance = a.balance - amount
	for _, v := range a.subscriptions {
		v.balanceModified(*a)
	}
}
