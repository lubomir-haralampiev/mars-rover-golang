package instructions

import (
	"fmt"
	"mars-rover-golang/constants"
	plateauType "mars-rover-golang/plateau"
	roverType "mars-rover-golang/rover"
)

// ProcessInstruction processes the given instruction
func ProcessInstruction(instruction string, rover roverType.Rover, plateau plateauType.Plateau) roverType.Rover {
	switch instruction {
	case constants.TurnLeft:
		return TurnLeft(rover)
	case constants.TurnRight:
		return TurnRight(rover)
	case constants.Move:
		return Move(rover, plateau)
	default:
		panic(fmt.Errorf("unknown instruction %s", instruction))

	}
}
