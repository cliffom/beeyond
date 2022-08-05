package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Game represents our main game and handles
// rendering, updating the world state, inputs, etc
type Game struct {
	Screen  tcell.Screen
	World   *World
	Options *GameOptions
}

// Run is the main gameloop - takes a world grid of entities, handles movement,
// and renders. Time between loop iterations is determined by Game.Frametime
func (g *Game) Run() {
	s := g.Screen
	w := g.World

	// main loop
	for {
		// iterate through our Grid Y-axis
		for i := range w.Grid {
			// iterate through our cells w.Grid[j][i]
			for j, k := range w.Grid[i] {
				switch ent := k.(type) {
				case Entity:
					style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(ent.GetColor())

					// If the entity in the current cell has a velocity
					// then attempt movement
					if ent.HasVelocity() {
						x, y := ent.GetPosition()
						vx, vy := ent.GetVelocity()

						// Check the contents in the cell our entity
						// is attempting to occupy. Only allow movement
						// if the cell is unoccupied
						if *w.GetCellAt(x+vx, y+vy) == nil {
							w.ClearCellAt(x, y)
							ent.Move()
							w.PlaceEntity(ent)
						}
					}

					// set the content of the screen at the current cell
					// to the entities current animation frame
					s.SetContent(j, i, ent.Draw(), nil, style)
				default:
					// handle empty cells
					style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(0)
					s.SetContent(j, i, '\u0020', nil, style)
				}
			}
		}

		// render screen based on updates and delay before next iteration
		s.Show()
		time.Sleep(g.Options.Frametime())
	}
}

// HandleEvent is what handles incoming events (key presses, screen resizes, etc)
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

// HandleInputs is what handles user input events (from keyboard)
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

// NewGame initlizes a game given a screen to render on
// and a world
func NewGame(s tcell.Screen, w *World, o *GameOptions) *Game {
	return &Game{
		Screen:  s,
		World:   w,
		Options: o,
	}
}
