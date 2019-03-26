package rover

// Rover defines how a rover looks like.
// Only responsible for holding the data,
// not responsible for validating it(e.g. the orientation).
// This should be part of the logic processing the instructions.
type Rover struct {
	X           int
	Y           int
	Orientation string
}
