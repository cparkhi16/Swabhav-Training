package cell

import (
	"testing"
	"tictactoe/mark"
)

var newCell = NewCell(mark.Cross)

func TestNewCell(t *testing.T) {
	actual := mark.Cross
	expected := newCell.mark
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestGetMark(t *testing.T) {
	actual := mark.Cross
	expected := newCell.GetMark()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestSetMark(t *testing.T) {
	actual := mark.Empty
	newCell.SetMark(mark.Empty)
	expected := newCell.GetMark()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}
