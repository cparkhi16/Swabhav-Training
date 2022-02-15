package cell

import (
	"tictactoe/mark"
)

type Cell struct {
	mark mark.Mark
}

func NewCell(mark mark.Mark) *Cell {
	return &Cell{
		mark: mark,
	}
}

func (c *Cell) GetMark() mark.Mark {
	return c.mark
}

func (c *Cell) SetMark(newMark mark.Mark) {
	c.mark = newMark
}
