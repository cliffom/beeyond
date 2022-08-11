package main

import "github.com/gdamore/tcell/v2"

const (
	mountain = '\u26F0' // ⛰
	border   = '\u2588' // █
)

type Static struct {
	Position
	Sprite
}

// Move for Static is a noop as a static cannot move
func (s *Static) Move() {}

// HasVelocity returns false since a static has no Velocity
// We don't use the Velocity component for this entity but including
// function simplifies drawing a static
func (s *Static) HasVelocity() bool {
	return false
}

// GetVelocity returns the current velocity (vx, vy)
func (s *Static) GetVelocity() (vx, vy int) {
	return 0, 0
}

func NewStatic(x, y int, r rune) *Static {
	return &Static{
		Position: Position{
			x: x,
			y: y,
		},
		Sprite: Sprite{
			color:  tcell.ColorLightGrey,
			frames: []rune{r},
			frame:  0,
		},
	}
}

func NewStaticBorder(x, y int) *Static {
	return NewStatic(x, y, border)
}

func NewStaticMountain(x, y int) *Static {
	return NewStatic(x, y, mountain)
}
