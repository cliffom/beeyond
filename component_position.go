package main

type Position struct {
	x int
	y int
}

// SetPosition sets a position as (x, y)
func (p *Position) SetPosition(x, y int) {
	p.x = x
	p.y = y
}

// GetPosition returns a position as (x, y)
func (p *Position) GetPosition() (x, y int) {
	return p.x, p.y
}
