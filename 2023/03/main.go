package main

import "log"

const inputFileName = "test.txt"

func main() {
	err := part1()
	if err != nil {
		log.Fatalf("error when executing part 1: %s", err)
	}
}
