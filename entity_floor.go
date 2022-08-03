package main

var floorFrames = []rune{
	'\u2593',
}

type Floor struct {
	Position
	Sprite
}

func (f *Floor) Draw() rune {
	return f.frames[f.frame]
}

func (f *Floor) GetPosition() (x, y int) {
	return f.x, f.y
}

func NewFloor(x, y int) *Floor {
	return &Floor{
		Position: Position{
			x: x,
			y: y,
		},
		Sprite: Sprite{
			frames: floorFrames,
			frame:  0,
		},
	}
}
