package account

import (
	"fmt"
	"testing"
)

var testAcc = New(123, 1, 1000)

func TestNew(t *testing.T) {
	expectedAccNo := 123
	actualAccNo := testAcc.GetAccountNo()
	expectedAccHolderId := 1
	actualAccHolderId := testAcc.GetAccountHolderId()
	expectedAccBalance := 1000
	actualAccBalance := testAcc.GetBalance()
	if expectedAccNo != actualAccNo || expectedAccHolderId != actualAccHolderId || expectedAccBalance != actualAccBalance {
		t.Errorf("expectedAccNo %d, actualAccNo %d", expectedAccNo, actualAccNo)
		t.Errorf("expectedAccHolderId %d, actualAccHolderId %d", expectedAccHolderId, actualAccHolderId)

	}
}

func TestGetBalance(t *testing.T) {
	expected := 1000
	actual := testAcc.GetBalance()
	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	} else {
		fmt.Println("passed GetBalance")
	}
}

func TestSetBalance(t *testing.T) {
	expected := 2000
	testAcc.SetBalance(2000)
	actual := testAcc.GetBalance()
	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	} else {
		fmt.Println("passed GetBalance")
	}
}
