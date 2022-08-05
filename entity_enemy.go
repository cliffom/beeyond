package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

var enemyFrames = []rune{
	'\u25F4',
	'\u25F5',
	'\u25F6',
	'\u25F7',
}

type Enemy struct {
	Position
	Velocity
	Sprite
	tick  int
	delay int
}

// Draw returns the rune that represents the current state of our Bee
func (e *Enemy) Draw() rune {
	if e.frame > len(e.frames)-1 {
		e.frame = 0
	}
	frame := e.frames[e.frame]
	e.frame++
	return frame
}

// Move checks for an entity in what would be our Bee's occupying
// cell post-movement. If the cell-to-be-occupied has no occupying
// entity, update the Bee's position based on movement vectors.
func (e *Enemy) Move() {
	e.x += e.vx
	e.y += e.vy
	e.vx = 0
	e.vy = 0
}

func (e *Enemy) HasVelocity() bool {
	if e.tick < e.delay {
		e.tick++
		return false
	}
	e.tick = 0
	x := rand.Intn(3)
	y := rand.Intn(3)
	switch x {
	case 0:
		e.vx = 0
	case 1:
		e.vx = -1
	case 2:
		e.vx = 1
	}
	switch y {
	case 0:
		e.vy = 0
	case 1:
		e.vy = -1
	case 2:
		e.vy = 1
	}

	return true
}

// NewEnemy returns a new Bee
func NewEnemy(x, y int) *Enemy {
	return &Enemy{
		delay: rand.Intn(10) + 10,
		Position: Position{
			x: x,
			y: y,
		},
		Velocity: Velocity{
			vx: 0,
			vy: 0,
		},
		Sprite: Sprite{
			frames: enemyFrames,
			frame:  0,
			color:  tcell.ColorRed,
		},
	}
}
