package multiply

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	actual := Multiply(4, 6)
	expected := 24
	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
