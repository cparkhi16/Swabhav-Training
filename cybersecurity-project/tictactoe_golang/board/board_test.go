package board

import (
	"testing"
	"tictactoe/mark"
)

var newBoard = NewBoard(3)

func TestNewBoard(t *testing.T) {
	actual := true
	expected := newBoard.IsEmpty()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestGetSize(t *testing.T) {
	actual := 3
	expected := newBoard.GetSize()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestGetEntryAt(t *testing.T) {
	actual := mark.Empty
	expected := newBoard.GetEntryAt(0, 0)
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestMakeEntryToBoard(t *testing.T) {
	actual := mark.Cross
	newBoard.MakeEntryToBoard(0, 0, mark.Cross)
	expected := newBoard.GetEntryAt(0, 0)
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestIsFull(t *testing.T) {
	actual := false
	expected := newBoard.IsFull()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestIsEmpty(t *testing.T) {
	actual := false
	expected := newBoard.IsEmpty()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}
