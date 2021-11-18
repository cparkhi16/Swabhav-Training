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
