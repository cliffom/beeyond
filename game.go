package main

import (
	"os"
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
					if ent.HasVelocity() {
						w.ClearCellAt(ent.GetPosition())
						ent.Move()
						w.PlaceEntity(ent)
					}

					s.SetContent(j, i, ent.Draw(), nil, style)
				}
			}
		}
		s.Show()
		time.Sleep(g.Frametime)
	}
}

func (g *Game) HandleEvent(evt chan tcell.Event, quit chan struct{}) {
	s := g.Screen

	select {
	case ev := <-evt:
		switch event := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			g.HandleInput(event.Key())
		}
	case <-quit:
		s.Fini()
		os.Exit(0)
	}
}

func (g *Game) HandleInput(k tcell.Key) {
	w := g.World
	s := g.Screen

	switch k {
	case tcell.KeyUp:
		if !w.MovePlayer(UP) {
			s.Beep()
		}
	case tcell.KeyRight:
		if !w.MovePlayer(RIGHT) {
			s.Beep()
		}
	case tcell.KeyDown:
		if !w.MovePlayer(DOWN) {
			s.Beep()
		}
	case tcell.KeyLeft:
		if !w.MovePlayer(LEFT) {
			s.Beep()
		}
	case tcell.KeyEscape:
		s.Clear()
		s.Fini()
		os.Exit(0)
	}
}

func NewGame(s tcell.Screen, w *World) *Game {
	return &Game{
		Screen:    s,
		World:     w,
		Frametime: 24 * time.Millisecond,
	}
}
