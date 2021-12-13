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
func (ra *Result) CheckWinning() (ResultAnalysis, string) {
	// check horizontal
	/*for _, r := range ra.Board.GameBoard.Cells {
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
	return ra.resutlt*/
	i := 0
	test := false
	//horizantel test
	for i < 9 {
		test = ra.Board.GameBoard.Cells[i] == ra.Board.GameBoard.Cells[i+1] && ra.Board.GameBoard.Cells[i+1] == ra.Board.GameBoard.Cells[i+2] &&
			ra.Board.GameBoard.Cells[i] != ""
		if !test {
			i += 3
		} else {
			ra.resutlt = Win
			return ra.resutlt, ra.Board.GameBoard.Cells[i]
		}
	}
	i = 0
	//vertical test
	for i < 3 {
		test = ra.Board.GameBoard.Cells[i] == ra.Board.GameBoard.Cells[i+3] && ra.Board.GameBoard.Cells[i+3] == ra.Board.GameBoard.Cells[i+6] && ra.Board.GameBoard.Cells[i] != ""
		if !test {
			i += 1
		} else {
			ra.resutlt = Win
			return ra.resutlt, ra.Board.GameBoard.Cells[i]
		}
	}

	//diagonal 1 test
	if ra.Board.GameBoard.Cells[0] == ra.Board.GameBoard.Cells[4] && ra.Board.GameBoard.Cells[4] == ra.Board.GameBoard.Cells[8] && ra.Board.GameBoard.Cells[0] != "" {
		return ra.resutlt, ra.Board.GameBoard.Cells[i]
	}
	//diagonal 2 test
	if ra.Board.GameBoard.Cells[2] == ra.Board.GameBoard.Cells[4] && ra.Board.GameBoard.Cells[4] == ra.Board.GameBoard.Cells[6] && ra.Board.GameBoard.Cells[2] != "" {
		return ra.resutlt, ra.Board.GameBoard.Cells[i]
	}
	if ra.Board.Size == 9 {
		ra.resutlt = Draw
		return ra.resutlt, ""
	}
	ra.resutlt = GameOn
	return ra.resutlt, ""
}
