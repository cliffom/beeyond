package main

type Delay struct {
	tick  int
	delay int
}

// NextTick returns true and resets Delay.Tick if Delay.tick == Delay.Delay
// otherwise it increments tick
func (d *Delay) NextTick() bool {
	if d.tick < d.delay {
		d.tick++
		return false
	}
	d.tick = 0
	return true
}
