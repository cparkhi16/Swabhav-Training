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

func TestMakeMove(t *testing.T) {
	game := New()
	game.Initialize(3, "Chinmay", "Keyur")
	var list = []struct {
		mov      string
		p        b.Mark
		expected bool
	}{{
		"121", b.X, false,
	}, {
		"11", b.O, true,
	}, {
		"11", b.X, false,
	}, {
		"5", b.X, false,
	}}
	for _, val := range list {
		actual := game.MakeMove(val.mov, val.p)
		if actual != val.expected {
			t.Errorf("Error found for makeMove function ")
		}
	}
}

func TestInitialize(t *testing.T) {
	var list = []struct {
		size     int
		n1       string
		n2       string
		expected string
	}{{
		0, "Chinmay", "keyur", "size should not be less than or equal to one",
	}, {
		-5, "Chinmay", "keyur", "size should not be less than or equal to one",
	},
	}
	for _, val := range list {
		game := New()
		actual := game.Initialize(val.size, val.n1, val.n2)
		if actual.Error() != val.expected {
			t.Errorf("Error found for initialize function ")
		}
	}
}

func TestTakeTurns(t *testing.T) {
	game := New()
	game.Initialize(3, "Chinmay", "Keyur")
	game.Board.GameBoard.CurrentMark = b.X
	game.takeTurns()
	actual := game.Board.GameBoard.CurrentMark
	expected := b.O
	if actual != expected {
		t.Errorf("Error while taking player turns ")
	}
}
