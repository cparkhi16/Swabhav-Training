package game

import (
	"fmt"
	"tictactoe/board"
	"tictactoe/player"
	"tictactoe/result"
	"tictactoe/resultAnalyzer"
)

type Game struct {
	players        []*player.Player
	currentPlayer  *player.Player
	board          *board.Board
	resultAnalyzer *resultAnalyzer.ResultAnalyzer
}

func NewGame(players []*player.Player, currentPlayer *player.Player, board *board.Board, resultAnalyzer *resultAnalyzer.ResultAnalyzer) *Game {
	return &Game{
		players:        players,
		currentPlayer:  currentPlayer,
		board:          board,
		resultAnalyzer: resultAnalyzer,
	}
}

func (g *Game) PlayGame() {
	g.currentPlayer = g.players[0]
	var location int = -1
	var currentRow int
	var currentCol int
	var newResult result.Result
	for {
	start:
		fmt.Printf("For Player-%s Enter location-", g.currentPlayer.GetName())
		_, err := fmt.Scanf("%d\n", &location)
		if err != nil {
			fmt.Scanf("%d", &location)
		}
		if location < 0 {
			fmt.Println("Invalid entry please try again!")
			goto start
		}
		ok := g.board.MakeEntryToBoardSingleDigit(location, g.currentPlayer.GetMark())
		if !ok {
			fmt.Println("Invalid entry please try again!")
			goto start
		}
		currentRow = location / g.board.GetSize()
		currentCol = location % g.board.GetSize()
		fmt.Println(currentRow, currentCol)
		g.board.DisplayBoard()
		newResult = g.resultAnalyzer.CheckBoard(currentRow, currentCol)
		if newResult == result.Winner {
			fmt.Println(g.currentPlayer.GetName(), " is a winner!")
			break
		} else if newResult == result.Tie {
			fmt.Println("game tie!")
			break
		}
		if g.currentPlayer == g.players[0] {
			g.currentPlayer = g.players[1]
		} else {
			g.currentPlayer = g.players[0]
		}
	}
}
