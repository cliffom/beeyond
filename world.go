package main

import "math/rand"

// World represents our world as a 2-dimensional grid
// and a user-controlled player
type World struct {
	Grid    [][]Entity
	Enemies []*Enemy
	Player  Player
}

// PlaceEntity places an entity on the world grid at
// the entity's current position as long as there isn't
// an existing entity occupying the grid position
func (w *World) PlaceEntity(e Entity) bool {
	x, y := e.GetPosition()
	if w.Grid[y][x] == nil {
		w.Grid[y][x] = e
		return true
	}

	return false
}

// ClearCellAt removes any entity from the cell position (x, y)
func (w *World) ClearCellAt(x, y int) {
	w.Grid[y][x] = nil
}

// GetCellAt returns the entity at the cell position (x, y)
func (w *World) GetCellAt(x, y int) *Entity {
	return &w.Grid[y][x]
}

// MovePlayer will attempt to update the position of a player
// based on the incoming direction. Passes the contents of the cell
// in the would-be new position as the player's awareness
func (w *World) MovePlayer(d int) bool {
	var vx, vy int
	switch d {
	case UP:
		vx = 0
		vy = -1
	case RIGHT:
		vx = 1
		vy = 0
	case DOWN:
		vx = 0
		vy = 1
	case LEFT:
		vx = -1
		vy = 0
	}

	x, y := w.Player.GetPosition()
	cell := *w.GetCellAt(x+vx, y+vy)
	return w.Player.UpdatePosition(d, cell)
}

// NewWorld initializes and returns a World. Responsible for
// drawing the world border and placing the player in their
// starting position
func NewWorld(w, h int, p Player, o *GameOptions) *World {
	grid := make([][]Entity, h)
	for i := range grid {
		grid[i] = make([]Entity, w)
	}

	// Place our player in the world
	x, y := p.GetPosition()
	grid[y][x] = p

	world := &World{
		Grid:    grid,
		Player:  p,
		Enemies: make([]*Enemy, o.Enemies),
	}

	placeMountainRanges(w, h, o.Mountains, world)

	// Initialize the borders of our world
	for i := range world.Grid {
		if i == 0 || i == len(world.Grid)-1 {
			for k := range world.Grid[i] {
				border := NewStaticBorder(k, i)
				world.PlaceEntity(border)
			}
		} else {
			leftBorder := NewStaticBorder(0, i)
			rightBorder := NewStaticBorder(len(world.Grid[i])-1, i)
			world.PlaceEntity(leftBorder)
			world.PlaceEntity(rightBorder)
		}
	}

	// Put some enemies into our world
	for i := o.Enemies - 1; i >= 0; i-- {
		ex := rand.Intn(w)
		ey := rand.Intn(h)
		if *world.GetCellAt(ex, ey) == nil {
			world.Enemies[i] = NewEnemy(ex, ey)
			world.PlaceEntity(world.Enemies[i])
		}
	}

	return world
}

/*****************************************************************************
 *                  Helper Functions for Mountain Ranges                     *
 *****************************************************************************/

func placeMountainRanges(w, h, n int, world *World) {
	mtnRangeWidth := 5
	mtnRangeHeight := 5

	for i := 0; i < n; i++ {
		x := rand.Intn(w-mtnRangeWidth-1) + 1
		y := rand.Intn(h-mtnRangeHeight-1) + 1
		placeMountainRange(x, y, mtnRangeWidth, mtnRangeHeight, world)
	}
}

func placeMountainRange(x, y, w, h int, world *World) {
	grid := make([][]int, w)
	for i := range grid {
		grid[i] = make([]int, h)
		for j := range grid[i] {
			if rand.Intn(6) < 3 {
				mountain := NewStaticMountain(x+i, y+j)
				world.PlaceEntity(mountain)
			}
		}
	}
}
