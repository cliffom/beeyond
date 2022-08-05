package main

type Velocity struct {
	vx int
	vy int
}

// HasVelocity determines if there is a non-zero velocity
// in either direction
func (v *Velocity) HasVelocity() bool {
	return v.vx != 0 || v.vy != 0
}

// GetVelocity returns the current velocity (vx, vy)
func (v *Velocity) GetVelocity() (vx, vy int) {
	return v.vx, v.vy
}

// SetVelocity sets a velocity to the incoming (vx, vy)
func (v *Velocity) SetVelocity(vx, vy int) {
	v.vx = vx
	v.vy = vy
}
