package player

import "testing"

func TestNewPlayer(t *testing.T) {
	actual := NewPlayers("C", "K")
	if actual == nil {
		t.Errorf("New Player returned nil !")
	}
}
