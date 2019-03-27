package plateau

import (
	"mars-rover-golang/rover"
)

// Plateau defines how a plateau looks like
type Plateau struct {
	SizeX  int
	SizeY  int
	Rovers []rover.Rover
}
