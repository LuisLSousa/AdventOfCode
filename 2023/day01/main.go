package main

import "log"

const inputFileName = "input.txt"

func main() {
	err := part1()
	if err != nil {
		log.Fatalf("error when executing part 1: %s", err)
	}

	err = part2()
	if err != nil {
		log.Fatalf("error when executing part 2: %s", err)
	}
}
