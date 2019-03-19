package lineprocessor

import (
	"fmt"
	"log"
	"mars-rover-golang/plateau"
	"strconv"
	"strings"
)

func convertSizeToInt(size string) int {
	intValue, err := strconv.Atoi(size)
	if err != nil {
		panic(fmt.Errorf("Can not convert %v to string", size))
	}

	return intValue
}

func convertSizesToInt(sizesLine string) (int, int) {
	sizes := strings.Split(sizesLine, string(' '))

	if length := len(sizes); length != 2 {
		panic(fmt.Errorf("Expected the plateau line to have length of 2 but got %v", length))
	}

	// @TODO parallelize it. use channels ?
	sizeX := convertSizeToInt(sizes[0])
	sizeY := convertSizeToInt(sizes[1])

	return sizeX, sizeY
}

// ProcessLines processes all lines of the input
func ProcessLines(lines []string) {

	sizeX, sizeY := convertSizesToInt(lines[0])
	log.Printf("%T %v", sizeX, sizeX)
	log.Printf("%T %v", sizeY, sizeY)

	plateau := plateau.Plateau{
		SizeX: sizeX,
		SizeY: sizeY,
	}

	log.Println(plateau)
	// for _, line := range lines {
	// 	log.Println(line)
	// }
}
