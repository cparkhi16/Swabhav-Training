package tictactoegame

import (
	b "cmdgame/board"
	pl "cmdgame/player"
	r "cmdgame/resultAnalyzer"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	if size <= 1 {
		return fmt.Errorf("size should not be less than or equal to one")
	} else {
		g.Player = pl.NewPlayers()
		g.Player.Players = make(map[string]b.Mark)
		n1, n2 := g.Player.GetPlayerDetails()
		g.Player.SetPlayerDetails(n1, n2)
		cell := b.NewCell()
		g.Board = b.MakeNewBoard(size, cell)
		g.Board.GameBoard.CurrentMark = ""
		g.Board.GameBoard.Marks = [2]b.Mark{b.X, b.O}
		g.Result = r.NewAnalyzer(g.Board)
		g.GameInitialized = true
	}
	return nil
}
func (g *TicTacToeGame) CheckIsValidMove(mov string) (bool, int, int) {
	split := strings.Split(mov, "")
	if len(split) > 2 || len(split) == 1 {
		fmt.Println("Invalid move :", mov)
		return false, -1, -1
	}
	r, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println("Invalid move :", mov)
		return false, -1, -1
	}
	c, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Println("Invalid move:", mov)
		return false, -1, -1
	}
	r = r - 1
	c = c - 1

	switch {
	case r < 0, r >= g.Board.Size:
		fmt.Println("Invalid move:", mov)
		return false, -1, -1
	}

	switch {
	case c < 0, c >= g.Board.Size:
		fmt.Println("Invalid move:", mov)
		return false, -1, -1
	}
	if g.Board.GameBoard.Cells[r][c] != "" {
		fmt.Println(mov, "is already occupied on the board !! ")
		return false, -1, -1
	}
	return true, r, c
}
func (g *TicTacToeGame) MakeMove(mov string, mark b.Mark) bool {
	isValid, r, c := g.CheckIsValidMove(mov)
	if isValid {
		if g.Board.GameBoard.Cells[r][c] == "" {
			g.Board.GameBoard.Cells[r][c] = string(mark)
			return true
		}
	}
	return false
}

func (g *TicTacToeGame) ShowMenu() {
	fmt.Printf("Choose x/o for player1 %v \n", g.Player.Player1)
	var input string
	fmt.Scanln(&input)
	if input != "x" && input != "o" {
		fmt.Println("Please provide valid input to start the game ")
		g.ShowMenu()
	} else {
		g.Board.GameBoard.CurrentMark = b.Mark(input)
		g.Player.Players[g.Player.Player1] = b.Mark(input)
	}
}

func (g *TicTacToeGame) GameOver() {
	fmt.Println("End Game!!")
	os.Exit(0)
}
func (g *TicTacToeGame) getMarkForCurrentPlayer() string {
	if g.Board.GameBoard.CurrentMark == g.Player.Players[g.Player.Player1] {
		fmt.Println("Current player : ", g.Player.Player1)
	} else {
		fmt.Println("Current player : ", g.Player.Player2)
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
		successfulMove := g.MakeMove(input, g.Board.GameBoard.CurrentMark)
		if !successfulMove {
			return false
		}
		gameStatus := g.Result.CheckWinning()
		if gameStatus == r.Win {
			g.Board.ShowBoard()
			if g.Player.Players[g.Player.Player1] == g.Board.GameBoard.CurrentMark {
				fmt.Println(g.Player.Player1, "wins!")
			} else {
				fmt.Println(g.Player.Player2, "wins!")
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
	if g.Board.GameBoard.CurrentMark == b.X {
		g.Board.GameBoard.CurrentMark = b.O
	} else {
		g.Board.GameBoard.CurrentMark = b.X
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
