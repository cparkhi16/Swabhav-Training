package main

import (
	g "cmdgame/game"
	"cmdgame/tictactoegame"
)

func StartGame(game g.Game) {
	game.Play()
}
func main() {
	t := tictactoegame.New()
	StartGame(t)
}
