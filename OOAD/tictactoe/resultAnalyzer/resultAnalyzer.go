package resultAnalyzer

import b "cmdgame/board"

type Result struct {
	Board   *b.Board
	resutlt ResultAnalysis
}
type ResultAnalysis int

const (
	Draw   ResultAnalysis = 0
	Win    ResultAnalysis = 1
	GameOn ResultAnalysis = 3
)

func NewAnalyzer(Board *b.Board) *Result {
	return &Result{Board: Board}
}
func (ra *Result) CheckWinning() ResultAnalysis {
	// check horizontal
	for _, r := range ra.Board.GameBoard.Cells {
		check := make(map[string]int)

		for _, c := range r {
			if c == "" {
				continue
			} else {
				if check[c] == 0 {
					check[c] = 1
				} else {
					check[c] += 1
				}

				if check[c] == ra.Board.Size {
					//return true, false
					ra.resutlt = Win
					return ra.resutlt
				}
			}
		}
	}

	// check vertical
	for c := 0; c < ra.Board.Size; c++ {
		check := make(map[string]int)

		for r := 0; r < ra.Board.Size; r++ {
			if ra.Board.GameBoard.Cells[r][c] == "" {
				continue
			} else {
				if check[ra.Board.GameBoard.Cells[r][c]] == 0 {
					check[ra.Board.GameBoard.Cells[r][c]] = 1
				} else {
					check[ra.Board.GameBoard.Cells[r][c]] += 1
				}

				if check[ra.Board.GameBoard.Cells[r][c]] == ra.Board.Size {
					//return true, false
					ra.resutlt = Win
					return ra.resutlt
				}
			}
		}
	}

	check := make(map[string]int)
	for i := 0; i < ra.Board.Size; i++ {
		if ra.Board.GameBoard.Cells[i][i] == "" {
			continue
		} else {
			if check[ra.Board.GameBoard.Cells[i][i]] == 0 {
				check[ra.Board.GameBoard.Cells[i][i]] = 1
			} else {
				check[ra.Board.GameBoard.Cells[i][i]] += 1
			}

			if check[ra.Board.GameBoard.Cells[i][i]] == ra.Board.Size {
				//return true, false
				ra.resutlt = Win
				return ra.resutlt
			}
		}
	}
	check = make(map[string]int)
	decr := ra.Board.Size - 1
	for i := 0; i < ra.Board.Size; i++ {
		if ra.Board.GameBoard.Cells[i][decr] == "" {
			continue
		} else {
			if check[ra.Board.GameBoard.Cells[i][decr]] == 0 {
				check[ra.Board.GameBoard.Cells[i][decr]] = 1
			} else {
				check[ra.Board.GameBoard.Cells[i][decr]] += 1
			}

			if check[ra.Board.GameBoard.Cells[i][decr]] == ra.Board.Size {
				//return true, false
				ra.resutlt = Win
				return ra.resutlt
			}
		}

		decr -= 1
	}
	isBoardFull := true
	for i := 0; i < ra.Board.Size; i++ {
		for j := 0; j < ra.Board.Size; j++ {
			if ra.Board.GameBoard.Cells[i][j] == "" {
				isBoardFull = false
			}
		}
	}
	if isBoardFull {
		//return false, true
		ra.resutlt = Draw
		return ra.resutlt
	}
	//return false, false
	ra.resutlt = GameOn
	return ra.resutlt
}
