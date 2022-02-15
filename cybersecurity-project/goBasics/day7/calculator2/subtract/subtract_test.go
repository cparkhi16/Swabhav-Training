package subtract

import (
	"fmt"
	"testing"
)

func TestSubtract(t *testing.T) {

	var list = []struct {
		arg1     int
		arg2     int
		expected int
	}{
		{20, 10, 10},
		{0, 2, -2},
	}

	for i, v := range list {
		actual := Subtract(v.arg1, v.arg2)
		expected := v.expected
		if expected != actual {
			t.Errorf("expected %d, actual %d", v.expected, actual)
		} else {
			fmt.Println("passed test-", i)
		}
	}

}
