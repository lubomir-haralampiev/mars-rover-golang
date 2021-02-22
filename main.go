package main

import (
	"bufio"
	"log"
	"mars-rover-golang/lineprocessor"
	"os"
)

const inputFile = "input.txt"

func main() {
	// If the file is big it's better to process every line right after it's read.
	// As this is my first time coding in Go I concentrate on the solution and read all the lines, because they are just 5.
	var lines []string

	log.SetOutput(os.Stdout)
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

	log.Printf("read %d lines from the input", len(lines))

	result := lineprocessor.ProcessLines(lines)
	log.Println()
	log.Println("Output:")
	log.Println("-------")
	for _, line := range result {
		log.Println(line)
	}
}
