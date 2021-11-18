package calculator

import "testing"

func TestAddBulk(t *testing.T) {
	var l = []struct {
		n1, n2   int
		expected int
	}{{
		10, 20, 30,
	}, {
		11, 10, 21,
	},
	}
	for _, val := range l {
		actual := Add(val.n1, val.n2)
		if val.expected != actual {
			t.Errorf("Error found for addition actual %v and expected %v", actual, val.expected)
		}
	}
}

func TestSubtractBulk(t *testing.T) {
	var l = []struct {
		n1, n2   int
		expected int
	}{{
		10, 20, -10,
	}, {
		11, 10, 1,
	},
	}
	for _, val := range l {
		actual := Subtract(val.n1, val.n2)
		if val.expected != actual {
			t.Errorf("Error found for subtraction actual %v and expected %v", actual, val.expected)
		}
	}
}
func TestSquareRootBulk(t *testing.T) {
	var l = []struct {
		n1       float64
		expected float64
	}{{
		4, 2,
	}, {
		9, 3,
	},
	}
	for _, val := range l {
		actual := SquareRoot(val.n1)
		if val.expected != actual {
			t.Errorf("Error found for squareroot actual %v and expected %v", actual, val.expected)
		}
	}
}

func TestPower(t *testing.T) {
	var l = []struct {
		n1       int
		n2       int
		expected int
	}{{
		2, 1, 2,
	}, {
		2, 3, 8,
	},
	}
	for _, val := range l {
		actual := ipow(val.n1, val.n2)
		if val.expected != actual {
			t.Errorf("Error found for power actual %v and expected %v", actual, val.expected)
		}
	}
}
