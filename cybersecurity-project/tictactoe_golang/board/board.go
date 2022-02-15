package board

import (
	"fmt"
	"tictactoe/cell"
	"tictactoe/mark"
)

type Board struct {
	cells []*cell.Cell
	size  uint8
}

func (b *Board) getCellLocation(row int, col int) int {
	return (row * int(b.size)) + col
}

func NewBoard(size uint8) *Board {
	cells := make([]*cell.Cell, size*size)
	for i := 0; i < int(size)*int(size); i++ {
		cells[i] = cell.NewCell(mark.Empty)
	}
	return &Board{
		cells: cells,
		size:  size,
	}
}

func (b *Board) GetSize() int {
	return int(b.size)
}

func (b *Board) DisplayBoard() {
	s := int(b.size)
	count := 1
	for i := 0; i < s*s; i++ {
		fmt.Printf("%v", *b.cells[i])
		count++
		if count > s {
			fmt.Println()
			count = 1
		}
	}
}

func (b *Board) GetEntryAt(row int, col int) mark.Mark {
	return b.cells[b.getCellLocation(row, col)].GetMark()
}

func (b *Board) MakeEntryToBoard(row int, col int, newMark mark.Mark) bool {
	if row > int(b.size)-1 || col > int(b.size)-1 || b.cells[b.getCellLocation(row, col)].GetMark() != mark.Empty {
		return false
	}
	b.cells[b.getCellLocation(row, col)].SetMark(newMark)
	return true
}

func (b *Board) MakeEntryToBoardSingleDigit(location int, newMark mark.Mark) bool {
	if location > int(b.size)*int(b.size)-1 || b.cells[location].GetMark() != mark.Empty || location < 0 {
		return false
	}
	b.cells[location].SetMark(newMark)
	return true
}

func (b *Board) IsFull() bool {
	for i := 0; i < int(b.size)*int(b.size); i++ {
		if b.cells[i].GetMark() == mark.Empty {
			return false
		}
	}
	return true
}

func (b *Board) IsEmpty() bool {
	for i := 0; i < int(b.size)*int(b.size); i++ {
		if b.cells[i].GetMark() != mark.Empty {
			return false
		}
	}
	return true
}
