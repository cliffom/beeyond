package main

import "github.com/gdamore/tcell/v2"

type Sprite struct {
	color  tcell.Color
	frames []rune
	frame  int
}

// Draw returns the rune that represents the current state of our Bee
func (s *Sprite) Draw() rune {
	return s.frames[s.frame]
}

// GetColor returns the current color of a Sprite
func (s *Sprite) GetColor() tcell.Color {
	return s.color
}
