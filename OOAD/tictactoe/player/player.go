package player

type Player struct {
	Player1 string
	Player2 string
}

func NewPlayers(p1, p2 string) *Player {
	return &Player{Player1: p1, Player2: p2}
}
