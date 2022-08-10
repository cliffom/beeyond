package main

import (
	"fmt"
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

	// Initialize a new tview form to get user options
	options := &GameOptions{}
	exit := getOptions(options)
	if exit {
		os.Exit(0)
	}

	s, err := getScreen()
	if err != nil {
		log.Fatalf("could not init screen: %v", err)
	}

	game := NewGame(s, options)

	// listen for events
	evt := make(chan tcell.Event)
	quit := make(chan struct{})
	go s.ChannelEvents(evt, quit)

	go game.Run()

	for {
		game.HandleEvent(evt, quit)
	}
}

// getOptions draws a tcell.View to get game options
func getOptions(options *GameOptions) (quit bool) {
	app := tview.NewApplication()
	optionsForm := tview.NewForm().
		AddDropDown("Framerate", []string{"1", "15", "30", "60"}, 2, func(option string, optionIndex int) {
			fps, _ := strconv.ParseFloat(option, 32)
			options.Framerate = float32(fps)
		}).
		AddDropDown("Enemies", []string{"1", "5", "10", "25"}, 1, func(option string, optionIndex int) {
			e, _ := strconv.Atoi(option)
			options.Enemies = e
		}).
		AddButton("Start", func() {
			app.Stop()
		}).
		AddButton("Quit", func() {
			app.Stop()
			quit = true
		})
	optionsForm.SetBorder(true).
		SetTitle(" Game Options ").
		SetTitleAlign(tview.AlignCenter)

	descBox := tview.NewTextView()
	beeyondText, _ := os.ReadFile("welcome.txt")
	fmt.Fprintln(descBox, string(beeyondText))
	descBox.SetBorder(true)
	descBox.SetTitle(" beeyond ")

	flex := tview.NewFlex().
		AddItem(descBox, 0, 1, false).
		AddItem(tview.NewBox(), 1, 1, false).
		AddItem(optionsForm, 32, 1, true)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

	return quit
}

// getScreen initializes and returns a new tcell.Screen
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
