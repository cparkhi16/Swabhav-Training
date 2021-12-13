package board

import (
	"fmt"
)

type Board struct {
	Size      int
	GameBoard *Cell
}

func GetDesiredBoardSize() int {
	fmt.Println("Enter desired size of board")
	var size int
	fmt.Scanln(&size)
	return size
}
func MakeNewBoard(size int, c *Cell) *Board {
	GameBoard := make([]string, size)
	/*for r, _ := range GameBoard {
		GameBoard[r] = make([]string, size)
	}*/
	c.Cells = GameBoard
	return &Board{GameBoard: c, Size: size}
}
func (b *Board) ShowBoard() error {
	if b == nil {
		return fmt.Errorf("null Board receiver")
	}
	for i, v := range b.GameBoard.Cells {
		if v == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(v)
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")

		}

	}
	return nil
}
