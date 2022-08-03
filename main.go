package main

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

func main() {
	s, err := getScreen()
	if err != nil {
		log.Fatalf("could not init screen: %v", err)
	}

	world := NewWorld(s.Size())
	bee := NewBee()

	world.PlaceEntity(bee)

	for i := range world.Grid {
		if i == 0 || i == len(world.Grid)-1 {
			for k := range world.Grid[i] {
				border := NewBorder(k, i)
				world.PlaceEntity(border)
			}
		} else {
			leftBorder := NewBorder(0, i)
			rightBorder := NewBorder(len(world.Grid[i])-1, i)
			world.PlaceEntity(leftBorder)
			world.PlaceEntity(rightBorder)
		}
	}

	// listen for events
	evt := make(chan tcell.Event)
	quit := make(chan struct{})
	go s.ChannelEvents(evt, quit)

	go func(s tcell.Screen, w *World) {
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
			time.Sleep(24 * time.Millisecond)
		}
	}(s, world)

	for {
		eventHandler(evt, quit, s, world, bee)
	}
}

func getScreen() (tcell.Screen, error) {
	// Initialize a new screen
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := s.Init(); err != nil {
		return nil, err
	}

	// Set our default styles
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	return s, nil
}

func eventHandler(evt chan tcell.Event, quit chan struct{}, s tcell.Screen, w *World, b *Bee) {
	select {
	case ev := <-evt:
		switch event := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			handleKeyPress(event.Key(), s, w, b)
		}
	case <-quit:
		s.Fini()
		os.Exit(0)
	}
}

// handleKeyPress handles key press events from the user
func handleKeyPress(k tcell.Key, s tcell.Screen, w *World, b *Bee) {
	switch k {
	case tcell.KeyUp:
		if !move(UP, 0, -1, b, w) {
			s.Beep()
		}
	case tcell.KeyRight:
		if !move(RIGHT, 1, 0, b, w) {
			s.Beep()
		}
	case tcell.KeyDown:
		if !move(DOWN, 0, 1, b, w) {
			s.Beep()
		}
	case tcell.KeyLeft:
		if !move(LEFT, -1, 0, b, w) {
			s.Beep()
		}
	case tcell.KeyEscape:
		s.Clear()
		s.Fini()
		os.Exit(0)
	}
}

func move(d, vx, vy int, b *Bee, w *World) bool {
	x, y := b.GetPosition()
	cell := *w.GetCellAt(x+vx, y+vy)
	w.ClearCellAt(b.GetPosition())
	moved := b.Move(d, cell)
	w.PlaceEntity(b)
	return moved
}
