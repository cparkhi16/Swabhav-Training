package squareRoot

import (
	"testing"
)

func TestSquareRoot(t *testing.T) {
	actual := SquareRoot(4)
	var expected float64 = 2
	if expected != actual {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}
