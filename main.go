package main

import (
	"log"
	"math/rand"
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
	rand.Seed(time.Now().UnixNano())
	s, err := getScreen()
	if err != nil {
		log.Fatalf("could not init screen: %v", err)
	}

	w, h := s.Size()
	bee := NewBee(w/2, h/2)
	world := NewWorld(w, h, bee)

	options := &GameOptions{
		Framerate: 30,
	}
	game := NewGame(s, world, options)

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
