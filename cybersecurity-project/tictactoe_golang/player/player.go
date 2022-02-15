package player

import (
	"tictactoe/mark"
)

type Player struct {
	name string
	mark mark.Mark
}

func NewPlayer(name string, mark mark.Mark) *Player {
	return &Player{
		name: name,
		mark: mark,
	}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetMark() mark.Mark {
	return p.mark
}

func (p *Player) SetName(name string) {
	p.name = name
}

func (p *Player) SetMark(mark mark.Mark) {
	p.mark = mark
}
