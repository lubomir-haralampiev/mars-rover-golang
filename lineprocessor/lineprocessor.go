package lineprocessor

import (
	"fmt"
	"log"
	"mars-rover-golang/plateau"
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

	log.Println(newPlateau)

	// log.Println(lines[1])
	log.Println(readRoverData(lines[1]))

	// for _, line := range lines {
	// 	log.Println(line)
	// }
}
