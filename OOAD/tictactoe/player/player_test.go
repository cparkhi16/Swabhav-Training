package player

import ("testing")

func TestNewPlayer(t *testing.T) {
	actual := NewPlayers()
	if actual == nil {
		t.Errorf("New Player returned nil !")
	}
}

func TestSetPlayerDetails(t *testing.T) {
	p := NewPlayers()
	//p.Players=make(map[string]b.Mark)
	p.SetPlayerDetails("Chinmay")
	if p.Name != "Chinmay" {
		t.Errorf("Wrong player details")
	}
}

