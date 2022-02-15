package customer

import (
	"fmt"
	"testing"
)

var testCustomer1 = New(1, "test1", 123, 1000)
var testCustomer2 = New(2, "test2", 456, 2000)

func TestNew(t *testing.T) {
	testCustomer := New(1, "test1", 123, 1000)
	fmt.Println(testCustomer.id)
	expectedName := "test1"
	actualName := testCustomer.GetFirstName()
	expectedBalance := 100
	testCustomerAcc, _ := testCustomer.GetAccountFromNo(123)
	actualBalance := testCustomerAcc.GetBalance()
	if actualName != expectedName || actualBalance != expectedBalance {
		t.Errorf("expectedName %s, actualName %s", expectedName, actualName)
		t.Errorf("expectedBalance %d, actualBalance %d", expectedBalance, actualBalance)
	} else {
		fmt.Println("testing of New func passed")
	}

}

func TestTransferMoney(t *testing.T) {
	err := testCustomer1.TransferMoney(testCustomer2, 123, 456, 1000)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		expected1 := 0
		expected2 := 3000
		testCustomer1Acc, _ := testCustomer1.GetAccountFromNo(123)
		actual1 := testCustomer1Acc.GetBalance()
		testCustomer2Acc, _ := testCustomer2.GetAccountFromNo(456)
		actual2 := testCustomer2Acc.GetBalance()
		fmt.Println(actual1, actual2)
		if actual1 != expected1 || actual2 != expected2 {
			t.Errorf("expected1 %d, actual1 %d", expected1, actual1)
			t.Errorf("expected2 %d, actual2 %d", expected2, actual2)
		} else {
			fmt.Println("testing transfer money successful")
		}
	}

}
