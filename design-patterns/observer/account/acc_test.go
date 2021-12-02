package account

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	var list = []struct {
		UserName string
		email    string
		number   uint16
		Balance  uint16
		Expected *Account
	}{{
		"Chinmay", "cp@gmail.com", 231, 1000, nil,
	}, {
		"Rajesh", "rp@gmail.com", 232, 2200, nil,
	}}
	for _, val := range list {
		actual := New(val.UserName, val.email, val.number, val.Balance)
		if val.Expected == actual {
			t.Errorf("Error found for New Account creation")
		}
	}
}

func TestDeposit(t *testing.T) {
	account := New("Chinmay", "cp@gmail.com", 231, 1000)
	var expected uint16 = 1200
	account.Deposit(200)
	actual := account.Balance
	if actual != expected {
		t.Errorf("Error found for Depositing amount")
	}
}
func TestWithDraw(t *testing.T) {
	account := New("Chinmay", "cp@gmail.com", 231, 1000)
	var expected uint16 = 800
	account.WithDraw(200)
	actual := account.Balance
	if actual != expected {
		t.Errorf("Error found for WithDrawing amount")
	}
}

func TestAddSubscriber(t *testing.T) {
	account := New("Chinmay", "cp@gmail.com", 231, 1000)
	email := NewEmailSubscription("email")
	account.AddSubscription(email)
	var expected int = 1
	actual := len(account.subscriptions)
	if actual != expected {
		t.Errorf("Error found for Adding subscriber")
	}
}

func TestRemoveSubscriber(t *testing.T) {
	account := New("Chinmay", "cp@gmail.com", 231, 1000)
	email := NewEmailSubscription("email")
	account.RemoveSubscription(email)
	var expected int = 0
	actual := len(account.subscriptions)
	if actual != expected {
		t.Errorf("Error found for Adding subscriber")
	}
}
