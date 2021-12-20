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
		0, "minimum size should be three",
	}, {
		-5, "minimum size should be three",
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
	game.currentPlayer.Mark = b.X
	game.takeTurns()
	actual := game.currentPlayer.Mark
	expected := b.O
	if actual != expected {
		t.Errorf("Error while taking player turns ")
	}
}
