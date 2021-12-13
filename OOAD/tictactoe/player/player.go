package player

import (
	b "cmdgame/board"
	"fmt"
)

type Player struct {
	Player1 string
	Player2 string
	Players map[string]b.Mark
}

func NewPlayers() *Player {
	return &Player{}
}

func (p *Player) GetPlayerDetails() (string, string) {
	fmt.Println("Enter name of player1")
	var n1 string
	fmt.Scanln(&n1)
	p.Player1 = n1
	fmt.Println("Enter name of player2")
	var n2 string
	fmt.Scanln(&n2)
	p.Player2 = n2
	return n1, n2
}

func (p *Player) SetPlayerDetails(n1, n2 string) {
	p.Player1 = n1
	p.Player2 = n2
}
