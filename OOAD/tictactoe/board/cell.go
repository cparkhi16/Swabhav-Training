package board

type Cell struct {
	//Cells       []string
	
	CurrentMark Mark
}

func NewCell(m Mark) *Cell {
	
	return &Cell{CurrentMark:m}
}
func (c *Cell)GetMark()Mark{
	return  c.CurrentMark
}

func (c *Cell)SetMark(m Mark){
	c.CurrentMark=m
}