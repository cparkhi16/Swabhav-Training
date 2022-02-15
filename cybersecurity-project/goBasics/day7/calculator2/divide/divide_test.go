package divide

import (
	"testing"
)

func TestDivide(t *testing.T) {
	actual := Divide(6.4, 2.5)
	var expected float32 = 2.56
	if expected != actual {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}
