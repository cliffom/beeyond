package main

import "time"

type Delay struct {
	ticker time.Ticker
}

// Tick determines whether or not an update should be allowed
func (d *Delay) Tick() bool {
	for {
		select {
		case <-d.ticker.C:
			return true
		default:
			return false
		}
	}
}
