package lineprocessor

import (
	"fmt"
	"log"
	"mars-rover-golang/instructions"
	"mars-rover-golang/plateau"
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

// ProcessLines processes all lines of the input
func ProcessLines(lines []string) {
	sizeX, sizeY := readPlateauSize(lines[0])

	newPlateau := plateau.Plateau{
		SizeX: sizeX,
		SizeY: sizeY,
	}

	log.Printf("created a new plateau %v", newPlateau)

	x, y, orientation := readRoverData(lines[1])
	rover := roverType.Rover{X: x, Y: y, Orientation: orientation}
	log.Printf("created a new rover %v", rover)

	for _, letter := range lines[2] {
		instruction := string(letter)
		log.Printf(`received instruction "%v"`, instruction)
		rover = instructions.ProcessInstruction(instruction, rover)
		log.Printf(`rover state after processing the instruction "%v": %v`, instruction, rover)
	}
	log.Println(rover)
	// @TODO add to the Plateau
	// As the rovers don't run concurrently it is safe
	// to save the rovers position after processing all the instructions.
	// If the rovers should be able to run concurrently,
	// we have to save their position on the plateau after every instruction.

	// newRover = instructions.ProcessInstruction("L", newRover)
	// log.Println(newRover)

	// for _, line := range lines {
	// 	log.Println(line)
	// }
}
