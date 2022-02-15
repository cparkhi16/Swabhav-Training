package main

import (
	"fmt"
	"tictactoe/board"
	"tictactoe/game"
	"tictactoe/mark"
	"tictactoe/player"
	"tictactoe/resultAnalyzer"
)

func main() {
	fmt.Printf("Enter board size-")
	var size int
getSizeAgain:
	fmt.Scanf("%d\n", &size)
	if size <= 2 {
		fmt.Printf("Invalid Entry! Minimum board size should be 3, Enter board size-")
		fmt.Scanf("%d", &size)
		goto getSizeAgain
	}
	board1 := board.NewBoard(uint8(size))
	resultAnalyzer1 := resultAnalyzer.NewResultAnalyzer(board1)

	var playerAname string
	var playerBname string
	var playerAmark string

	playerA := player.NewPlayer("a", mark.Not)
	playerB := player.NewPlayer("b", mark.Cross)

	fmt.Println("Enter playerA name-")
enternameAAgain:
	_, err := fmt.Scanf("%s\n", &playerAname)
	if playerAname == " " || playerAname == "" {
		fmt.Println("Invalid Entry, Enter playerA name-")
		goto enternameAAgain
	}

	fmt.Println("Enter playerA mark (cross or not)-")
entermarkAagain:
	n, err := fmt.Scanf("%s\n", &playerAmark)
	if err != nil || n != 1 {
		// handle invalid input
		fmt.Println(n, err)
	}

	if playerAmark != "cross" && playerAmark != "not" {
		fmt.Println("Invalid Entry, Enter playerA mark (cross or not)-")
		goto entermarkAagain
	} else {
		playerA.SetName(playerAname)
		if playerAmark == "cross" {
			playerA.SetMark(mark.Cross)
			playerB.SetMark(mark.Not)
		} else {
			playerA.SetMark(mark.Not)
			playerB.SetMark(mark.Cross)
		}
	}

	fmt.Println("Enter playerB name-")
enternameBAgain:
	fmt.Scanf("%s", &playerBname)
	if playerBname == " " || playerBname == "" {
		fmt.Println("Invalid Entry, Enter playerB name-")
		goto enternameBAgain
	}
	playerB.SetName(playerBname)

	var players []*player.Player
	players = append(players, playerA)
	players = append(players, playerB)
	currentPlayer := playerA

	ticgame := game.NewGame(players, currentPlayer, board1, resultAnalyzer1)
	ticgame.PlayGame()

}
