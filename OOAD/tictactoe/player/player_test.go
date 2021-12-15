package player

import ("testing"
b "cmdgame/board")

func TestNewPlayer(t *testing.T) {
	actual := NewPlayers()
	if actual == nil {
		t.Errorf("New Player returned nil !")
	}
}

func TestSetPlayerDetails(t *testing.T) {
	p := NewPlayers()
	p.Players=make(map[string]b.Mark)
	p.SetPlayerDetails("Chinmay")
	if p.Name != "Chinmay" {
		t.Errorf("Wrong player details")
	}
}

func TestGetPlayerFromMap(t *testing.T){
	p := NewPlayers()
	m:=make(map[string]b.Mark)
	m["Chinmay"]=b.X
	expected:="Chinmay"
	expectedStatus:=true
	actual,actualStatus:=p.GetPlayerFromMap(m,b.X)
	if actual!=expected || expectedStatus!=actualStatus{
		t.Errorf("Wrong player got from map")
	}
	
}