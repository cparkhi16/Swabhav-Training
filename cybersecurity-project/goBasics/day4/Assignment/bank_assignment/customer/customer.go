package customer

//package name and struct name should be same, while importing we can ignore the name then.
//acronyms should be in capital for eg. ID not Id
import (
	account "bank/account"
	customError "bank/customError"
	"fmt"
	"strconv"
)

type customer struct {
	id          int
	name        string
	accountList []*account.Account
}

func New(id int, name string, accountNo int, balance int) *customer {
	return &customer{
		id:          id,
		name:        name,
		accountList: []*account.Account{account.New(accountNo, id, balance)},
	}
}

func (c *customer) DisplayDetails() {
	fmt.Printf("ID-%d ", c.id)
	fmt.Printf("Name-%s\n", c.name)
	for i, a := range c.accountList {
		fmt.Printf("Account-%d --->", i)
		fmt.Printf("AccountNo-%d ", a.GetAccountNo())
		fmt.Printf("Account Balance-%d\n", a.GetBalance())
	}
}

func (c *customer) AddAccount(accountNo int, balance int) {
	c.accountList = append(c.accountList, account.New(accountNo, c.id, balance))
}

func (c *customer) GetAccountFromNo(no int) (*account.Account, *customError.CustomError) {
	for _, v := range c.accountList {
		if v.GetAccountNo() == no {
			return v, nil
		}
	}
	return nil, customError.New("Account no-"+strconv.Itoa(no)+" of user "+c.name+" is not found", customError.FATAL)
}

func (c *customer) GetFirstName() string {
	return c.name
}

func (c *customer) SetFirstName(newName string) {
	c.name = newName
}

func (c *customer) TransferMoney(toCustomer *customer, senderaccountNo int, receiveraccountno int, amount int) *customError.CustomError { //we can return error as well to make it more general
	senderAccount, err := c.GetAccountFromNo(senderaccountNo)
	if err != nil {
		return err
	}
	if senderAccount.GetBalance() >= amount && amount > 0 {
		receiverAccount, err := toCustomer.GetAccountFromNo(receiveraccountno)
		if err != nil {
			return err
		}
		receiverAccount.SetBalance(receiverAccount.GetBalance() + amount)
		senderAccount.SetBalance(senderAccount.GetBalance() - amount)
	} else {
		return customError.New("Account balance of user "+c.name+" is not sufficient", customError.FATAL)
	}
	return nil
}
