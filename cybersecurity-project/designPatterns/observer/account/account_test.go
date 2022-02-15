package account

import "testing"

var actualAcc = NewAccount(123, "test", "test@test.com", 0)

func TestNewAccount(t *testing.T) {
	expectedAcc := account{
		number:   123,
		username: "test",
		email:    "test@test.com",
		balance:  0,
	}
	if actualAcc.number != expectedAcc.number || actualAcc.username != expectedAcc.username || actualAcc.email != expectedAcc.email || actualAcc.balance != expectedAcc.balance {
		t.Errorf("TestNewAccount-expected not equal to actual acocunt")
	}
}

func TestAddSubscriber(t *testing.T) {
	actualAcc.AddSubscriber(NewEmailSubscription())
	actuallen := len(actualAcc.subscriptions)
	expectedlen := 1
	if actualAcc.subscription[0] == NewEmailSubscription() {
		t.Errorf("TestAddSubscriber-expected not equal to actual acocunt")
	}
}

func TestRemoveSubscriber(t *testing.T) {
	actualAcc.RemoveSubscriber(NewEmailSubscription())
	actuallen := len(actualAcc.subscriptions)
	expectedlen := 0
	if actuallen != expectedlen {
		t.Errorf("TestRemoveSubscriber-expected not equal to actual acocunt")
	}
}

func TestDeposit(t *testing.T) {
	actualAcc.Deposit(1000)
	actualBalance := actualAcc.balance
	expectedBalance := uint16(1000)
	if actualBalance != expectedBalance {
		t.Errorf("TestDeposit-expected=%d not equal to actual=%d", expectedBalance, actualBalance)
	}
}

func TestWithdraw(t *testing.T) {
	actualAcc.Withdraw(1000)
	actualBalance := actualAcc.balance
	expectedBalance := uint16(0)
	if actualBalance != expectedBalance {
		t.Errorf("TestWithdraw-expected=%d not equal to actual=%d", expectedBalance, actualBalance)
	}
}
