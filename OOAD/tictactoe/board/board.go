package board

import (
	"fmt"
)

type Board struct {
	Size      int
	GameBoard *Cell
}

func MakeNewBoard(size int, c *Cell) *Board {
	GameBoard := make([][]string, size)
	for r, _ := range GameBoard {
		GameBoard[r] = make([]string, size)
	}
	c.Cells = GameBoard
	return &Board{GameBoard: c, Size: size}
}
func (b *Board) ShowBoard() error {
	if b == nil {
		return fmt.Errorf("null Board receiver")
	}
	for r, _ := range b.GameBoard.Cells {
		for _, c := range b.GameBoard.Cells[r] {
			switch c {
			case "":
				fmt.Printf(" %v ", "-")
			default:
				fmt.Printf(" %v ", c)
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("")
	return nil
}
