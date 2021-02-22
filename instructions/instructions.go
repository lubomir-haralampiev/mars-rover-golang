package instructions

import (
	"fmt"
	"mars-rover-golang/constants"
	plateauType "mars-rover-golang/plateau"
	roverType "mars-rover-golang/rover"
)

// ProcessInstruction processes the given instruction
//
// Note on using "panic" - panic may not always be the better choice.
// If an instruction can not be processed but we want proceed with the next one
// then we can recover the panic here and retun an Error as a second return value.
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
