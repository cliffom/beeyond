package main

var beeFrames = []rune{
	'\u25B2',
	'\u25B6',
	'\u25BC',
	'\u25C0',
}

type Bee struct {
	Position
	Velocity
	Sprite
}

func (b *Bee) Draw() rune {
	return b.frames[b.frame]
}

func (b *Bee) Move(dir int, e Entity) bool {
	b.frame = dir

	if e == nil {
		switch dir {
		case UP:
			b.vx = 0
			b.vy = -1
		case RIGHT:
			b.vx = 1
			b.vy = 0
		case DOWN:
			b.vx = 0
			b.vy = 1
		case LEFT:
			b.vx = -1
			b.vy = 0
		}

		b.x += b.vx
		b.y += b.vy

		return true
	}

	return false
}

func (b *Bee) GetPosition() (x, y int) {
	return b.x, b.y
}

func NewBee() *Bee {
	return &Bee{
		Position: Position{
			x: 5,
			y: 5,
		},
		Velocity: Velocity{
			vx: 0,
			vy: 0,
		},
		Sprite: Sprite{
			frames: beeFrames,
			frame:  0,
		},
	}
}
