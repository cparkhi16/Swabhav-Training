package board

import (
	"fmt"
	"strconv"
)

type Board struct {
	Size      int
	GameBoard []*Cell
}

func GetDesiredBoardSize() int {
	fmt.Println("Enter desired size of board")
	var size int
	fmt.Scanln(&size)
	return size
}
func MakeNewBoard(size int) *Board {
	GameBoard := make([]*Cell, size*size)
	//c.Cells = GameBoard
	for i:=0;i<size*size;i++{
		GameBoard[i]=NewCell(Empty)
	}
	return &Board{GameBoard: GameBoard, Size: size}
}
func (b *Board) ShowBoard() error {
	if b == nil {
		return fmt.Errorf("null Board receiver")
	}
	c:=1
	for i:=0;i< b.Size*b.Size ;i++{
		c++
		if b.GameBoard[i].GetMark() == Empty {
			fmt.Printf("-")
		} else {
			fmt.Printf( string(b.GameBoard[i].GetMark()))
		}
		/*fmt.Printf( string(b.GameBoard[i].GetMark()))
		fmt.Printf("|")*/
		if c>b.Size{
			//fmt.Printf("|")
			fmt.Println("")
			c=1
		}else{
			fmt.Printf("|")
		}
	}
	
	return nil
}

func (b *Board) CheckIsValidMove(mov string) (bool, int) {
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

	if b.GameBoard[r].GetMark() != Empty{
		fmt.Println(mov, "is already occupied on the board !! ")
		return false, -1
	}
	return true, r
}
func (b *Board) MakeMove(mov string, mark Mark) bool {
	isValid, pos := b.CheckIsValidMove(mov)
	if isValid {
		if b.GameBoard[pos].GetMark() == Empty{
			b.GameBoard[pos].SetMark(mark)
			return true
		}
	}
	return false
}
