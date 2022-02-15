package power

import (
	"testing"
)

func TestPower(t *testing.T) {
	actual := Power(4, 2)
	var expected float64 = 16
	if expected != actual {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}
