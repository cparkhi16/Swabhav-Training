package accountHolder

import (
	"testing"
)

func TestNewAccountHolder(t *testing.T) {
	var list = []struct {
		fName          string
		lname          string
		contact        string
		age            uint8
		accountNumber  uint8
		id             int
		balance        int
		gender         Gender
		expected       string
		expectedPerson *person
	}{{
		"Chinmay", "Parkhi", "7876675432", 21, 100, 400, 1000, Male, "", nil,
	}, {
		"Rajesh", "Patil", "9879901234", 22, 123, 500, 1200, Male, "", nil,
	}}
	for _, val := range list {
		actual, err := NewAccountHolder(val.fName, val.lname, val.contact, val.age, val.accountNumber, val.id, val.balance, val.gender)
		if val.expected != err.Error() && val.expectedPerson != actual {
			t.Errorf("Error found for NewAccountPresent Error actual %v and expected %v", err.Error(), val.expected)
			t.Errorf("Error found for NewAccountPresent Pointer actual %v and expected %v", actual, val.expectedPerson)
		}
	}
}

func TestIsAccountPresent(t *testing.T) {
	accTest, _ := NewAccountHolder("Chinmay", "Parkhi", "7876675432", 21, 100, 400, 1000, Male)
	var list = []struct {
		accountNumber uint8
		expected      string
	}{{
		100, "",
	}, {
		19, "Account does not exist",
	},
	}
	for _, val := range list {
		actual := accTest.isAccountPresent(val.accountNumber)
		if val.expected != actual.Error() {
			t.Errorf("Error found for isAccountPrsesnt actual %v and expected %v", actual, val.expected)
		}
	}
}

func TestShareMoney(t *testing.T) {
	userOne, _ := NewAccountHolder("Chinmay", "Parkhi", "7876675432", 21, 231, 400, 1000, Male)
	userTwo, _ := NewAccountHolder("Rajesh", "Patil", "9879901234", 22, 123, 500, 1200, Male)
	//userOne.GetAccountDetails()
	//userTwo.GetAccountDetails()
	var list = []struct {
		accTwo           *person
		accountNumberOne uint8
		accountNumberTwo uint8
		money            int
		expected         string
		err              string
	}{{
		userTwo, 231, 123, 500, "Transfer successful", "",
	}, {
		userTwo, 231, 123, -100, "Transfer unsuccessful", "negative money can't be transfered",
	},
	}
	for _, val := range list {
		issuccess, err := userOne.ShareMoney(val.accTwo, val.accountNumberOne, val.accountNumberTwo, val.money)
		if val.err != err.Error() && val.expected != issuccess {
			t.Errorf("Error found for Sharing Money Error actual %v and expected %v", err.Error(), val.err)
			t.Errorf("Error found for Succsess message actual %v and expected %v", issuccess, val.expected)
		}
		userOne, _ = NewAccountHolder("Chinmay", "Parkhi", "7876675432", 21, 231, 400, 1000, Male)
		val.accTwo, _ = NewAccountHolder("Rajesh", "Patil", "9879901234", 22, 123, 500, 1200, Male)
	}
}

func TestUpdateAccount(t *testing.T) {
	userOne, _ := NewAccountHolder("Chinmay", "Parkhi", "7876675432", 21, 231, 400, 1000, Male)
	//userTwo, _ := NewAccountHolder("Rajesh", "Patil", "9879901234", 22, 123, 500, 1200, Male)
	var list = []struct {
		accountNumber uint8
		balance       int
		expected      string
	}{{
		231, 1000, "",
	}, {
		123, -800, "Negative balance not accepted",
	},
	}
	for _, val := range list {
		err := userOne.UpdateAccount(val.accountNumber, val.balance)
		if val.expected != err.Error() {
			t.Errorf("Error found for UpdateAccount actual %v and expected %v", err.Error(), val.expected)
		}

	}
}

func TestGetBalance(t *testing.T) {
	userOne, _ := NewAccountHolder("Chinmay", "Parkhi", "7876675432", 21, 231, 400, 1000, Male)
	//userTwo, _ := NewAccountHolder("Rajesh", "Patil", "9879901234", 22, 123, 500, 1200, Male)
	var list = []struct {
		accountNumber uint8
		expected      int
	}{{
		231, 1000,
	}, {
		123, 1200,
	}}
	for _, val := range list {
		balance := userOne.GetBalance(val.accountNumber)
		if val.expected != balance {
			t.Errorf("Error found for GetBalance Error actual %v and expected %v", balance, val.expected)
		}
		userOne, _ = NewAccountHolder("Rajesh", "Patil", "9879901234", 22, 123, 500, 1200, Male)
	}

}
