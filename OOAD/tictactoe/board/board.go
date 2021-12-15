package board

import (
	"fmt"
	"strconv"
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
	GameBoard := make([]string, size*size)
	/*for r, _ := range GameBoard {
		GameBoard[r] = make([]string, size)
	}*/
	//fmt.Println("V",GameBoard)
	c.Cells = GameBoard
	return &Board{GameBoard: c, Size: size}
}
func (b *Board) ShowBoard() error {
	if b == nil {
		return fmt.Errorf("null Board receiver")
	}
	for i, v := range b.GameBoard.Cells {
		if v == string(Empty) {
			fmt.Printf("-")
		} else {
			fmt.Printf(v)
		}
		
		if i > 0 && (i+1)%b.Size==0 {
			//fmt.Println("Val", (i+1)%((b.Size/2)-1))
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")

		}

	}
	
	return nil
}

func (b *Board) CheckIsValidMove(mov string) (bool, int) {
	//split := strings.Split(mov, "")
	/*if len(split) > 2 || len(split) == 1 {
		fmt.Println("Invalid move :", mov)
		return false, -1, -1
	}*/
	/*r, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println("Invalid move :", mov)
		return false, -1, -1
	}
	c, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Println("Invalid move:", mov)
		return false, -1, -1
	}*/
	r, err := strconv.Atoi(mov)
	if err != nil {
		fmt.Println("Invalid move :", mov)
		return false, -1
	}
	r = r - 1

	switch {
	case r < 0, r >= b.Size*b.Size:
		fmt.Println("Invalid move:", mov)
		return false, r
	}

	if b.GameBoard.Cells[r] != string(Empty) {
		fmt.Println(mov, "is already occupied on the board !! ")
		return false, -1
	}
	return true, r
}
func (b *Board) MakeMove(mov string, mark Mark) bool {
	isValid, pos := b.CheckIsValidMove(mov)
	if isValid {
		if b.GameBoard.Cells[pos] == string(Empty) {
			b.GameBoard.Cells[pos] = string(mark)
			return true
		}
	}
	return false
}
