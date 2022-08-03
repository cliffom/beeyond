package main

import "github.com/gdamore/tcell/v2"

type Player interface {
	Entity
	UpdatePosition(d int, e Entity) bool
	HandleInput(k tcell.Key) bool
}
