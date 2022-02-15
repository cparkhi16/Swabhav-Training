package calculator

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	actual := Add(4, 6)
	expected := 10
	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

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

func TestMultiply(t *testing.T) {
	actual := Multiply(4, 6)
	expected := 24
	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestDivide(t *testing.T) {
	actual := Divide(6, 2)
	expected := 3
	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
