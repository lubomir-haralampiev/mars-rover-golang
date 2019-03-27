package instructions

import (
	"fmt"
	"mars-rover-golang/constants"
	roverType "mars-rover-golang/rover"
)

// Move makes the rover move forward
func Move(rover roverType.Rover) roverType.Rover {
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

	return roverCopy
}

/*
my implementation in node js:

const ensureMovementIsPossible = (rover, plateau) => {
    if (plateau.rovers.find(existingRover => existingRover.x === rover.x && existingRover.y === rover.y)) {
        throw Error('At this position is already another rover');
    }

    if (rover.x > plateau.sizeX || rover.y > plateau.sizeY || rover.x < 0 || rover.y < 0) {
        throw Error('Reached the end of the plateau, can not move forward');
    }

    return rover;
};

export default (rover, plateau) => {
    switch (rover.orientation) {
        case NORTH:
            return ensureMovementIsPossible(createRover(rover.x, rover.y + 1, rover.orientation), plateau);
        case WEST:
            return ensureMovementIsPossible(createRover(rover.x - 1, rover.y, rover.orientation), plateau);
        case SOUTH:
            return ensureMovementIsPossible(createRover(rover.x, rover.y - 1, rover.orientation), plateau);
        case EAST:
            return ensureMovementIsPossible(createRover(rover.x + 1, rover.y, rover.orientation), plateau);
        default:
            throw Error(`Can not move the rover because the orientation "${rover.orientation}" is unknown`);
    }
};
*/
