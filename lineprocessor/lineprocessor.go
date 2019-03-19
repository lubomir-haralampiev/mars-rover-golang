package lineprocessor

import (
	"fmt"
	"log"
	"mars-rover-golang/plateau"
	"strconv"
	"strings"
)

func convertSizeToInt(size string, channel chan int) {
	intValue, err := strconv.Atoi(size)
	if err != nil {
		panic(fmt.Errorf("Can not convert %v to string", size))
	}

	channel <- intValue
}

func convertSizesToInt(sizesLine string) (int, int) {
	sizes := strings.Split(sizesLine, string(' '))

	if length := len(sizes); length != 2 {
		panic(fmt.Errorf("Expected the plateau line to have length of 2 but got %v", length))
	}

	channel := make(chan int)
	go convertSizeToInt(sizes[0], channel)
	go convertSizeToInt(sizes[1], channel)

	return <-channel, <-channel
}

// ProcessLines processes all lines of the input
func ProcessLines(lines []string) {

	sizeX, sizeY := convertSizesToInt(lines[0])
	log.Printf("%T %v", sizeX, sizeX)
	log.Printf("%T %v", sizeY, sizeY)

	newPlateau := plateau.Plateau{
		SizeX: sizeX,
		SizeY: sizeY,
	}

	log.Println(newPlateau)
	// for _, line := range lines {
	// 	log.Println(line)
	// }
}
