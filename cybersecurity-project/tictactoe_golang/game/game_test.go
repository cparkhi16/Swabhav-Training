package game

import (
	"tictactoe/board"
	"tictactoe/mark"
	"tictactoe/player"
	"tictactoe/resultAnalyzer"
)

var board1 = board.NewBoard(3)
var resultAnalyzer1 = resultAnalyzer.NewResultAnalyzer(board1)

var playerA = player.NewPlayer("a", mark.Not)
var playerB = player.NewPlayer("b", mark.Cross)
var players []*player.Player

func TestPlayGame() {
	players = append(players, playerA)
	players = append(players, playerB)
	var currentPlayer = playerA
	var newGame = NewGame(players, currentPlayer, board1, resultAnalyzer1)
	newGame.PlayGame()
}
