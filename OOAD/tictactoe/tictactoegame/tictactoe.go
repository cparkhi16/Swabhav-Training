package tictactoegame

import (
	b "cmdgame/board"
	pl "cmdgame/player"
	r "cmdgame/resultAnalyzer"
	"fmt"
	"os"
	"strconv"
)

type TicTacToeGame struct {
	GameInitialized bool
	Player          *pl.Player
	Board           *b.Board
	Result          *r.Result
}

func New() *TicTacToeGame {
	return &TicTacToeGame{GameInitialized: false}
}
func (g *TicTacToeGame) Initialize(size int) error {
	if size <= 2 {
		return fmt.Errorf("minimum size should be three")
	} else {
		g.Player = pl.NewPlayers()
		g.Player.Players = make(map[string]b.Mark)
		n1,err := g.Player.GetPlayerDetails()
		if err!=nil{
			return err
		}
		g.Player.SetPlayerDetails(n1)
		n2,err:=g.Player.GetPlayerDetails()
		if err!=nil{
			return err
		}
		g.Player.SetPlayerDetails(n2)
		g.Player.Mark=b.Empty
		g.Board = b.MakeNewBoard(size)
		g.Player.Marks = [2]b.Mark{b.X, b.O}
		g.Result = r.NewAnalyzer(g.Board)
		g.GameInitialized = true
	}
	return nil
}

func (g *TicTacToeGame) ShowMenu() {
	fmt.Printf("Choose x/o for player2 %v \n", g.Player.Name)
	var input string
	fmt.Scanln(&input)
	if input != string(b.X) && input != string(b.O) {
		fmt.Println("Please provide valid input to start the game ")
		g.ShowMenu()
	} else {
		g.Player.Mark = b.Mark(input)
		g.Player.Players[g.Player.Name] = b.Mark(input)
	}
}

func (g *TicTacToeGame) GameOver() {
	fmt.Println("End Game!!")
	os.Exit(0)
}
func (g *TicTacToeGame) getMarkForCurrentPlayer() string {
	
	if g.Player.Mark == g.Player.Players[g.Player.Name] {
		p,_:=g.Player.GetPlayerFromMap(g.Player.Players,g.Player.Mark)
		fmt.Println("Current player : ",p)
	} else {
		k,_:=g.Player.GetPlayerFromMap(g.Player.Players,b.Empty)
		fmt.Println("Current player : ",k)
	}
	g.Board.ShowBoard()
	fmt.Print("Make a move or enter 'exit' to end the game : ")
	var input string
	fmt.Scanln(&input)
	return input
}
func (g *TicTacToeGame) inGame() bool {
	input := g.getMarkForCurrentPlayer()
	switch input {
	case "":
		return false
	case "exit":
		g.GameOver()
	default:
		successfulMove := g.Board.MakeMove(input, g.Player.Mark)
		if !successfulMove {
			return false
		}
		val,_:= strconv.Atoi(input)
		currRow:=(val-1)/g.Board.Size
		currCol:=(val-1)%g.Board.Size
		gameStatus, _ := g.Result.CheckWinning(currRow,currCol)
		if gameStatus == r.Win {
			g.Board.ShowBoard()
			if g.Player.Players[g.Player.Name] == g.Player.Mark {
				p,_:=g.Player.GetPlayerFromMap(g.Player.Players,g.Player.Mark)
				fmt.Println(p, "wins!")
			} else {
				k,_:=g.Player.GetPlayerFromMap(g.Player.Players,b.Empty)
				fmt.Println(k, "wins!")
			}
			g.GameOver()
		} else if gameStatus == r.Draw {
			g.Board.ShowBoard()
			fmt.Println("Game draw")
			g.GameOver()
		}

		g.takeTurns()
	}
	return true
}

func (g *TicTacToeGame) takeTurns() {
	if g.Player.Mark == b.X {
		g.Player.Mark = b.O
	} else {
		g.Player.Mark = b.X
	}
}

func (g *TicTacToeGame) Play() {
	size := b.GetDesiredBoardSize()
	for {
		switch g.GameInitialized {
		case false:
			e := g.Initialize(size)
			if e != nil {
				fmt.Println(e.Error())
				g.Play()
			} else {
				g.ShowMenu()
			}
		default:
			g.inGame()
		}
	}
}
