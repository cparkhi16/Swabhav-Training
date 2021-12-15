package tictactoegame

import (
	b "cmdgame/board"
	"testing"
	
)

func TestNewTicTacToe(t *testing.T) {
	actual := New()
	if actual == nil {
		t.Errorf("New Tictactoe returned nil !")
	}
}


func TestInitialize(t *testing.T) {
	var list = []struct {
		size     int
		expected string
	}{{
		0, "size should not be less than or equal to one",
	}, {
		-5, "size should not be less than or equal to one",
	},
	}
	for _, val := range list {
		game := New()
		actual := game.Initialize(val.size)
		if actual.Error() != val.expected {
			t.Errorf("Error found for initialize function ")
		}
	}
}

func TestTakeTurns(t *testing.T) {
	game := New()
	game.Initialize(3)
	game.Board.GameBoard.CurrentMark = b.X
	game.takeTurns()
	actual := game.Board.GameBoard.CurrentMark
	expected := b.O
	if actual != expected {
		t.Errorf("Error while taking player turns ")
	}
}
