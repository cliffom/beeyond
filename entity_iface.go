package main

type Entity interface {
	Draw() rune
	GetPosition() (x, y int)
}
