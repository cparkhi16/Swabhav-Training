package player

import (
	"testing"
	"tictactoe/mark"
)

var newPlayer = NewPlayer("test", mark.Cross)

func TestNewPlayer(t *testing.T) {
	actualName := "test"
	expectedName := newPlayer.name
	actualMark := mark.Cross
	expectedMark := newPlayer.mark

	if actualName != expectedName {
		t.Errorf("expected:%v and actual:%v", expectedName, actualName)
	} else if actualMark != expectedMark {
		t.Errorf("expected:%v and actual:%v", expectedMark, actualMark)
	}
}

func TestGetMark(t *testing.T) {
	actual := mark.Cross
	expected := newPlayer.GetMark()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestGetName(t *testing.T) {
	actual := "test"
	expected := newPlayer.GetName()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}
