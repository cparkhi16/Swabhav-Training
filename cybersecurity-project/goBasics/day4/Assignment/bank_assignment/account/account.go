package account

type Account struct {
	accountNo       int
	accountHolderId int
	balance         int
}

func New(accountNo int, accountHolderId int, balance int) *Account {
	return &Account{
		accountNo:       accountNo,
		accountHolderId: accountHolderId,
		balance:         balance,
	}
}
func (a *Account) GetBalance() int {

	return a.balance
}

func (a *Account) SetBalance(newBalance int) {
	a.balance = newBalance
}

func (a *Account) GetAccountHolderId() int {
	return a.accountHolderId
}

func (a *Account) SetAccountHolderId(newId int) {
	a.accountHolderId = newId
}

func (a *Account) GetAccountNo() int {
	return a.accountNo
}

func (a *Account) SetAccountNo(newAccountNo int) {
	a.accountNo = newAccountNo
}
