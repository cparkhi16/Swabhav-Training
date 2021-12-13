package board

type Cell struct {
	Cells       []string
	Marks       [2]Mark
	CurrentMark Mark
}

func NewCell() *Cell {
	return &Cell{}
}
