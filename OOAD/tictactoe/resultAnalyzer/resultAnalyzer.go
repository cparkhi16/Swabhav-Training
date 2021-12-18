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
func (ra *Result) CheckWinning(currRow,currCol int) (ResultAnalysis, string) {
	
	i := 0
	test := false
	//horizantel test
	/*for i < ra.Board.Size*ra.Board.Size {
		
		if ra.Board.GameBoard.Cells[i]!=""{
		test = ra.CheckRow(i)
		}
		
		if !test {
			i += ra.Board.Size
		} else {
			ra.resutlt = Win
			return ra.resutlt, ra.Board.GameBoard.Cells[i]
		}
		
	}*/
	if ra.Board.GameBoard[currRow*ra.Board.Size].GetMark()!=b.Empty{
	test = ra.CheckRow(currRow*ra.Board.Size)
	}
	if test{
		ra.resutlt = Win
		return ra.resutlt,string( ra.Board.GameBoard[i].GetMark()) 
	}
	i = 0
	//vertical test
	testCol:=false
	/*for i < (ra.Board.Size) {
		
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
*/
if ra.Board.GameBoard[currCol].GetMark()!=b.Empty{
	testCol = ra.CheckCol(currCol)
	}
	if testCol{
		ra.resutlt = Win
		return ra.resutlt,string( ra.Board.GameBoard[i].GetMark()) 
	}
	if(currRow==currCol){
	//diagonal 1 test
	testDiagOne:=false
	if  ra.Board.GameBoard[0].GetMark() != b.Empty {
		testDiagOne=ra.CheckFirstDiagonal()
		if testDiagOne{
		ra.resutlt=Win
		return ra.resutlt,string( ra.Board.GameBoard[i].GetMark())
		}
	}
}	
   if(currRow+currCol==ra.Board.Size-1){
	//diagonal 2 test
	testDiagTwo:=false
	if  ra.Board.GameBoard[ra.Board.Size-1].GetMark() != b.Empty {
		testDiagTwo=ra.CheckSecondDiagonal()
		if testDiagTwo{
		ra.resutlt=Win
		return ra.resutlt,string( ra.Board.GameBoard[i].GetMark())
		}
	}
}
	isBoardFull := true
	for i := 0; i < ra.Board.Size*ra.Board.Size; i++ {
		//fmt.Println("Size ",ra.Board.Size)
			if ra.Board.GameBoard[i].GetMark() == b.Empty {
				isBoardFull = false
			}
		
	}
	if  isBoardFull==true {
		ra.resutlt = Draw
		return ra.resutlt, string(b.Empty)
	}
	ra.resutlt = GameOn
	return ra.resutlt, string(b.Empty)
}
func (ra *Result)CheckRow(i int)bool{
	val:=ra.Board.GameBoard[i].GetMark()
	for j:=0;j<ra.Board.Size;j++{
		//fmt.Println("Size ===",j)
		//fmt.Println("val ",val)
	//fmt.Println("val of roe c ", ra.Board.GameBoard[i+j].GetMark())
		//c:=ra.Board.GameBoard.Cells[i+j]
	 if val!=ra.Board.GameBoard[i+j].GetMark() { 
				return false
			}
			continue 				
		}	
	
	//fmt.Println("Outside for ")
	return true
}

func (ra *Result)CheckCol(i int )bool{
	val:=ra.Board.GameBoard[i].GetMark()
	index:=0
	for j:=0;j<ra.Board.Size;j++{
		
		//fmt.Println("Index == ",index)
		//fmt.Println("vsl here ",val)
		//fmt.Println("ra val--- ",ra.Board.GameBoard.Cells[i+index])
		if val!=ra.Board.GameBoard[i+index].GetMark() {
			return false
		}
		index=index+ra.Board.Size
		continue
		
	}
	
	return true
}

func (ra *Result)CheckFirstDiagonal()bool{
	val:=ra.Board.GameBoard[0].GetMark()
	index:=0
	for i:=0;i<ra.Board.Size;i++{
		if ra.Board.GameBoard[index].GetMark()!=val{
			
			//if index<ra.Board.Size{
			return false
			//}
		}
		index=index+ra.Board.Size+1
		continue

	}
return true
}


func (ra *Result)CheckSecondDiagonal()bool{
	val:=ra.Board.GameBoard[ra.Board.Size-1].GetMark()
	index:=(ra.Board.Size-1)*2
	for i:=0;i<ra.Board.Size-1;i++{
		if ra.Board.GameBoard[index].GetMark()!=val{
			return false
		}
		index=index+(ra.Board.Size-1)
		continue
	}
return true
}