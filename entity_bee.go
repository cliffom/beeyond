package main

import "github.com/gdamore/tcell/v2"

var beeFrames = []rune{
	'\u25B2',
	'\u25B6',
	'\u25BC',
	'\u25C0',
}

type Bee struct {
	Position
	Velocity
	Sprite
}

// Move checks for an entity in what would be our Bee's occupying
// cell post-movement. If the cell-to-be-occupied has no occupying
// entity, update the Bee's position based on movement vectors.
func (b *Bee) Move() {
	b.x += b.vx
	b.y += b.vy
	b.vx = 0
	b.vy = 0
}

// UpdatePosition updates a Bee's velocity (vx, vy) and position (x, y)
// based on an intended direction
func (b *Bee) UpdatePosition(d int, e Entity) bool {
	b.frame = d

	if e != nil {
		return false
	}

	switch d {
	case UP:
		b.SetVelocity(0, -1)
	case RIGHT:
		b.SetVelocity(1, 0)
	case DOWN:
		b.SetVelocity(0, 1)
	case LEFT:
		b.SetVelocity(-1, 0)
	}

	return true
}

// NewBee returns a new Bee
func NewBee(x, y int) *Bee {
	return &Bee{
		Position: Position{
			x: x,
			y: y,
		},
		Sprite: Sprite{
			frames: beeFrames,
			frame:  0,
			color:  tcell.ColorYellow,
		},
	}
}
