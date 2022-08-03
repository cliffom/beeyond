package main

type World struct {
	Grid [][]Entity
}

func (w *World) PlaceEntity(e Entity) bool {
	x, y := e.GetPosition()
	if w.Grid[y][x] == nil {
		w.Grid[y][x] = e
		return true
	}

	return false
}

func (w *World) ClearCellAt(x, y int) {
	w.Grid[y][x] = nil
}

func (w *World) GetCellAt(x, y int) *Entity {
	return &w.Grid[y][x]
}

func NewWorld(w, h int) *World {
	grid := make([][]Entity, h)
	for i := range grid {
		grid[i] = make([]Entity, w)
	}

	return &World{
		Grid: grid,
	}
}
