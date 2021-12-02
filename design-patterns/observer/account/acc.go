package account

type Account struct {
	subscriptions []Subscriber
	UserName      string
	email         string
	number        uint16
	Balance       uint16
}

func New(name, email string, number, balance uint16) *Account {
	return &Account{
		UserName: name,
		email:    email,
		Balance:  balance,
		number:   number,
	}
}

func (a *Account) AddSubscription(s Subscriber) {
	a.subscriptions = append(a.subscriptions, s)
}

func (a *Account) RemoveSubscription(s Subscriber) {
	a.subscriptions = removeFromslice(a.subscriptions, s)
}
func (a *Account) NotifyAll() {
	for _, observer := range a.subscriptions {
		observer.BalanceModified(*a)
	}
}

func removeFromslice(subList []Subscriber, observerToRemove Subscriber) []Subscriber {
	subListLength := len(subList)
	for i, observer := range subList {
		if observerToRemove.GetID() == observer.GetID() {
			subList[subListLength-1], subList[i] = subList[i], subList[subListLength-1]
			return subList[:subListLength-1]
		}
	}
	return subList
}

func (a *Account) Deposit(amt uint16) {
	a.Balance = a.Balance + amt
}
func (a *Account) WithDraw(amt uint16) {
	a.Balance = a.Balance - amt
}
