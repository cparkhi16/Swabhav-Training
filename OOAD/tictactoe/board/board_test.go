package board

import "testing"

func TestNewCell(t *testing.T) {
	actual := NewCell(Empty)
	if actual == nil {
		t.Errorf("New Cell returned nil instance !")
	}
}
func TestMakeNewBoard(t *testing.T){
	//cell := NewCell(Empty)
	ActualBoard := MakeNewBoard(3)
	if ActualBoard==nil{
		t.Errorf("New Board returned nil instance !")
	}

}
func TestShowBoard(t *testing.T) {
	//cell := NewCell()
	Board := MakeNewBoard(3)
	expected := error(nil)
	actual := Board.ShowBoard()
	if actual != expected {
		t.Errorf("Error found for show board")
	}
}

func TestMakeMove(t *testing.T) {
	//game := New()
	//game.Initialize(3)
	//cell := NewCell()
	Board := MakeNewBoard(3)
	var list = []struct {
		mov      string
		p        Mark
		expected bool
	}{{
		"121", X, false,
	}, {
		"1", O, true,
	}, {
		"1", X, false,
	}, {
		"500", X, false,
	}}
	for _, val := range list {
		actual,_ := Board.MakeMove(val.mov, val.p)
		if actual != val.expected {
			t.Errorf("Error found for makeMove function ")
		}
	}
}
