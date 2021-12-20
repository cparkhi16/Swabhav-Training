package player

import (
	b "cmdgame/board"
	"fmt"
)

type Player struct {
	Name string
	Mark b.Mark
	Players map[string]b.Mark
}

func NewPlayers() *Player {
	return &Player{}
}


func (p *Player) SetPlayerDetails(n1 string) error{
	for k,_:=range p.Players{
		if k==n1{
			return fmt.Errorf("please enter different name of player")
		}
	}
	p.Name=n1
	p.Players[p.Name]=b.Empty
	return nil
}
func (p *Player)GetPlayerFromMap(m map[string]b.Mark,value b.Mark)(key string,ok bool){
	for k,v:=range m{
		if v==value{
			key=k
			ok=true
			return
		}
	}
	return
}