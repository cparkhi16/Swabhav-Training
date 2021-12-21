package board

import (
	"fmt"
	"strconv"
)

type Board struct {
	Size      int
	GameBoard []*Cell
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

func (b *Board) CheckIsValidMove(mov string) (bool, int,string) {
	r, err := strconv.Atoi(mov)
	if err != nil {
		s:=fmt.Sprintf("Invalid move : %v", mov)
		return false, -1,s
	}
	r = r - 1

	switch {
	case r < 0, r >= b.Size*b.Size:
		//fmt.Println("Invalid move:", mov)
		s:=fmt.Sprintf("Invalid move : %v", mov)
		return false, r,s
	}

	if b.GameBoard[r].GetMark() != Empty{
		//fmt.Println(mov, "is already occupied on the board !! ")
		s:=fmt.Sprintf(" %v is already occupied on the board !! ",mov)
		return false, -1,s
	}
	return true, r,"Valid move"
}
func (b *Board) MakeMove(mov string, mark Mark) (bool,string) {
	isValid, pos,s := b.CheckIsValidMove(mov)
	if isValid {
		if b.GameBoard[pos].GetMark() == Empty{
			b.GameBoard[pos].SetMark(mark)
			return true,s
		}
	}
	return false,s
}
