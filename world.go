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
func NewWorld(width, height int, p Player, o *GameOptions) *World {
	grid := make([][]Entity, height)
	for i := range grid {
		grid[i] = make([]Entity, width)
	}

	w := &World{
		Grid:    grid,
		Player:  p,
		Enemies: make([]*Enemy, o.Enemies),
	}

	w.PlaceEntity(p)
	placeBorders(w)
	placeMountainRanges(width, height, o.Mountains, w)
	placeEnemies(width, height, o.Enemies, w)

	return w
}

/*****************************************************************************
 *                      Helper Functions for Borders                         *
 *****************************************************************************/

func placeBorders(w *World) {
	for i := range w.Grid {
		if i == 0 || i == len(w.Grid)-1 {
			for k := range w.Grid[i] {
				border := NewStaticBorder(k, i)
				w.PlaceEntity(border)
			}
		} else {
			leftBorder := NewStaticBorder(0, i)
			rightBorder := NewStaticBorder(len(w.Grid[i])-1, i)
			w.PlaceEntity(leftBorder)
			w.PlaceEntity(rightBorder)
		}
	}
}

/*****************************************************************************
 *                      Helper Functions for Enemies                         *
 *****************************************************************************/

func placeEnemies(width, height, enemies int, w *World) {
	i := 0
	for i < enemies {
		x := rand.Intn(width)
		y := rand.Intn(height)
		if *w.GetCellAt(x, y) == nil {
			w.Enemies[i] = NewEnemy(x, y)
			w.PlaceEntity(w.Enemies[i])
			i++
		}
	}
}

/*****************************************************************************
 *                  Helper Functions for Mountain Ranges                     *
 *****************************************************************************/

func placeMountainRanges(width, height, ranges int, w *World) {
	mtnRangeWidth := 5
	mtnRangeHeight := 5

	for i := 0; i < ranges; i++ {
		x := rand.Intn(width-mtnRangeWidth-1) + 1
		y := rand.Intn(height-mtnRangeHeight-1) + 1
		placeMountain(x, y, mtnRangeWidth, mtnRangeHeight, w)
	}
}

func placeMountain(x, y, width, height int, w *World) {
	grid := make([][]int, width)
	for i := range grid {
		grid[i] = make([]int, height)
		for j := range grid[i] {
			if rand.Intn(6) < 3 {
				mountain := NewStaticMountain(x+i, y+j)
				w.PlaceEntity(mountain)
			}
		}
	}
}
