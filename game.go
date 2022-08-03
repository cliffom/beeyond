package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen    tcell.Screen
	World     *World
	Frametime time.Duration
}

func (g *Game) Run() {
	s := g.Screen
	w := g.World
	for {
		s.Clear()
		for i := range w.Grid {
			for j, k := range w.Grid[i] {
				switch ent := k.(type) {
				case Entity:
					style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(ent.GetColor())
					s.SetContent(j, i, ent.Draw(), nil, style)
				}
			}
		}
		s.Show()
		time.Sleep(g.Frametime)
	}
}

func NewGame(s tcell.Screen, w *World) *Game {
	return &Game{
		Screen:    s,
		World:     w,
		Frametime: 24 * time.Millisecond,
	}
}
