package resultAnalyzer

import (
	"testing"
	"tictactoe/board"
	"tictactoe/mark"
	"tictactoe/result"
)

var board1 *board.Board = board.NewBoard(2)

var r = NewResultAnalyzer(board1)

func TestCheckBoard(t *testing.T) {
	board1.MakeEntryToBoard(0, 0, mark.Cross)
	board1.MakeEntryToBoard(1, 1, mark.Cross)
	actual := result.Winner
	expected := r.CheckBoard()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}
