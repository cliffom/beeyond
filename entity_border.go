package main

var borderFrames = []rune{
	'\u2588',
}

type Border struct {
	Position
	Sprite
}

func (b *Border) Draw() rune {
	return b.frames[b.frame]
}

func (b *Border) GetPosition() (x, y int) {
	return b.x, b.y
}

func NewBorder(x, y int) *Border {
	return &Border{
		Position: Position{
			x: x,
			y: y,
		},
		Sprite: Sprite{
			frames: borderFrames,
			frame:  0,
		},
	}
}
