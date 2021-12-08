package board

import "testing"

func TestNewCell(t *testing.T) {
	actual := NewCell()
	if actual == nil {
		t.Errorf("New Cell returned nil instance !")
	}
}

func TestShowBoard(t *testing.T) {
	cell := NewCell()
	Board := MakeNewBoard(3, cell)
	expected := error(nil)
	actual := Board.ShowBoard()
	if actual != expected {
		t.Errorf("Error found for show board")
	}
}
