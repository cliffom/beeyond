package main

type Player interface {
	Entity
	UpdatePosition(d int, e Entity) bool
}
