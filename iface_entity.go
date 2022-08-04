package main

import "github.com/gdamore/tcell/v2"

type Entity interface {
	Draw() rune
	Move()
	GetPosition() (x, y int)
	HasVelocity() bool
	GetVelocity() (vx, vy int)
	GetColor() tcell.Color
}
