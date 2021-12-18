package player

import (
	b "cmdgame/board"
	"fmt"
)

type Player struct {
	Name string
	Players map[string]b.Mark
	Mark b.Mark
	Marks [2]b.Mark
}

func NewPlayers() *Player {
	return &Player{}
}

func (p *Player) GetPlayerDetails() (string,error) {
	fmt.Println("Enter name of player")
	var n1 string
	fmt.Scanln(&n1)
	if n1==string(b.Empty){
		return n1 ,fmt.Errorf("please enter a valid name")
	}
	return n1,nil
}

func (p *Player) SetPlayerDetails(n1 string) {
	p.Name=n1
	p.Players[p.Name]=b.Empty
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