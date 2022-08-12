package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type Entity interface {
	GetID() uuid.UUID
	GetCategory() string
	Draw() rune
	Move()
	GetPosition() (x, y int)
	HasVelocity() bool
	GetVelocity() (vx, vy int)
	GetColor() tcell.Color
}
