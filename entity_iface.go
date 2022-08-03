package main

import "github.com/gdamore/tcell/v2"

type Entity interface {
	Draw() rune
	Move()
	GetPosition() (x, y int)
	HasVelocity() bool
	GetColor() tcell.Color
}
