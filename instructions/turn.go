package instructions

import (
	"fmt"
	"mars-rover-golang/constants"
	roverType "mars-rover-golang/rover"
)

// Another solution(maybe a more elegant one) is to convert the strings to degrees,
// calculate the new degrees and convert them back to strings. This way we could support not only 90 degrees.
// Keep it it simpler for now as the main goal is to get in touch with Go.

var transitionsLeft = map[string]string{
	constants.North: constants.West,
	constants.West:  constants.South,
	constants.South: constants.East,
	constants.East:  constants.North,
}

var transitionsRight = map[string]string{
	constants.North: constants.East,
	constants.West:  constants.North,
	constants.South: constants.West,
	constants.East:  constants.South,
}

func getNewOrientation(rover roverType.Rover, transitionMap map[string]string) string {
	newOrientation, ok := transitionMap[rover.Orientation]

	if ok == false {
		panic(fmt.Errorf(`Unsupported orientation "%v" of the Rover "%v"`, rover.Orientation, rover))
	}

	return newOrientation
}

// TurnLeft makes a rover turn left
func TurnLeft(rover roverType.Rover) roverType.Rover {
	return roverType.Rover{
		X:           rover.X,
		Y:           rover.Y,
		Orientation: getNewOrientation(rover, transitionsLeft),
	}
}

// TurnRight makes a rover turn right
func TurnRight(rover roverType.Rover) roverType.Rover {
	return roverType.Rover{
		X:           rover.X,
		Y:           rover.Y,
		Orientation: getNewOrientation(rover, transitionsRight),
	}
}
