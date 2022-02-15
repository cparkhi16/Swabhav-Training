package resultAnalyzer

import (
	"tictactoe/board"
	"tictactoe/mark"
	"tictactoe/result"
)

type ResultAnalyzer struct {
	board *board.Board
}

func NewResultAnalyzer(board *board.Board) *ResultAnalyzer {
	return &ResultAnalyzer{
		board: board,
	}
}

func (r *ResultAnalyzer) CheckBoard(currentRow, currentCol int) result.Result {
	//fmt.Println("checkRows-", r.checkRows())
	//fmt.Println("checkColumns-", r.checkColumns())
	//fmt.Println("checkDiagonals-", r.checkDiagonals())
	if r.board.IsFull() {
		return result.Tie
	} else if r.checkRows(currentRow) || r.checkColumns(currentCol) || r.checkDiagonals(currentRow, currentCol) {
		return result.Winner
	}
	return result.InProgress

}

func (r *ResultAnalyzer) checkRows(currentRow int) bool {
	var rowValue mark.Mark = r.board.GetEntryAt(currentRow, 0)
	rowCount := 0
	for col := 0; col < r.board.GetSize(); col++ {
		if rowValue != mark.Empty && rowValue == r.board.GetEntryAt(currentRow, col) {
			rowCount++
		}
	}
	//fmt.Println(rowCount)
	if rowCount == r.board.GetSize() {
		return true
	}
	return false
}

func (r *ResultAnalyzer) checkColumns(currentCol int) bool {
	var columnValue mark.Mark = r.board.GetEntryAt(0, currentCol)
	columnCount := 0
	for row := 0; row < r.board.GetSize(); row++ {
		if columnValue != mark.Empty && columnValue == r.board.GetEntryAt(row, currentCol) {
			columnCount++
		}
	}
	//fmt.Println(columnCount)
	if columnCount == r.board.GetSize() {
		return true
	}
	return false
}

func (r *ResultAnalyzer) checkDiagonals(currentRow, currentCol int) bool {
	var isItDiagonal1Entry bool = false
	var isItDiagonal2Entry bool = false
	if currentCol == currentRow {
		isItDiagonal1Entry = true
	}
	if currentRow == r.board.GetSize()-currentCol-1 {
		isItDiagonal2Entry = true
	}

	diagonal1Count := 0
	diagonal2Count := 0

	if isItDiagonal1Entry {
		//left-up to right-bottom diagonal
		var diagonal1 = r.board.GetEntryAt(0, 0)
		for i := 1; i < r.board.GetSize(); i++ {
			if diagonal1 == r.board.GetEntryAt(i, i) && r.board.GetEntryAt(i, i) != mark.Empty {
				diagonal1Count++
			}
		}
	} else if isItDiagonal2Entry {
		//right-up to left-bottom diagonal
		var diagonal2 = r.board.GetEntryAt(0, r.board.GetSize()-1)
		for i := 1; i < r.board.GetSize(); i++ {
			if diagonal2 == r.board.GetEntryAt(i, r.board.GetSize()-i-1) && r.board.GetEntryAt(i, r.board.GetSize()-i-1) != mark.Empty {
				diagonal2Count++
			}
		}
	}
	//fmt.Println(diagonal1Count, diagonal2Count)
	if diagonal1Count == r.board.GetSize()-1 || diagonal2Count == r.board.GetSize()-1 {
		return true
	}
	return false

}
