package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Enemy struct {
	Position
	Velocity
	Sprite
	Delay
}

// Draw returns the current animation frame but also increments
// the frame for the next draw call
func (e *Enemy) Draw() rune {
	frame := e.frames[e.frame]

	e.frame++
	if e.frame > len(e.frames)-1 {
		e.frame = 0
	}

	return frame
}

// Move checks for an entity in what would be our Enemy's occupying
// cell post-movement. If the cell-to-be-occupied has no occupying
// entity, update the Enemy's position based on movement vectors.
func (e *Enemy) Move() {
	e.SetPosition(e.x+e.vx, e.y+e.vy)
}

// HasVelocity overrides Velocity.HasVelocity to randomize
// movement direction and speed (time between movement intervals)
func (e *Enemy) HasVelocity() bool {
	if !e.Tick() {
		return false
	}

	e.SetRandomVelocity()
	return true
}

// NewEnemy returns a new Enemy
func NewEnemy(x, y int) *Enemy {
	delay := time.Duration(rand.Intn(501)+250) * time.Millisecond
	return &Enemy{
		Delay: Delay{
			ticker: *time.NewTicker(delay),
		},
		Position: Position{
			x: x,
			y: y,
		},
		Velocity: Velocity{
			vx: 0,
			vy: 0,
		},
		Sprite: Sprite{
			frames: getEnemyFrames(),
			frame:  0,
			color:  tcell.ColorRed,
		},
	}
}

// getEnemyFrames is a helper function that adds frames of each animation
// in order to slow down the perceived time between animation updates
func getEnemyFrames() []rune {
	enemyFrames := []rune{
		'\u25F4',
		'\u25F5',
		'\u25F6',
		'\u25F7',
	}

	numFramesPerState := rand.Intn(7) + 4
	totalFrames := len(enemyFrames) * numFramesPerState

	runes := make([]rune, totalFrames)

	for i, v := range enemyFrames {
		for j := numFramesPerState * i; j < numFramesPerState+(i*numFramesPerState); j++ {
			runes[j] = v
		}
	}

	return runes
}
