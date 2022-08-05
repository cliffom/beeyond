package main

type Position struct {
	x int
	y int
}

// GetPosition returns a position as (x, y)
func (p *Position) GetPosition() (x, y int) {
	return p.x, p.y
}
