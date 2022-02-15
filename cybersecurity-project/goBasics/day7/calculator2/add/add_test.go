package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	actual := add(2.0, 6.0)
	var expected float32 = 10.0
	if expected != actual {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}
