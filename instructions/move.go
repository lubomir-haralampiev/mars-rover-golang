package instructions

import (
	"errors"
	"fmt"
	"mars-rover-golang/constants"
	plateauType "mars-rover-golang/plateau"
	roverType "mars-rover-golang/rover"
)

func ensureMovementIsPossible(plateau plateauType.Plateau, rover roverType.Rover) roverType.Rover {
	for i := range plateau.Rovers {
		if plateau.Rovers[i].X == rover.X && plateau.Rovers[i].Y == rover.Y {
			panic(errors.New("Can not move the rover because the target position is occupied"))
		}
	}

	if rover.X > plateau.SizeX || rover.Y > plateau.SizeY || rover.X < 0 || rover.Y < 0 {
		panic(errors.New("Reached the end of the plateau, can not move forward"))
	}

	return rover
}

// Move makes the rover move forward
func Move(rover roverType.Rover, plateau plateauType.Plateau) roverType.Rover {
	// !!! the copyong only works because its values are all of primitive types which are referenced by value.
	// If the rover has properties of complex types, a deep copying must be implemented!
	roverCopy := rover

	switch rover.Orientation {
	case constants.North:
		roverCopy.Y++
	case constants.West:
		roverCopy.X--
	case constants.South:
		roverCopy.Y--
	case constants.East:
		roverCopy.X++
	default:
		panic(fmt.Errorf(`Cano not move the rover %v because the orientation "%v" is unknown`, rover, rover.Orientation))
	}

	return ensureMovementIsPossible(plateau, roverCopy)
}
