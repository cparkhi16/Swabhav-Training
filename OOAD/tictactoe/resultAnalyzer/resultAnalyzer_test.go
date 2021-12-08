package resultAnalyzer

import (
	b "cmdgame/board"
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	c := b.NewCell()
	board := b.MakeNewBoard(3, c)
	actual := NewAnalyzer(board)
	if actual == nil {
		t.Errorf("New Analyzer returned nil !")
	}
}

func TestCheckWinning(t *testing.T) {
	c := b.NewCell()
	board := b.MakeNewBoard(3, c)
	var list = []struct {
		board    [][]string
		expected ResultAnalysis
	}{{
		[][]string{{"x", "x", "x"}, {"", "o", ""}, {"o", "", ""}}, Win,
	}, {
		[][]string{{"x", "", "x"}, {"", "o", ""}, {"o", "", ""}}, GameOn,
	}, {
		[][]string{{"x", "o", "x"}, {"o", "x", "x"}, {"o", "x", "o"}}, Draw,
	}, {
		[][]string{{"x", "o", "x"}, {"x", "x", "x"}, {"x", "x", "o"}}, Win,
	}}
	for _, val := range list {
		r := NewAnalyzer(board)
		board.GameBoard.Cells = val.board
		actual := r.CheckWinning()
		if actual != val.expected {
			t.Errorf("Error found in actual : %v expected : %v check", actual, val.expected)
		}
	}
}
