package main

import "github.com/gdamore/tcell/v2"

type Sprite struct {
	color  tcell.Color
	frames []rune
	frame  int
}
