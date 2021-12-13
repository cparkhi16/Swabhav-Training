package player

import "testing"

func TestNewPlayer(t *testing.T) {
	actual := NewPlayers()
	if actual == nil {
		t.Errorf("New Player returned nil !")
	}
}

func TestSetPlayerDetails(t *testing.T) {
	p := NewPlayers()
	p.SetPlayerDetails("Chinmay", "Keyur")
	if p.Player1 != "Chinmay" || p.Player2 != "Keyur" {
		t.Errorf("Wrong player details")
	}
}
