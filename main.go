package main

import (
	"bufio"
	"log"
	"os"
)

const inputFile = "input.txt"

func main() {
	var lines []string

	// log.SetOutput(os.Stdout)
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// !file.Close has the error at the first place!
		if err := file.Close(); err != nil {
			log.Printf("Error closing the input file - %s", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// defer func() {
	// 	panicErr := recover()
	// 	//if panicErr != nil {
	// 	log.Println(panicErr)
	// 	panic(errors.New("a new constructred error"))
	// 	//}
	// }()
	// panic("Error 666")
	log.Println(lines)
}
