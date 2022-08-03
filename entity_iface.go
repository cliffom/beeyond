package main

import "github.com/gdamore/tcell/v2"

type Entity interface {
	Draw() rune
	GetPosition() (x, y int)
	GetColor() tcell.Color
}
