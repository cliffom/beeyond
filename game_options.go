package main

import "time"

type GameOptions struct {
	Framerate float32
}

// Frametime returns a time.Duration based on the given framerate
// ex: Framerate 30 is approximately a frametime of 0.033ms
func (o *GameOptions) Frametime() time.Duration {
	frametimeInSeconds := 1 / o.Framerate * float32(time.Second)
	return time.Duration(frametimeInSeconds)
}
