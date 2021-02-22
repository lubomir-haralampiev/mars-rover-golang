package lineprocessor

import (
	"fmt"
	"log"
	"mars-rover-golang/instructions"
	plateauType "mars-rover-golang/plateau"
	roverType "mars-rover-golang/rover"
	"strconv"
	"strings"
)

func convertStringToInt(size string) int {
	intValue, err := strconv.Atoi(size)
	if err != nil {
		panic(fmt.Errorf("Can not convert %v to string", size))
	}

	return intValue
}

func readPlateauSize(sizesLine string) (int, int) {
	sizes := strings.Split(sizesLine, string(' '))

	if length := len(sizes); length != 2 {
		panic(fmt.Errorf("Expected the plateau line to have length of 2 but got %v", length))
	}

	return convertStringToInt(sizes[0]), convertStringToInt(sizes[1])
}

func readRoverData(roverLine string) (int, int, string) {
	roverData := strings.Split(roverLine, string(' '))

	if length := len(roverData); length != 3 {
		panic(fmt.Errorf("Expected the rover line to have 3 items but got %v", length))
	}

	return convertStringToInt(roverData[0]), convertStringToInt(roverData[1]), roverData[2]
}

func formatforOutput(rover roverType.Rover) string {
	return fmt.Sprintf("%v %v %v", rover.X, rover.Y, rover.Orientation)
}

// ProcessLines processes all lines of the input
func ProcessLines(lines []string) []string {
	var result []string

	sizeX, sizeY := readPlateauSize(lines[0])
	lines = lines[1:]

	plateau := plateauType.Plateau{
		SizeX: sizeX,
		SizeY: sizeY,
	}

	log.Printf("created a new plateau %v", plateau)

	for i, line := range lines {
		// After creating the plateau the odd lines are the rovers and the even ones the instructions.
		// we'll pick only the odd lines to create the rovers and in the same step take the next line with the instructions
		if i%2 != 0 {
			continue
		}

		x, y, orientation := readRoverData(line)
		rover := roverType.Rover{X: x, Y: y, Orientation: orientation}
		log.Printf("created a new rover %v", rover)

		for _, letter := range lines[i+1] {
			instruction := string(letter)
			log.Printf(`received instruction "%v"`, instruction)
			rover = instructions.ProcessInstruction(instruction, rover, plateau)
			log.Printf(`rover state after processing the instruction "%v": %v`, instruction, rover)
		}
		// As the rovers don't run concurrently it is safe
		// to add the rover to the plateau after processing all the instructions.
		// If the rovers should be able to run concurrently,
		// we have to save their position on the plateau after every instruction.
		plateau.Rovers = append(plateau.Rovers, rover)
		result = append(result, formatforOutput(rover))
	}

	return result
}
