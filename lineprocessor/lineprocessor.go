package lineprocessor

import (
	"log"
	"mars-rover-golang/plateau"
	"strings"
)

// ProcessLines processes all lines of the input
func ProcessLines(lines []string) {
	sizes := strings.Split(lines[0], string(' '))

	sizeX, sizeY := sizes[0], sizes[1]
	log.Println(sizeX)
	log.Println(sizeY)

	// @TODO convert to integer:
	// v := "10"
	// if s, err := strconv.Atoi(v); err == nil {
	// 	fmt.Printf("%T, %v", s, s)
	// }

	plateau := plateau.Plateau{
		SizeX: 5,
		SizeY: 5,
	}

	log.Println(plateau)
	// for _, line := range lines {
	// 	log.Println(line)
	// }
}
