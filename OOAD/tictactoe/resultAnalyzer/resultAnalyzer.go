package resultAnalyzer

import (b "cmdgame/board")


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
	
	i := 0
	test := false
	//horizantel test
	for i < ra.Board.Size*ra.Board.Size {
		
		if ra.Board.GameBoard.Cells[i]!=""{
		test = ra.CheckRow(i)
		}
		
		if !test {
			i += ra.Board.Size
		} else {
			ra.resutlt = Win
			return ra.resutlt, ra.Board.GameBoard.Cells[i]
		}
		
	}
	i = 0
	//vertical test
	testCol:=false
	for i < (ra.Board.Size) {
		
		if ra.Board.GameBoard.Cells[i]!=""{
		testCol = ra.CheckCol(i)
		}
		//fmt.Println("test val",testCol)
		if !testCol {
			i += 1
		} else {
			ra.resutlt = Win
			return ra.resutlt, ra.Board.GameBoard.Cells[i]
		}
	}

	//diagonal 1 test
	testDiagOne:=false
	if  ra.Board.GameBoard.Cells[0] != "" {
		testDiagOne=ra.CheckFirstDiagonal()
		if testDiagOne{
		ra.resutlt=Win
		return ra.resutlt, ra.Board.GameBoard.Cells[i]
		}
	}
	//diagonal 2 test
	testDiagTwo:=false
	if  ra.Board.GameBoard.Cells[ra.Board.Size-1] != "" {
		testDiagTwo=ra.CheckSecondDiagonal()
		if testDiagTwo{
		ra.resutlt=Win
		return ra.resutlt, ra.Board.GameBoard.Cells[i]
		}
	}
	isBoardFull := true
	for i := 0; i < ra.Board.Size*ra.Board.Size; i++ {
		//fmt.Println("Size ",ra.Board.Size)
		
			if ra.Board.GameBoard.Cells[i] == "" {
				isBoardFull = false
			}
		
	}
	if  isBoardFull==true {
		ra.resutlt = Draw
		return ra.resutlt, ""
	}
	ra.resutlt = GameOn
	return ra.resutlt, ""
}
func (ra *Result)CheckRow(i int)bool{
	val:=ra.Board.GameBoard.Cells[i]
	//check:=make(map[string]int)
	
	for j:=0;j<ra.Board.Size;j++{
		//fmt.Println("Size ===",j)
		//fmt.Println("val ",ra.Board.GameBoard.Cells[i+j])
	//fmt.Println("val of roe c ", ra.Board.GameBoard.Cells[i+j])
		//c:=ra.Board.GameBoard.Cells[i+j]
	 if val==ra.Board.GameBoard.Cells[i+j]{ 
				continue
			} else{
				return false
			}
			
		}	
	
	//fmt.Println("Outside for ")
	return true
}

func (ra *Result)CheckCol(i int )bool{
	val:=ra.Board.GameBoard.Cells[i]
	index:=0
	for j:=0;j<ra.Board.Size;j++{
		
		//fmt.Println("Index == ",index)
		//fmt.Println("vsl here ",val)
		//fmt.Println("ra val--- ",ra.Board.GameBoard.Cells[i+index])
		if val==ra.Board.GameBoard.Cells[i+index] {
			index=index+ra.Board.Size
			continue
		}else{
			return false
		}
		
	}
	
	return true
}

func (ra *Result)CheckFirstDiagonal()bool{
	val:=ra.Board.GameBoard.Cells[0]
	index:=0
	for i:=0;i<ra.Board.Size;i++{
		if ra.Board.GameBoard.Cells[index]==val{
			
			//if index<ra.Board.Size{
			index=index+ra.Board.Size+1
			continue
			//}
		}else{
			return false
		}

	}
return true
}


func (ra *Result)CheckSecondDiagonal()bool{
	val:=ra.Board.GameBoard.Cells[ra.Board.Size-1]
	index:=(ra.Board.Size-1)*2
	for i:=0;i<ra.Board.Size-1;i++{
		if ra.Board.GameBoard.Cells[index]==val{
			index=index+(ra.Board.Size-1)
			continue
		}else{
			return false
		}
	}
return true
}