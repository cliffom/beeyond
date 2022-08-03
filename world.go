package main

type World struct {
	Grid   [][]Entity
	Player Player
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

func NewWorld(w, h int, p Player) *World {
	grid := make([][]Entity, h)
	for i := range grid {
		grid[i] = make([]Entity, w)
	}

	// Place our player in the world
	x, y := p.GetPosition()
	grid[y][x] = p

	world := &World{
		Grid:   grid,
		Player: p,
	}

	// Initialize the borders of our world
	for i := range world.Grid {
		if i == 0 || i == len(world.Grid)-1 {
			for k := range world.Grid[i] {
				border := NewBorder(k, i)
				world.PlaceEntity(border)
			}
		} else {
			leftBorder := NewBorder(0, i)
			rightBorder := NewBorder(len(world.Grid[i])-1, i)
			world.PlaceEntity(leftBorder)
			world.PlaceEntity(rightBorder)
		}
	}

	return world
}
