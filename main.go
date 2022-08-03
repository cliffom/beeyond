package main

import (
	"log"

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

	bee := NewBee()
	w, h := s.Size()
	world := NewWorld(w, h, bee)

	// Initialize the borders of our world
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

	game := NewGame(s, world)

	// listen for events
	evt := make(chan tcell.Event)
	quit := make(chan struct{})
	go s.ChannelEvents(evt, quit)

	go game.Run()

	for {
		game.HandleEvent(evt, quit)
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
