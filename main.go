package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func main() {
	rand.Seed(time.Now().UnixNano())

	framerate, enemies, exit := getOptions()
	if exit {
		os.Exit(0)
	}

	s, err := getScreen()
	if err != nil {
		log.Fatalf("could not init screen: %v", err)
	}

	w, h := s.Size()
	bee := NewBee(w/2, h/2)
	world := NewWorld(w, h, bee, enemies)

	options := &GameOptions{
		Framerate: framerate,
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

func getOptions() (framerate float32, enemies int, quit bool) {
	app := tview.NewApplication()
	form := tview.NewForm().
		AddDropDown("Framerate", []string{"1", "15", "30", "60"}, 2, func(option string, optionIndex int) {
			fps, _ := strconv.ParseFloat(option, 32)
			framerate = float32(fps)
		}).
		AddDropDown("Enemies", []string{"1", "5", "10", "25"}, 1, func(option string, optionIndex int) {
			e, _ := strconv.Atoi(option)
			enemies = e
		}).
		AddButton("Start", func() {
			app.Stop()
		}).
		AddButton("Quit", func() {
			app.Stop()
			quit = true
		})
	form.SetBorder(true).SetTitle(" beeyond - Game Options ").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}

	return framerate, enemies, quit
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
