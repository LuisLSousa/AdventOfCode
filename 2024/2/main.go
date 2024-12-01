package main

import (
	"bufio"
	"log"
	"os"
)

const inputFileName = "input.txt"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// do stuff
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s", err)
	}
}
