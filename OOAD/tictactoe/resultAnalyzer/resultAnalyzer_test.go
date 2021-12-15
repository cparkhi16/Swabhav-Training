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
func TestCheckRow(t *testing.T){
	c := b.NewCell()
	board := b.MakeNewBoard(3, c)
	var list = []struct {
		i int
		board    []string
		expected bool
	}{{
		0,[]string{"x", "x", "x" ,"", "o", "","o", "", ""}, true,
	}, {2, []string{"x", "", "x","", "o", "","o", "", ""}, false,
	}}
	for _, val := range list {
		r := NewAnalyzer(board)
		board.GameBoard.Cells = val.board
		actual :=r.CheckRow(val.i)
		if actual != val.expected {
			t.Errorf("Error found in actual : %v expected : %v check for check row", actual, val.expected)
		}
	}
}
func TestCheckCol(t *testing.T){
	c := b.NewCell()
	board := b.MakeNewBoard(3, c)
	var list = []struct {
		i int
		board    []string
		expected bool
	}{{
		0,[]string{"x", "o", "x" ,"x", "o", "","x", "", ""}, true,
	}, {2, []string{"x", "", "x","", "o", "","o", "", ""}, false,
	}}
	for _, val := range list {
		r := NewAnalyzer(board)
		board.GameBoard.Cells = val.board
		actual :=r.CheckCol(val.i)
		if actual != val.expected {
			t.Errorf("Error found in actual : %v expected : %v check for check col", actual, val.expected)
		}
	}
}
func TestCheckFirstDiagonal(t *testing.T){
	c := b.NewCell()
	board := b.MakeNewBoard(3, c)
	var list = []struct {
		board    []string
		expected bool
	}{{
		[]string{"x", "o", "x" ,"o", "x", "","", "o", "x"}, true,
	}, {[]string{"x", "", "x","", "o", "","o", "", ""}, false,
	}}
	for _, val := range list {
		r := NewAnalyzer(board)
		board.GameBoard.Cells = val.board
		actual :=r.CheckFirstDiagonal()
		if actual != val.expected {
			t.Errorf("Error found in actual : %v expected : %v check for check col", actual, val.expected)
		}
	}
}
func TestCheckSecondDiagonal(t *testing.T){
	c := b.NewCell()
	board := b.MakeNewBoard(3, c)
	var list = []struct {
		board    []string
		expected bool
	}{{
		[]string{"o", "o", "x" ,"o", "x", "","x", "o", "x"}, true,
	}, {[]string{"x", "", "x","", "o", "","o", "", ""}, false,
	}}
	for _, val := range list {
		r := NewAnalyzer(board)
		board.GameBoard.Cells = val.board
		actual :=r.CheckSecondDiagonal()
		if actual != val.expected {
			t.Errorf("Error found in actual : %v expected : %v check for check col", actual, val.expected)
		}
	}
}
func TestCheckWinning(t *testing.T) {
	c := b.NewCell()
	board := b.MakeNewBoard(3, c)
	var list = []struct {
		board    []string
		expected ResultAnalysis
	}{{
		[]string{"x", "x", "x" ,"", "o", "","o", "", ""}, Win,
	}, {
		[]string{"x", "", "x","", "o", "","o", "", ""}, GameOn,
	}, {
		[]string{"x", "o", "x","o", "x", "x","o", "x", "o"}, Draw,
	}, {
		[]string{"x", "o", "x","x", "x", "x","x", "x", "o"}, Win,
	}}
	for _, val := range list {
		r := NewAnalyzer(board)
		board.GameBoard.Cells = val.board
		actual,_ := r.CheckWinning()
		if actual != val.expected {
			t.Errorf("Error found in actual : %v expected : %v check", actual, val.expected)
		}
	}
}
