package player

import (
	b "cmdgame/board"
	)

type Player struct {
	Name string
	Mark b.Mark
	//Players map[string]b.Mark
}

func NewPlayers() *Player {
	return &Player{}
}

func (p *Player)SetPlayerMark(m b.Mark){
	p.Mark=m
}

func (p *Player) SetPlayerDetails(n1 string) {
	p.Name=n1
	//p.Players[p.Name]=b.Empty
}
