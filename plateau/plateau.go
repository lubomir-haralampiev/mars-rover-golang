package plateau

import (
	"mars-rover-golang/rover"
)

// Plateau defines how a plateau looks like
type Plateau struct {
	SizeX int
	SizeY int
	// Intentionally not modeling the plateau as array of arrays with fixed size.
	// Because the only purpose is holding the rovers and check for collisions.
	// If visualisation is needed, then a real grid can still be done on the GUI side.
	Rovers []rover.Rover
}
