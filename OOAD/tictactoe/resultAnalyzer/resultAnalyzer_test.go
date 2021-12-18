package resultAnalyzer

import (
	b "cmdgame/board"
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	//c := b.NewCell()
	board := b.MakeNewBoard(3)
	actual := NewAnalyzer(board)
	if actual == nil {
		t.Errorf("New Analyzer returned nil !")
	}
}
func TestCheckRow(t *testing.T){
	board := b.MakeNewBoard(3)
	r := NewAnalyzer(board)
	board.MakeMove("1",b.X)
	board.MakeMove("2",b.X)
	board.MakeMove("3",b.X)
	board.MakeMove("4",b.Empty)
	board.MakeMove("5",b.O)
	board.MakeMove("6",b.O)
	board.MakeMove("7",b.Empty)
	board.MakeMove("8",b.O)
	board.MakeMove("9",b.O)
	actual := r.CheckRow(0)
	expected:=true
	if actual!=expected{
		t.Errorf("Error found in actual : %v expected : %v check", actual, expected)
	}
}
func TestCheckCol(t *testing.T){
	board := b.MakeNewBoard(3)
	r := NewAnalyzer(board)
	board.MakeMove("1",b.O)
	board.MakeMove("2",b.X)
	board.MakeMove("3",b.X)
	board.MakeMove("4",b.O)
	board.MakeMove("5",b.X)
	board.MakeMove("6",b.O)
	board.MakeMove("7",b.O)
	board.MakeMove("8",b.Empty)
	board.MakeMove("9",b.O)
	actual:= r.CheckCol(0)
	expected:=true
	if actual!=expected{
		t.Errorf("Error found in actual : %v expected : %v check", actual, expected)
	}
}
func TestCheckFirstDiagonal(t *testing.T){
	board := b.MakeNewBoard(3)
	r := NewAnalyzer(board)
	board.MakeMove("1",b.X)
	board.MakeMove("2",b.O)
	board.MakeMove("3",b.X)
	board.MakeMove("4",b.O)
	board.MakeMove("5",b.X)
	board.MakeMove("6",b.O)
	board.MakeMove("7",b.Empty)
	board.MakeMove("8",b.O)
	board.MakeMove("9",b.X)
	actual:= r.CheckFirstDiagonal()
	expected:=true
	if actual!=expected{
		t.Errorf("Error found in actual : %v expected : %v check", actual, expected)
	}
}
func TestCheckSecondDiagonal(t *testing.T){
	board := b.MakeNewBoard(3)
	r := NewAnalyzer(board)
	board.MakeMove("1",b.X)
	board.MakeMove("2",b.O)
	board.MakeMove("3",b.X)
	board.MakeMove("4",b.O)
	board.MakeMove("5",b.X)
	board.MakeMove("6",b.O)
	board.MakeMove("7",b.X)
	board.MakeMove("8",b.O)
	board.MakeMove("9",b.Empty)
	actual:= r.CheckSecondDiagonal()
	expected:=true
	if actual!=expected{
		t.Errorf("Error found in actual : %v expected : %v check", actual, expected)
	}
}
func TestCheckWinning(t *testing.T) {
	//c := b.NewCell()
	board := b.MakeNewBoard(3)
	r := NewAnalyzer(board)
	board.MakeMove("1",b.X)
	board.MakeMove("2",b.X)
	board.MakeMove("3",b.X)
	board.MakeMove("4",b.Empty)
	board.MakeMove("5",b.O)
	board.MakeMove("6",b.O)
	board.MakeMove("7",b.Empty)
	board.MakeMove("8",b.O)
	board.MakeMove("9",b.O)
	actual,_ := r.CheckWinning(0,2)
	expected:=Win
	if actual!=expected{
		t.Errorf("Error found in actual : %v expected : %v check", actual, expected)
	}
	
}
