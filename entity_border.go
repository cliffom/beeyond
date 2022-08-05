package main

import "github.com/gdamore/tcell/v2"

type Border struct {
	Position
	Sprite
}

// Move for Border is a noop as a border cannot move
func (b *Border) Move() {}

// HasVelocity returns false since a static border has no Velocity
// We don't use the Velocity component for this entity but including
// function simplifies drawing a border
func (b *Border) HasVelocity() bool {
	return false
}

// GetVelocity returns the current velocity (vx, vy)
func (b *Border) GetVelocity() (vx, vy int) {
	return 0, 0
}

func NewBorder(x, y int) *Border {
	return &Border{
		Position: Position{
			x: x,
			y: y,
		},
		Sprite: Sprite{
			color:  tcell.ColorLightGrey,
			frames: []rune{'\u2588'},
			frame:  0,
		},
	}
}
